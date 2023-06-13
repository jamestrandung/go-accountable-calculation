package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormula_MarshalJSON(t *testing.T) {
	mockValue1 := newMockValueWithFormula(t)
	mockSyntaxOperand1 := &SyntaxOperand{Name: "TestOperand1"}
	mockValue1.On("ToSyntaxOperand", OpTransparent).Return(mockSyntaxOperand1).Once()

	mockValue2 := newMockValueWithFormula(t)
	mockSyntaxOperand2 := &SyntaxOperand{Name: "TestOperand2", StageIdx: 1}
	mockValue2.On("ToSyntaxOperand", OpTransparent).Return(mockSyntaxOperand2).Once()

	mockValue3 := newMockValueWithFormula(t)
	innerNode := NewSyntaxNode(OpCategoryTwoValMiddleOp, OpTransparent, "TestInnerOp", []any{"staticValue"})
	mockSyntaxOperand3 := NewSyntaxOperandWithFormula(innerNode, true)
	mockValue3.On("ToSyntaxOperand", OpTransparent).Return(mockSyntaxOperand3).Once()

	syntaxNode := NewSyntaxNode(OpCategoryFunctionCall, OpTransparent, "TestOp", []any{mockValue1, mockValue2, mockValue3})

	wantedJSON := "{\"Category\":\"FunctionCall\",\"Operation\":\"TestOp\",\"Operands\":[{\"Name\":\"TestOperand1\"},{\"Name\":\"TestOperand2\",\"StageIdx\":1},{\"Node\":{\"Category\":\"TwoValMiddleOp\",\"Operation\":\"TestInnerOp\",\"Operands\":[{\"StaticValue\":\"staticValue\"}]},\"WrapInParentheses\":true}]}"

	actualJSON, err := syntaxNode.MarshalJSON()
	assert.Equal(t, wantedJSON, string(actualJSON), "marshal result should be %v", wantedJSON)
	assert.Nil(t, err, "error should be nil")
}

func TestNewSyntaxOperand(t *testing.T) {
	mockValueOps, cleanup := MockValueOps(t)
	defer cleanup()

	mockValue := newMockValueWithFormula(t)

	mockValueOps.On("Identify", mockValue).Return("TestIdentity").Once()

	actual := NewSyntaxOperand(mockValue)
	wanted := &SyntaxOperand{Name: "TestIdentity"}

	assert.Equal(t, wanted, actual)
}

func TestNewSyntaxOperandWithStageIdx(t *testing.T) {
	mockValueOps, cleanup := MockValueOps(t)
	defer cleanup()

	mockValue := newMockValueWithFormula(t)

	mockValueOps.On("Identify", mockValue).Return("TestIdentity").Once()

	actual := NewSyntaxOperandWithStageIdx(mockValue, 5)
	wanted := &SyntaxOperand{Name: "TestIdentity", StageIdx: 5}

	assert.Equal(t, wanted, actual)
}

func TestNewSyntaxOperandWithStaticValue(t *testing.T) {
	value := "test"

	actual := NewSyntaxOperandWithStaticValue(value)
	wanted := &SyntaxOperand{StaticValue: value}

	assert.Equal(t, wanted, actual)
}

func TestNewSyntaxOperandWithNode(t *testing.T) {
	node := &SyntaxNode{}
	wrapInParentheses := true

	actual := NewSyntaxOperandWithFormula(node, wrapInParentheses)
	wanted := &SyntaxOperand{Node: node, WrapInParentheses: wrapInParentheses}

	assert.Equal(t, wanted, actual)
}
