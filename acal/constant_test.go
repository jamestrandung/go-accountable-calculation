package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstant_IsNil(t *testing.T) {
	var nilConstant *Constant[int]

	assert.True(t, nilConstant.IsNil())

	constant := NewConstant(2)

	assert.False(t, constant.IsNil())
}

func TestConstant_ToSyntaxOperand(t *testing.T) {
	constant := NewConstant(2)

	actual := constant.ToSyntaxOperand(OpTransparent)

	assert.Equal(
		t, &SyntaxOperand{
			StaticValue: "2",
		}, actual,
	)
}

func TestConstant_GetFormula(t *testing.T) {
	constant := NewConstant(2)

	actual := constant.GetFormulaFn()()

	assert.Equal(
		t, &SyntaxNode{
			category: OpCategoryAssignStatic,
			op:       OpTransparent,
			opDesc:   "2",
			operands: nil,
		}, actual,
	)
}
