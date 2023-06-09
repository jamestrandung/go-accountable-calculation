package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
	"github.com/shopspring/decimal"
)

const (
	opPlus acal.Op = iota
	opMinus
	opMultiply
	opDivide
	opAbs
	opCeil
	opFloor
	opRound
	opPow

	opEquals acal.Op = iota
	opNotEquals
	opLargerThanEquals
	opLargerThan
	opSmallerThanEquals
	opSmallerThan

	maxOpName          = "Max"
	minOpName          = "Min"
	medianOpName       = "Median"
	boundBetweenOpName = "BoundBetween"
)

var (
	opDesc = map[acal.Op]string{
		opPlus:     "+",
		opMinus:    "-",
		opMultiply: "*",
		opDivide:   "/",
		opAbs:      "Abs",
		opCeil:     "Ceil",
		opFloor:    "Floor",
		opRound:    "Round",
		opPow:      "Pow",
	}

	opLevel = map[acal.Op]int{
		opPlus:     0,
		opMinus:    0,
		opMultiply: 1,
		opDivide:   1,
	}
)

type opProvider struct {
	tv acal.TypedValue[decimal.Decimal]
}

// IsPositive returns whether this value is positive.
func (p opProvider) IsPositive() boolean.Simple {
	return p.LargerThan(Zero)
}

// IsNegative returns whether this value is negative.
func (p opProvider) IsNegative() boolean.Simple {
	return p.SmallerThan(Zero)
}

// Equals returns whether this value equals to the input value.
func (p opProvider) Equals(v acal.TypedValue[decimal.Decimal]) boolean.Simple {
	return boolean.PerformBinaryLogicOp[decimal.Decimal](
		p.tv, v, opEquals, "==", func(a, b decimal.Decimal) bool {
			return a.Equal(b)
		},
	)
}

// NotEquals returns whether this value does not equal to the input value.
func (p opProvider) NotEquals(v acal.TypedValue[decimal.Decimal]) boolean.Simple {
	return boolean.PerformBinaryLogicOp[decimal.Decimal](
		p.tv, v, opNotEquals, "!=", func(a, b decimal.Decimal) bool {
			return !a.Equal(b)
		},
	)
}

// LargerThan returns whether this value is larger than the input value.
func (p opProvider) LargerThan(v acal.TypedValue[decimal.Decimal]) boolean.Simple {
	return boolean.PerformBinaryLogicOp[decimal.Decimal](
		p.tv, v, opLargerThan, ">", func(a, b decimal.Decimal) bool {
			return a.GreaterThan(b)
		},
	)
}

// LargerThanEquals returns whether this value is larger than or equal to the input value.
func (p opProvider) LargerThanEquals(v acal.TypedValue[decimal.Decimal]) boolean.Simple {
	return boolean.PerformBinaryLogicOp[decimal.Decimal](
		p.tv, v, opLargerThanEquals, ">=", func(a, b decimal.Decimal) bool {
			return a.GreaterThanOrEqual(b)
		},
	)
}

// SmallerThan returns whether this value is smaller than the input value.
func (p opProvider) SmallerThan(v acal.TypedValue[decimal.Decimal]) boolean.Simple {
	return boolean.PerformBinaryLogicOp[decimal.Decimal](
		p.tv, v, opSmallerThan, "<", func(a, b decimal.Decimal) bool {
			return a.LessThan(b)
		},
	)
}

// SmallerThanEquals returns whether this value is smaller than or equal to the input value.
func (p opProvider) SmallerThanEquals(v acal.TypedValue[decimal.Decimal]) boolean.Simple {
	return boolean.PerformBinaryLogicOp[decimal.Decimal](
		p.tv, v, opSmallerThanEquals, "<=", func(a, b decimal.Decimal) bool {
			return a.LessThanOrEqual(b)
		},
	)
}

func (p opProvider) Plus(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opPlus, "+", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Add(b)
		},
	)
}

func (p opProvider) Minus(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opMinus, "-", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Sub(b)
		},
	)
}

func (p opProvider) Multiply(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opMultiply, "*", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Mul(b)
		},
	)
}

func (p opProvider) Divide(tv acal.TypedValue[decimal.Decimal]) Simple {
	return PerformBinaryDecimalOp(
		p.tv, tv, opDivide, "/", func(a, b decimal.Decimal) decimal.Decimal {
			return a.Div(b)
		},
	)
}
