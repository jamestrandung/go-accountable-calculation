package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
	"github.com/my-shopspring/decimal"
)

// Value represents an acal.Value of float kind.
//
//go:generate mockery --name=Value --case underscore --inpackage
type Value interface {
	acal.TypedValue[decimal.Decimal]

	// Decimal returns the value of this Value as a decimal.Decimal.
	// If it's nil, a decimal.Decimal value of 0 is returned.
	Decimal() decimal.Decimal
	// Float returns the value of this Value as a float64.
	// If it's nil, 0 is returned.
	Float() float64
}

// Interface governs the methods that Value should provide.
//
//go:generate mockery --name=Interface --case underscore --inpackage
type Interface interface {
	Value
	booleanOpProvider
	decimalOpProvider
}

type decimalOpProvider interface {
	// Add returns this value + input value.
	Add(acal.TypedValue[decimal.Decimal]) Simple
	// Sub returns this value - input value.
	Sub(acal.TypedValue[decimal.Decimal]) Simple
	// Mul returns this value * input value.
	Mul(acal.TypedValue[decimal.Decimal]) Simple
	// Div returns this value / input value.
	Div(acal.TypedValue[decimal.Decimal]) Simple
	// Neg returns -(this value).
	Neg() Simple
	// Inv returns 1 / this value.
	Inv() Simple
	// Abs returns the absolute amount of this value.
	Abs()
	// Ceil returns the nearest integer value greater than or equal to this value.
	Ceil()
	// Floor returns the nearest integer value less than or equal to this value.
	Floor() Simple
	// Round rounds this value to the given decimal places. If places < 0, it will
	// round the integer part to the nearest 10^(-places).
	//
	// Example:
	//  (5.45).Round(1) // 5.5
	// 	(545).Round(-1) // 550
	//
	Round(decimalPlace acal.TypedValue[decimal.Decimal]) Simple
}

type booleanOpProvider interface {
	// IsZero returns whether this value is zero.
	IsZero() boolean.Simple
	// IsPositive returns whether this value is positive.
	IsPositive() boolean.Simple
	// IsNegative returns whether this value is negative.
	IsNegative() boolean.Simple
	// Equals returns whether this value equals to the input value.
	Equals(acal.TypedValue[decimal.Decimal]) boolean.Simple
	// NotEquals returns whether this value does not equal to the input value.
	NotEquals(acal.TypedValue[decimal.Decimal]) boolean.Simple
	// LargerThan returns whether this value is larger than the input value.
	LargerThan(acal.TypedValue[decimal.Decimal]) boolean.Simple
	// LargerThanEquals returns whether this value is larger than or equal to the input value.
	LargerThanEquals(acal.TypedValue[decimal.Decimal]) boolean.Simple
	// SmallerThan returns whether this value is smaller than the input value.
	SmallerThan(acal.TypedValue[decimal.Decimal]) boolean.Simple
	// SmallerThanEquals returns whether this value is smaller than or equal to the input value.
	SmallerThanEquals(acal.TypedValue[decimal.Decimal]) boolean.Simple
}
