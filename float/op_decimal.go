package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/shopspring/decimal"
	"sort"
)

const (
	opAdd acal.Op = iota
	opSubtract
	opMultiply
	opDivide
)

var (
	opLevel = map[acal.Op]int{
		opAdd:      0,
		opSubtract: 0,
		opMultiply: 1,
		opDivide:   1,
	}
)

// Add returns this value + input value.
func (p opProvider) Add(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opAdd, "+", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Add(b)
		},
	)
}

// Sub returns this value - input value.
func (p opProvider) Sub(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opSubtract, "-", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Sub(b)
		},
	)
}

// Mul returns this value * input value.
func (p opProvider) Mul(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opMultiply, "*", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Mul(b)
		},
	)
}

// Div returns this value / input value.
func (p opProvider) Div(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opDivide, "/", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Div(b)
		},
	)
}

// Neg returns -(this value).
func (p opProvider) Neg() Simple {
	return PerformUnaryDecimalOp(
		p.tv, "Negate", func(d decimal.Decimal) decimal.Decimal {
			return d.Neg()
		},
	)
}

// Inv returns 1 / this value.
func (p opProvider) Inv() Simple {
	return PerformUnaryDecimalOp(
		p.tv, "Inverse", func(d decimal.Decimal) decimal.Decimal {
			return decimal.NewFromInt(1).Div(d)
		},
	)
}

// Abs returns the absolute amount of this value.
func (p opProvider) Abs() Simple {
	return PerformUnaryDecimalOp(
		p.tv, "Absolute", func(d decimal.Decimal) decimal.Decimal {
			return d.Abs()
		},
	)
}

// Ceil returns the nearest integer value greater than or equal to this value.
func (p opProvider) Ceil() Simple {
	return PerformUnaryDecimalOp(
		p.tv, "Ceil", func(d decimal.Decimal) decimal.Decimal {
			return d.Ceil()
		},
	)
}

// Floor returns the nearest integer value less than or equal to this value.
func (p opProvider) Floor() Simple {
	return PerformUnaryDecimalOp(
		p.tv, "Floor", func(d decimal.Decimal) decimal.Decimal {
			return d.Floor()
		},
	)
}

// Round rounds this value to the given decimal places. If places < 0, it will
// round the integer part to the nearest 10^(-places).
//
// Example:
//
//		(5.45).Round(1) // 5.5
//	 	(545).Round(-1) // 550
func (p opProvider) Round(decimalPlace acal.TypedValue[decimal.Decimal]) Simple {
	return PerformDecimalFunctionCall(
		"Round",
		func(decimals ...decimal.Decimal) decimal.Decimal {
			if len(decimals) < 2 {
				return decimal.Zero
			}

			return decimals[0].Round((int32)(decimals[1].IntPart()))
		},
		p.tv, decimalPlace,
	)
}

// Max returns the largest amount amongst the given values.
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

// Min returns the smallest amount amongst the given values.
func Min(values ...acal.TypedValue[decimal.Decimal]) Simple {
	if len(values) == 0 {
		return NilFloat
	}

	return PerformDecimalFunctionCall(
		"Min",
		func(decimals ...decimal.Decimal) decimal.Decimal {
			if len(decimals) == 0 {
				return decimal.Zero
			}

			return decimal.Min(decimals[0], decimals[1:]...)
		}, values...,
	)
}

// Average returns the average amount of the given values.
func Average(values ...acal.TypedValue[decimal.Decimal]) Simple {
	if len(values) == 0 {
		return NilFloat
	}

	return PerformDecimalFunctionCall(
		"Average",
		func(decimals ...decimal.Decimal) decimal.Decimal {
			if len(decimals) == 0 {
				return decimal.Zero
			}

			return decimal.Avg(decimals[0], decimals[1:]...)
		}, values...,
	)
}

// Median returns the median of the given values.
func Median(values ...acal.TypedValue[decimal.Decimal]) Simple {
	if len(values) == 0 {
		return NilFloat
	}

	return PerformDecimalFunctionCall(
		"Median",
		func(decimals ...decimal.Decimal) decimal.Decimal {
			count := len(decimals)
			if count == 0 {
				return decimal.Zero
			}

			sort.Slice(
				decimals, func(i, j int) bool {
					return decimals[i].LessThan(decimals[j])
				},
			)

			if count%2 == 0 {
				return decimals[count/2-1].Add(decimals[count/2]).Div(decimal.NewFromInt(2))
			}

			return decimals[count/2]
		}, values...,
	)
}

// BoundBetween bounds the given value between the 2 limits.
func BoundBetween(
	boundMe acal.TypedValue[decimal.Decimal],
	lowerBound acal.TypedValue[decimal.Decimal],
	upperBound acal.TypedValue[decimal.Decimal],
) Simple {
	return PerformDecimalFunctionCall(
		"BoundBetween",
		func(decimals ...decimal.Decimal) decimal.Decimal {
			if len(decimals) < 3 {
				return decimal.Zero
			}

			result := decimal.Min(decimals[0], decimals[2])
			result = decimal.Max(result, decimals[1])

			return result
		},
		boundMe, lowerBound, upperBound,
	)
}

// Pow returns a^b.
func Pow(a, b acal.TypedValue[decimal.Decimal]) Simple {
	return PerformDecimalFunctionCall(
		"Pow",
		func(decimals ...decimal.Decimal) decimal.Decimal {
			if len(decimals) < 2 {
				return decimal.Zero
			}

			return decimals[0].Pow(decimals[1])
		},
		a, b,
	)
}
