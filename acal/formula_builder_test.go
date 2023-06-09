package acal

//func TestFormulaBuilder_NewFormulaFunctionCall(t *testing.T) {
//	aValMock := newMockValueWithFormula(t)
//	aValMock.On("SelfReplaceIfNil").Return(aValMock).Once()
//
//	staticValue := "staticValue"
//	testFnName := "testFnName"
//
//	expected := NewSyntaxNode(
//		OpCategoryFunctionCall,
//		OpTransparent,
//		testFnName,
//		[]any{aValMock, staticValue},
//	)
//
//	actual := FormulaBuilder.NewFormulaFunctionCall(testFnName, aValMock, staticValue)
//
//	assert.Equal(t, expected, actual)
//}
//
//func TestFormulaBuilder_NewFormulaTwoValMiddleOp(t *testing.T) {
//	aValMock1 := newMockValueWithFormula(t)
//	aValMock1.On("SelfReplaceIfNil").Return(aValMock1).Once()
//
//	aValMock2 := newMockValueWithFormula(t)
//	aValMock2.On("SelfReplaceIfNil").Return(aValMock2).Once()
//
//	testOp := OpTransparent
//	testOpDesc := "TestOpDesc"
//
//	expected := NewSyntaxNode(OpCategoryTwoValMiddleOp, testOp, testOpDesc, []any{aValMock1, aValMock2})
//
//	actual := FormulaBuilder.NewFormulaTwoValMiddleOp(aValMock1, aValMock2, testOp, "TestOpDesc")
//
//	assert.Equal(t, expected, actual)
//}
