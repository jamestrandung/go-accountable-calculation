package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerformUnaryLogicOp(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "value is nil",
			test: func(t *testing.T) {
				mockTypedValue := acal.NewMockTypedValue[int](t)
				mockTypedValue.On("IsNil").
					Return(true).
					Once()

				actual := PerformUnaryLogicOp[int](
					mockTypedValue, "DummyFnName", func(v int) bool {
						return false
					},
				)

				assert.Nil(t, actual)
			},
		},
		{
			desc: "value is not nil",
			test: func(t *testing.T) {
				fnName := "DummyFnName"

				mockTypedValue := acal.NewMockTypedValue[int](t)
				mockTypedValue.On("IsNil").
					Return(false).
					Once()
				mockTypedValue.On("GetTypedValue").
					Return(1).
					Once()

				mockFormulaBuilder, cleanupFn := acal.MockFormulaBuilder(t)
				defer cleanupFn()

				dummyFormula := &acal.SyntaxNode{}

				mockFormulaBuilder.On("NewFormulaFunctionCall", fnName, mockTypedValue).
					Return(dummyFormula).
					Once()

				actual := PerformUnaryLogicOp[int](
					mockTypedValue, fnName, func(v int) bool {
						assert.Equal(t, 1, v)
						return true
					},
				)

				assert.True(t, actual.Bool())

				formula := actual.GetFormulaFn()()

				assert.Equal(t, dummyFormula, formula)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(
			sc.desc, func(t *testing.T) {
				sc.test(t)
			},
		)
	}
}

func TestPerformBinaryLogicOp(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "both values are nil",
			test: func(t *testing.T) {
				mockTypedValue1 := acal.NewMockTypedValue[int](t)
				mockTypedValue1.On("IsNil").
					Return(true).
					Once()

				mockTypedValue2 := acal.NewMockTypedValue[int](t)
				mockTypedValue2.On("IsNil").
					Return(true).
					Once()

				actual := PerformBinaryLogicOp[int](
					mockTypedValue1, mockTypedValue2, acal.OpTransparent, "SomeOp", func(a, b int) bool {
						return false
					},
				)

				assert.Nil(t, actual)
			},
		},
		{
			desc: "at least one value is NOT nil",
			test: func(t *testing.T) {
				mockTypedValue1 := acal.NewMockTypedValue[int](t)
				mockTypedValue1.On("IsNil").
					Return(true).
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

				mockFormulaBuilder, cleanupFn := acal.MockFormulaBuilder(t)
				defer cleanupFn()

				dummyFormula := &acal.SyntaxNode{}

				mockFormulaBuilder.On("NewFormulaTwoValMiddleOp", mockTypedValue1, mockTypedValue2, acal.OpTransparent, "SomeOp").
					Return(dummyFormula).
					Once()

				actual := PerformBinaryLogicOp[int](
					mockTypedValue1, mockTypedValue2, acal.OpTransparent, "SomeOp", func(a, b int) bool {
						assert.Equal(t, 0, a)
						assert.Equal(t, 1, b)
						return true
					},
				)

				assert.True(t, actual.Bool())

				formula := actual.GetFormulaFn()()

				assert.Equal(t, dummyFormula, formula)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(
			sc.desc, func(t *testing.T) {
				sc.test(t)
			},
		)
	}
}
