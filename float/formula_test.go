package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerformUnaryDecimalOp(t *testing.T) {
	fnName := "DummyFnName"

	mockTypedValue := acal.NewMockTypedValue[decimal.Decimal](t)
	mockTypedValue.On("IsNil").
		Return(false).
		Once()
	mockTypedValue.On("GetTypedValue").
		Return(decimal.NewFromInt(10)).
		Once()

	actual := PerformUnaryDecimalOp(
		mockTypedValue, fnName, func(v decimal.Decimal) decimal.Decimal {
			assert.Equal(t, decimal.NewFromInt(10), v)
			return decimal.NewFromInt(99)
		},
	)

	assert.Equal(t, decimal.NewFromInt(99), actual.Decimal())

	formula := actual.GetFormulaFn()()

	assert.Equal(t, acal.NewSyntaxNode(acal.OpCategoryFunctionCall, acal.OpTransparent, fnName, []any{mockTypedValue}), formula)
}

func TestPerformBinaryDecimalOp(t *testing.T) {
	fnName := "DummyFnName"

	mockTypedValue1 := acal.NewMockTypedValue[decimal.Decimal](t)
	mockTypedValue1.On("IsNil").
		Return(false).
		Once()
	mockTypedValue1.On("GetTypedValue").
		Return(decimal.NewFromInt(0)).
		Once()

	mockTypedValue2 := acal.NewMockTypedValue[decimal.Decimal](t)
	mockTypedValue2.On("IsNil").
		Return(false).
		Once()
	mockTypedValue2.On("GetTypedValue").
		Return(decimal.NewFromInt(1)).
		Once()

	actual := PerformBinaryDecimalOp(
		mockTypedValue1, mockTypedValue2, acal.OpTransparent, fnName, func(a, b decimal.Decimal) decimal.Decimal {
			assert.Equal(t, decimal.NewFromInt(0), a)
			assert.Equal(t, decimal.NewFromInt(1), b)
			return decimal.NewFromInt(99)
		},
	)

	assert.Equal(t, decimal.NewFromInt(99), actual.Decimal())

	formula := actual.GetFormulaFn()()

	assert.Equal(t, acal.NewSyntaxNode(acal.OpCategoryTwoValMiddleOp, acal.OpTransparent, fnName, []any{mockTypedValue1, mockTypedValue2}), formula)
}

func TestPerformDecimalFunctionCall(t *testing.T) {
	fnName := "DummyFnName"

	mockTypedValue1 := acal.NewMockTypedValue[decimal.Decimal](t)
	mockTypedValue1.On("IsNil").
		Return(false).
		Once()
	mockTypedValue1.On("GetTypedValue").
		Return(decimal.NewFromInt(0)).
		Once()

	mockTypedValue2 := acal.NewMockTypedValue[decimal.Decimal](t)
	mockTypedValue2.On("IsNil").
		Return(false).
		Once()
	mockTypedValue2.On("GetTypedValue").
		Return(decimal.NewFromInt(1)).
		Once()

	actual := PerformDecimalFunctionCall(
		fnName, func(decimals ...decimal.Decimal) decimal.Decimal {
			assert.Equal(t, decimal.NewFromInt(0), decimals[0])
			assert.Equal(t, decimal.NewFromInt(1), decimals[1])
			return decimal.NewFromInt(99)
		}, mockTypedValue1, mockTypedValue2,
	)

	assert.Equal(t, decimal.NewFromInt(99), actual.Decimal())

	formula := actual.GetFormulaFn()()

	assert.Equal(t, acal.NewSyntaxNode(acal.OpCategoryFunctionCall, acal.OpTransparent, fnName, []any{mockTypedValue1, mockTypedValue2}), formula)
}
