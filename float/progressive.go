package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/my-shopspring/decimal"
)

// Progressive ...
type Progressive struct {
	*acal.Progressive[decimal.Decimal]
	opProvider
}

// MakeProgressive returns a new Progressive with the provided fields.
func MakeProgressive(name string) Progressive {
	core := acal.NewProgressive[decimal.Decimal](name)

	p := Progressive{
		Progressive: core,
	}

	p.opProvider = opProvider{
		tv: p,
	}

	p.WithFormatFn(FormatFn)

	return p
}

// GetTypedValue returns the typed value this Progressive contains.
func (p Progressive) GetTypedValue() decimal.Decimal {
	return acal.ExtractTypedValue[decimal.Decimal](p.Progressive)
}

// Anchor returns a new Simple initialized to the value of this
// Progressive and anchored with the given name.
func (p Progressive) Anchor(name string) Simple {
	return Simple{
		Simple: p.DoAnchor(name),
	}
}

// Decimal returns the value of this Progressive as a decimal.Decimal.
// If it's nil, a decimal.Decimal value of 0 is returned.
func (p Progressive) Decimal() decimal.Decimal {
	return acal.ExtractTypedValue[decimal.Decimal](p.Progressive)
}

// Float returns the value of this Progressive as a float64.
// If it's nil, 0 is returned.
func (p Progressive) Float() float64 {
	return p.Decimal().InexactFloat64()
}
