package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple_IsNil(t *testing.T) {
	var nilSimple *Simple
	assert.True(t, nilSimple.IsNil())

	fSimple := NewSimple("false", false)
	assert.False(t, fSimple.IsNil())
}

func TestSimple_GetTypedValue(t *testing.T) {
	var nilSimple *Simple
	assert.False(t, nilSimple.GetTypedValue())

	tSimple := NewSimple("true", true)
	assert.True(t, tSimple.GetTypedValue())

	fSimple := NewSimple("false", false)
	assert.False(t, fSimple.GetTypedValue())
}

func TestSimple_GetToSyntaxOperand(t *testing.T) {
	defer func(original func(s *Simple, nextOp acal.Op) *acal.SyntaxOperand) {
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

				simple := NewSimple("Simple", true)

				toBaseSyntaxOperand = func(s *Simple, nextOp acal.Op) *acal.SyntaxOperand {
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

				simple := &Simple{
					Simple: acal.NewSimpleWithFormula[bool](
						true, func() *acal.SyntaxNode {
							return dummySyntaxNode
						},
					),
				}

				toBaseSyntaxOperand = func(s *Simple, nextOp acal.Op) *acal.SyntaxOperand {
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
			desc: "nextOp is different from lastOp",
			test: func(t *testing.T) {
				dummyOp := opAnd
				dummySyntaxNode := acal.NewSyntaxNode(acal.OpCategoryFunctionCall, opOr, "OR", nil)

				simple := &Simple{
					Simple: acal.NewSimpleWithFormula[bool](
						true, func() *acal.SyntaxNode {
							return dummySyntaxNode
						},
					),
				}

				toBaseSyntaxOperand = func(s *Simple, nextOp acal.Op) *acal.SyntaxOperand {
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

func TestSimple_SelfReplaceIfNil(t *testing.T) {
	var nilSimple *Simple
	assert.Equal(t, NilBool, nilSimple.SelfReplaceIfNil())

	simple := NewSimple("Simple", true)
	assert.Equal(t, simple, simple.SelfReplaceIfNil())
}

func TestSimple_Bool(t *testing.T) {
	var nilSimple *Simple
	assert.False(t, nilSimple.Bool())

	tSimple := NewSimple("True", true)
	assert.True(t, tSimple.Bool())

	fSimple := NewSimple("False", false)
	assert.False(t, fSimple.Bool())
}

func TestSimple_Anchor(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil Simple",
			test: func(t *testing.T) {
				var nilSimple *Simple

				actual := nilSimple.Anchor("Something")

				assert.Equal(t, "Something", actual.GetName())
				assert.False(t, actual.Bool())
			},
		},
		{
			desc: "same Simple",
			test: func(t *testing.T) {
				simple := NewSimpleFrom(acal.NewConstant(true))

				actual := simple.Anchor("Something")

				assert.Equal(t, "Something", actual.GetName())
				assert.True(t, actual.Bool())
				assert.Equal(t, simple, actual)
			},
		},
		{
			desc: "new Simple",
			test: func(t *testing.T) {
				simple := NewSimple("AlreadyAnchored", true)

				actual := simple.Anchor("Something")

				assert.Equal(t, "Something", actual.GetName())
				assert.True(t, actual.Bool())
				assert.NotEqual(t, simple, actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
