package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
	"github.com/shopspring/decimal"
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

	Add(acal.TypedValue[decimal.Decimal]) Simple
	Sub(acal.TypedValue[decimal.Decimal]) Simple
	Mul(acal.TypedValue[decimal.Decimal]) Simple
	Div(acal.TypedValue[decimal.Decimal]) Simple
	Equals(acal.TypedValue[decimal.Decimal]) boolean.Simple
	NotEquals(acal.TypedValue[decimal.Decimal]) boolean.Simple
	LargerThan(acal.TypedValue[decimal.Decimal]) boolean.Simple
	LargerThanEquals(acal.TypedValue[decimal.Decimal]) boolean.Simple
	SmallerThan(acal.TypedValue[decimal.Decimal]) boolean.Simple
	SmallerThanEquals(acal.TypedValue[decimal.Decimal]) boolean.Simple
}
