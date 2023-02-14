package acal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFormulaBuilder_NewFormulaFunctionCall(t *testing.T) {
	aValMock := &mockValueWithFormula{}
	aValMock.On("SelfReplaceIfNil").Return(aValMock).Once()

	staticValue := "staticValue"
	testFnName := "testFnName"

	expected := NewSyntaxNode(
		OpCategoryFunctionCall,
		OpTransparent,
		testFnName,
		[]any{aValMock, staticValue},
	)

	actual := FormulaBuilder.NewFormulaFunctionCall(testFnName, aValMock, staticValue)

	assert.Equal(t, expected, actual)
	mock.AssertExpectationsForObjects(t, aValMock)
}

func TestFormulaBuilder_NewFormulaTwoValMiddleOp(t *testing.T) {
	aValMock1 := &mockValueWithFormula{}
	aValMock1.On("SelfReplaceIfNil").Return(aValMock1).Once()

	aValMock2 := &mockValueWithFormula{}
	aValMock2.On("SelfReplaceIfNil").Return(aValMock2).Once()

	testOp := OpTransparent
	testOpDesc := "TestOpDesc"

	expected := NewSyntaxNode(OpCategoryTwoValMiddleOp, testOp, testOpDesc, []any{aValMock1, aValMock2})

	actual := FormulaBuilder.NewFormulaTwoValMiddleOp(aValMock1, aValMock2, testOp, "TestOpDesc")

	assert.Equal(t, expected, actual)
	mock.AssertExpectationsForObjects(t, aValMock1, aValMock2)
}
