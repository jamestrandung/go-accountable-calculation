package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/shopspring/decimal"
)

// PerformUnaryDecimalOp returns a Simple to represent the result of
// performing unaryOpFn on the value of the provided acal.TypedValue.
func PerformUnaryDecimalOp(
	tv acal.TypedValue[decimal.Decimal],
	fnName string,
	unaryOpFn func(a decimal.Decimal) decimal.Decimal,
) Simple {
	tv = acal.PreProcessOperand(tv)

	return MakeSimpleWithFormula(
		unaryOpFn(tv.GetTypedValue()),
		func() *acal.SyntaxNode {
			return acal.NewFormulaForFunctionCall(fnName, tv)
		},
	)
}

// PerformBinaryDecimalOp returns a Simple to represent the result of
// performing binaryOpFn on the values of the provided acal.TypedValue.
func PerformBinaryDecimalOp(
	tv1 acal.TypedValue[decimal.Decimal],
	tv2 acal.TypedValue[decimal.Decimal],
	op acal.Op,
	opDesc string,
	binaryOpFn func(a, b decimal.Decimal) decimal.Decimal,
) Simple {
	tv1 = acal.PreProcessOperand(tv1)
	tv2 = acal.PreProcessOperand(tv2)

	return MakeSimpleWithFormula(
		binaryOpFn(tv1.GetTypedValue(), tv2.GetTypedValue()),
		func() *acal.SyntaxNode {
			return acal.NewFormulaForTwoValMiddleOp(tv1, tv2, op, opDesc)
		},
	)
}

// PerformDecimalFunctionCall returns a Simple to represent the result of
// performing fn on the values of the provided acal.TypedValue.
func PerformDecimalFunctionCall(
	fnName string,
	fn func(decimals ...decimal.Decimal) decimal.Decimal,
	values ...acal.TypedValue[decimal.Decimal],
) Simple {
	arguments := make([]any, len(values))
	decimals := make([]decimal.Decimal, len(values))

	for idx, value := range values {
		value = acal.PreProcessOperand(value)

		arguments[idx] = value
		decimals[idx] = value.GetTypedValue()
	}

	return MakeSimpleWithFormula(
		fn(decimals...),
		func() *acal.SyntaxNode {
			return acal.NewFormulaForFunctionCall(fnName, arguments...)
		},
	)
}
