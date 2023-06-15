package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerformUnaryLogicOp(t *testing.T) {
	fnName := "DummyFnName"

	mockTypedValue := acal.NewMockTypedValue[int](t)
	mockTypedValue.On("IsNil").
		Return(false).
		Once()
	mockTypedValue.On("GetTypedValue").
		Return(1).
		Once()

	actual := PerformUnaryLogicOp[int](
		mockTypedValue, fnName, func(v int) bool {
			assert.Equal(t, 1, v)
			return true
		},
	)

	assert.True(t, actual.Bool())

	formula := actual.GetFormulaFn()()

	assert.Equal(t, acal.NewSyntaxNode(acal.OpCategoryFunctionCall, acal.OpTransparent, fnName, []any{mockTypedValue}), formula)
}

func TestPerformBinaryLogicOp(t *testing.T) {
	dummyOp := opOr
	dummyOpDesc := "SomeOp"

	mockTypedValue1 := acal.NewMockTypedValue[int](t)
	mockTypedValue1.On("IsNil").
		Return(false).
		Once()
	mockTypedValue1.On("GetTypedValue").
		Return(0).
		Once()

	mockTypedValue2 := acal.NewMockTypedValue[int](t)
	mockTypedValue2.On("IsNil").
		Return(false).
		Once()
	mockTypedValue2.On("GetTypedValue").
		Return(1).
		Once()

	actual := PerformBinaryLogicOp[int](
		mockTypedValue1, mockTypedValue2, dummyOp, dummyOpDesc, func(a, b int) bool {
			assert.Equal(t, 0, a)
			assert.Equal(t, 1, b)
			return true
		},
	)

	assert.True(t, actual.Bool())

	formula := actual.GetFormulaFn()()

	assert.Equal(t, acal.NewSyntaxNode(acal.OpCategoryTwoValMiddleOp, dummyOp, dummyOpDesc, []any{mockTypedValue1, mockTypedValue2}), formula)
}
