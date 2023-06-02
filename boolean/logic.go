package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
)

// PerformUnaryLogicOp returns a boolean.Simple to represent the result of
// performing unaryOpFn on the values of the provided acal.TypedValue.
func PerformUnaryLogicOp[T any](
	tv acal.TypedValue[T],
	fnName string,
	unaryOpFn func(v T) bool,
) Simple {
	if acal.IsNilValue(tv) {
		return NilBool
	}

	return MakeSimpleWithFormula(unaryOpFn(tv.GetTypedValue()), acal.FormulaBuilder.NewFormulaFunctionCall(fnName, tv))
}

// PerformBinaryLogicOp returns a boolean.Simple to represent the result of
// performing binaryOpFn on the values of the provided acal.TypedValue.
func PerformBinaryLogicOp[T any](
	tv1 acal.TypedValue[T],
	tv2 acal.TypedValue[T],
	op acal.Op,
	opDesc string,
	binaryOpFn func(a, b T) bool,
) Simple {
	if acal.IsNilValue(tv1) && acal.IsNilValue(tv2) {
		return NilBool
	}

	return MakeSimpleWithFormula(
		binaryOpFn(acal.ExtractTypedValue(tv1), acal.ExtractTypedValue(tv2)),
		acal.FormulaBuilder.NewFormulaTwoValMiddleOp(tv1, tv2, op, opDesc),
	)
}
