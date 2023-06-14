package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/shopspring/decimal"
)

const (
	opAdd acal.Op = iota
	opSubtract
	opMultiply
	opDivide
	opAbs
	opCeil
	opFloor
	opRound
	opPow

	maxOpName          = "Max"
	minOpName          = "Min"
	medianOpName       = "Median"
	boundBetweenOpName = "BoundBetween"
)

func (p opProvider) Add(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opAdd, "+", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Add(b)
		},
	)
}

func (p opProvider) Sub(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opSubtract, "-", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Sub(b)
		},
	)
}

func (p opProvider) Mul(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opMultiply, "*", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Mul(b)
		},
	)
}

func (p opProvider) Div(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opDivide, "/", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Div(b)
		},
	)
}

func Max(values ...acal.TypedValue[decimal.Decimal]) Simple {
	if len(values) == 0 {
		return NilFloat
	}

	return PerformDecimalFunctionCall(
		"Max",
		func(decimals ...decimal.Decimal) decimal.Decimal {
			if len(decimals) == 0 {
				return decimal.Zero
			}

			return decimal.Max(decimals[0], decimals[1:]...)
		}, values...,
	)
}
