package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/shopspring/decimal"
)

// PerformUnaryDecimalOp returns a Simple to represent the result of
// performing unaryOpFn on the value of the provided acal.TypedValue.
func PerformUnaryDecimalOp(
    tv acal.TypedValue[decimal.Decimal],
    opDesc string,
    unaryOpFn func(a decimal.Decimal) decimal.Decimal,
) Simple {
    tv = acal.ReplaceIfNil(tv)

    return MakeSimpleWithFormula(
        unaryOpFn(tv.GetTypedValue()),
        acal.FormulaBuilder.NewFormulaFunctionCall(opDesc, tv),
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
    tv1 = acal.ReplaceIfNil(tv1)
    tv2 = acal.ReplaceIfNil(tv2)

    return MakeSimpleWithFormula(
        binaryOpFn(tv1.GetTypedValue(), tv2.GetTypedValue()),
        acal.FormulaBuilder.NewFormulaTwoValMiddleOp(tv1, tv2, op, opDesc),
    )
}

// PerformDecimalFunctionCall returns a Simple to represent the result of
// performing fn on the values of the provided acal.TypedValue.
func PerformDecimalFunctionCall(
    opDesc string,
    fn func(decimals ...decimal.Decimal) decimal.Decimal,
    values ...acal.TypedValue[decimal.Decimal],
) Simple {
    decimals := make([]decimal.Decimal, len(values))

    for idx, value := range values {
        value = acal.ReplaceIfNil(value)

        values[idx] = value
        decimals[idx] = value.GetTypedValue()
    }

    return MakeSimpleWithFormula(
        fn(decimals...),
        acal.FormulaBuilder.NewFormulaFunctionCall(opDesc, values),
    )
}
