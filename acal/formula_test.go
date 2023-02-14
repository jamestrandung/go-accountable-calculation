package acal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFormula_MarshalJSON(t *testing.T) {
	aValMock1 := &mockValueWithFormula{}
	syntaxOperandMock1 := &SyntaxOperand{Name: "TestOperand1"}
	aValMock1.On("ToSyntaxOperand", OpTransparent).Return(syntaxOperandMock1).Once()

	aValMock2 := &mockValueWithFormula{}
	syntaxOperandMock2 := &SyntaxOperand{Name: "TestOperand2", StageIdx: 1}
	aValMock2.On("ToSyntaxOperand", OpTransparent).Return(syntaxOperandMock2).Once()

	aValMock3 := &mockValueWithFormula{}
	innerNode := NewSyntaxNode(OpCategoryTwoValMiddleOp, OpTransparent, "TestInnerOp", []any{"staticValue"})
	syntaxOperandMock3 := NewSyntaxOperandWithFormula(innerNode, true)
	aValMock3.On("ToSyntaxOperand", OpTransparent).Return(syntaxOperandMock3).Once()

	syntaxNode := NewSyntaxNode(OpCategoryFunctionCall, OpTransparent, "TestOp", []any{aValMock1, aValMock2, aValMock3})

	wantedJSON := "{\"Category\":\"FunctionCall\",\"Operation\":\"TestOp\",\"Operands\":[{\"Name\":\"TestOperand1\"},{\"Name\":\"TestOperand2\",\"StageIdx\":1},{\"Node\":{\"Category\":\"TwoValMiddleOp\",\"Operation\":\"TestInnerOp\",\"Operands\":[{\"StaticValue\":\"staticValue\"}]},\"WrapInParentheses\":true}]}"

	actualJSON, err := syntaxNode.MarshalJSON()
	assert.Equal(t, wantedJSON, string(actualJSON), "marshal result should be %v", wantedJSON)
	assert.Nil(t, err, "error should be nil")
}

func TestNewSyntaxOperand(t *testing.T) {
	valueOpsMock, cleanup := MockValueOps()
	defer cleanup()

	aValMock := &mockValueWithFormula{}

	valueOpsMock.On("Identify", aValMock).Return("TestIdentity").Once()

	actual := NewSyntaxOperand(aValMock)
	wanted := &SyntaxOperand{Name: "TestIdentity"}

	assert.Equal(t, wanted, actual)
	mock.AssertExpectationsForObjects(t, valueOpsMock)
}

func TestNewSyntaxOperandWithStageIdx(t *testing.T) {
	valueOpsMock, cleanup := MockValueOps()
	defer cleanup()

	aValMock := &mockValueWithFormula{}

	valueOpsMock.On("Identify", aValMock).Return("TestIdentity").Once()

	actual := NewSyntaxOperandWithStageIdx(aValMock, 5)
	wanted := &SyntaxOperand{Name: "TestIdentity", StageIdx: 5}

	assert.Equal(t, wanted, actual)
	mock.AssertExpectationsForObjects(t, valueOpsMock)
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
