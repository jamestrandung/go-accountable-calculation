package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple_GetTypedValue(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Simple is nil",
			test: func(t *testing.T) {
				s := Simple{
					Simple: nil,
				}

				actual := s.GetTypedValue()
				assert.Equal(t, decimal.Decimal{}, actual)
			},
		},
		{
			desc: "inner Simple is not nil",
			test: func(t *testing.T) {
				s := MakeSimpleFromInt("Simple", 10)

				actual := s.GetTypedValue()
				assert.Equal(t, decimal.NewFromInt(10), actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestSimple_ToSyntaxOperand(t *testing.T) {
	defer func(original func(s Simple, nextOp acal.Op) *acal.SyntaxOperand) {
		toBaseSyntaxOperand = original
	}(toBaseSyntaxOperand)

	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "toBaseSyntaxOperand returns non-nil result",
			test: func(t *testing.T) {
				dummyOp := acal.OpTransparent
				dummySyntaxOperand := &acal.SyntaxOperand{}

				simple := MakeSimpleFromInt("Simple", 10)

				toBaseSyntaxOperand = func(s Simple, nextOp acal.Op) *acal.SyntaxOperand {
					assert.Equal(t, simple, s)
					assert.Equal(t, dummyOp, nextOp)
					return dummySyntaxOperand
				}

				actual := simple.ToSyntaxOperand(dummyOp)
				assert.Equal(t, dummySyntaxOperand, actual)
			},
		},
		{
			desc: "nextOp is transparent",
			test: func(t *testing.T) {
				dummyOp := acal.OpTransparent
				dummySyntaxNode := &acal.SyntaxNode{}

				simple := Simple{
					Simple: acal.NewSimpleWithFormula[decimal.Decimal](
						decimal.NewFromInt(10), func() *acal.SyntaxNode {
							return dummySyntaxNode
						},
					),
				}

				toBaseSyntaxOperand = func(s Simple, nextOp acal.Op) *acal.SyntaxOperand {
					assert.Equal(t, simple, s)
					assert.Equal(t, dummyOp, nextOp)
					return nil
				}

				expected := acal.NewSyntaxOperandWithFormula(dummySyntaxNode, false)

				actual := simple.ToSyntaxOperand(dummyOp)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc: "lastOp is transparent",
			test: func(t *testing.T) {
				dummyOp := opMultiply
				dummySyntaxNode := acal.NewSyntaxNode(acal.OpCategoryFunctionCall, acal.OpTransparent, "Round", nil)

				simple := Simple{
					Simple: acal.NewSimpleWithFormula[decimal.Decimal](
						decimal.NewFromInt(10), func() *acal.SyntaxNode {
							return dummySyntaxNode
						},
					),
				}

				toBaseSyntaxOperand = func(s Simple, nextOp acal.Op) *acal.SyntaxOperand {
					assert.Equal(t, simple, s)
					assert.Equal(t, dummyOp, nextOp)
					return nil
				}

				expected := acal.NewSyntaxOperandWithFormula(dummySyntaxNode, false)

				actual := simple.ToSyntaxOperand(dummyOp)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc: "nextOp level is lower than lastOp",
			test: func(t *testing.T) {
				dummyOp := opAdd
				dummySyntaxNode := acal.NewSyntaxNode(acal.OpCategoryFunctionCall, opMultiply, "*", nil)

				simple := Simple{
					Simple: acal.NewSimpleWithFormula[decimal.Decimal](
						decimal.NewFromInt(10), func() *acal.SyntaxNode {
							return dummySyntaxNode
						},
					),
				}

				toBaseSyntaxOperand = func(s Simple, nextOp acal.Op) *acal.SyntaxOperand {
					assert.Equal(t, simple, s)
					assert.Equal(t, dummyOp, nextOp)
					return nil
				}

				expected := acal.NewSyntaxOperandWithFormula(dummySyntaxNode, false)

				actual := simple.ToSyntaxOperand(dummyOp)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc: "nextOp level is same as lastOp",
			test: func(t *testing.T) {
				dummyOp := opAdd
				dummySyntaxNode := acal.NewSyntaxNode(acal.OpCategoryFunctionCall, opSubtract, "-", nil)

				simple := Simple{
					Simple: acal.NewSimpleWithFormula[decimal.Decimal](
						decimal.NewFromInt(10), func() *acal.SyntaxNode {
							return dummySyntaxNode
						},
					),
				}

				toBaseSyntaxOperand = func(s Simple, nextOp acal.Op) *acal.SyntaxOperand {
					assert.Equal(t, simple, s)
					assert.Equal(t, dummyOp, nextOp)
					return nil
				}

				expected := acal.NewSyntaxOperandWithFormula(dummySyntaxNode, false)

				actual := simple.ToSyntaxOperand(dummyOp)
				assert.Equal(t, expected, actual)
			},
		},
		{
			desc: "nextOp level is higher lastOp",
			test: func(t *testing.T) {
				dummyOp := opMultiply
				dummySyntaxNode := acal.NewSyntaxNode(acal.OpCategoryFunctionCall, opSubtract, "-", nil)

				simple := Simple{
					Simple: acal.NewSimpleWithFormula[decimal.Decimal](
						decimal.NewFromInt(10), func() *acal.SyntaxNode {
							return dummySyntaxNode
						},
					),
				}

				toBaseSyntaxOperand = func(s Simple, nextOp acal.Op) *acal.SyntaxOperand {
					assert.Equal(t, simple, s)
					assert.Equal(t, dummyOp, nextOp)
					return nil
				}

				expected := acal.NewSyntaxOperandWithFormula(dummySyntaxNode, true)

				actual := simple.ToSyntaxOperand(dummyOp)
				assert.Equal(t, expected, actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestSimple_Anchor(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Simple is nil",
			test: func(t *testing.T) {
				var s Simple

				actual := s.Anchor("Simple")

				assert.Equal(t, "Simple", actual.GetName())
				assert.Equal(t, decimal.NewFromInt(0), actual.Decimal())
			},
		},
		{
			desc: "old Simple has a name",
			test: func(t *testing.T) {
				s := MakeSimpleFromInt("Simple", 10)

				actual := s.Anchor("AnotherSimple")

				assert.Equal(t, "AnotherSimple", actual.GetName())
				assert.Equal(t, decimal.NewFromInt(10), actual.Decimal())
				assert.NotEqual(t, s, actual)
				assert.Equal(t, "Simple", s.GetName())
			},
		},
		{
			desc: "old Simple does not have a name",
			test: func(t *testing.T) {
				s1 := MakeSimpleFromInt("Simple", 10)
				s2 := MakeSimpleFromInt("AnotherSimple", 20)
				s3 := s1.Add(s2)

				actual := s3.Anchor("YetAnotherSimple")

				assert.Equal(t, "YetAnotherSimple", actual.GetName())
				assert.Equal(t, decimal.NewFromInt(30), actual.Decimal())
				assert.Equal(t, s3, actual)
				assert.Equal(t, "Simple", s1.GetName())
				assert.Equal(t, "AnotherSimple", s2.GetName())
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestSimple_Decimal(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Simple is nil",
			test: func(t *testing.T) {
				s := Simple{
					Simple: nil,
				}

				actual := s.Decimal()
				assert.Equal(t, decimal.Decimal{}, actual)
			},
		},
		{
			desc: "inner Simple is not nil",
			test: func(t *testing.T) {
				s := MakeSimpleFromInt("Simple", 10)

				actual := s.Decimal()
				assert.Equal(t, decimal.NewFromInt(10), actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestSimple_Float(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Simple is nil",
			test: func(t *testing.T) {
				s := Simple{
					Simple: nil,
				}

				actual := s.Float()
				assert.Equal(t, 0.0, actual)
			},
		},
		{
			desc: "inner Simple is not nil",
			test: func(t *testing.T) {
				s := MakeSimpleFromInt("Simple", 10)

				actual := s.Float()
				assert.Equal(t, 10.0, actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
