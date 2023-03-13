package op

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
)

// PerformUnaryLogicOp returns a boolean.Simple to represent the result of
// performing unaryOpFn on the values of the provided acal.TypedValue.
func PerformUnaryLogicOp[T any](
	tv acal.TypedValue[T],
	fnName string,
	unaryOpFn func(v T) bool,
) *boolean.Simple {
	if acal.IsNilValue(tv) {
		return nil
	}

	return boolean.NewSimpleWithFormula(
		unaryOpFn(tv.GetTypedValue()), func() *acal.SyntaxNode {
			return acal.FormulaBuilder.NewFormulaFunctionCall(fnName, tv)
		},
	)
}

// PerformBinaryLogicOp returns a boolean.Simple to represent the result of
// performing binaryOpFn on the values of the provided acal.TypedValue.
func PerformBinaryLogicOp[T any](
	tv1 acal.TypedValue[T],
	tv2 acal.TypedValue[T],
	op acal.Op,
	opDesc string,
	binaryOpFn func(a, b T) bool,
) *boolean.Simple {
	if acal.IsNilValue(tv1) && acal.IsNilValue(tv2) {
		return nil
	}

	return boolean.NewSimpleWithFormula(
		binaryOpFn(tv1.GetTypedValue(), tv2.GetTypedValue()), func() *acal.SyntaxNode {
			return acal.FormulaBuilder.NewFormulaTwoValMiddleOp(tv1, tv2, op, opDesc)
		},
	)
}
