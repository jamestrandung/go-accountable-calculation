package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormulaBuilder_NewFormulaFunctionCall(t *testing.T) {
	mockValue := newMockValueWithFormula(t)

	staticValue := "staticValue"
	testFnName := "testFnName"

	expected := NewSyntaxNode(
		OpCategoryFunctionCall,
		OpTransparent,
		testFnName,
		[]any{mockValue, staticValue},
	)

	actual := FormulaBuilder.NewFormulaFunctionCall(testFnName, mockValue, staticValue)

	assert.Equal(t, expected, actual)
}

func TestFormulaBuilder_NewFormulaTwoValMiddleOp(t *testing.T) {
	mockValue1 := newMockValueWithFormula(t)
	mockValue2 := newMockValueWithFormula(t)

	testOp := OpTransparent
	testOpDesc := "TestOpDesc"

	expected := NewSyntaxNode(OpCategoryTwoValMiddleOp, testOp, testOpDesc, []any{mockValue1, mockValue2})

	actual := FormulaBuilder.NewFormulaTwoValMiddleOp(mockValue1, mockValue2, testOp, "TestOpDesc")

	assert.Equal(t, expected, actual)
}

func TestPreProcessOperand(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "value is nil",
			test: func(t *testing.T) {
				mockValue := NewMockTypedValue[int](t)

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue).
					Return(true).
					Once()

				actual := PreProcessOperand[int](mockValue)

				assert.Equal(t, ZeroSimple[int]("NilReplacement"), actual)
			},
		},
		{
			desc: "value is a snapshooter",
			test: func(t *testing.T) {
				p := NewProgressive[int]("Progressive")

				actual := PreProcessOperand[int](p)

				assert.Equal(t, p.getSnapshot(), actual)
			},
		},
		{
			desc: "value is non-nil and not a snapshooter",
			test: func(t *testing.T) {
				mockValue := NewMockTypedValue[int](t)

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue).
					Return(false).
					Once()

				actual := PreProcessOperand[int](mockValue)

				assert.Equal(t, mockValue, actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
