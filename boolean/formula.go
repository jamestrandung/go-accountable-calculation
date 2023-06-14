package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
)

// PerformUnaryLogicOp returns a Simple to represent the result of
// performing unaryOpFn on the values of the provided acal.TypedValue.
func PerformUnaryLogicOp[T any](
	tv acal.TypedValue[T],
	fnName string,
	unaryOpFn func(v T) bool,
) Simple {
	tv = acal.PreProcessOperand(tv)

	return MakeSimpleWithFormula(
		unaryOpFn(tv.GetTypedValue()),
		func() *acal.SyntaxNode {
			return acal.FormulaBuilder.NewFormulaFunctionCall(fnName, tv)
		},
	)
}

// PerformBinaryLogicOp returns a Simple to represent the result of
// performing binaryOpFn on the values of the provided acal.TypedValue.
func PerformBinaryLogicOp[T any](
	tv1 acal.TypedValue[T],
	tv2 acal.TypedValue[T],
	op acal.Op,
	opDesc string,
	binaryOpFn func(a, b T) bool,
) Simple {
	tv1 = acal.PreProcessOperand(tv1)
	tv2 = acal.PreProcessOperand(tv2)

	return MakeSimpleWithFormula(
		binaryOpFn(tv1.GetTypedValue(), tv2.GetTypedValue()),
		func() *acal.SyntaxNode {
			return acal.FormulaBuilder.NewFormulaTwoValMiddleOp(tv1, tv2, op, opDesc)
		},
	)
}
