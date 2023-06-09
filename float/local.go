package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/shopspring/decimal"
)

// Local ...
type Local struct {
	*acal.Local[decimal.Decimal]
	opProvider
}

// MakeLocal returns a new Local with the provided fields.
func MakeLocal(name string, original acal.TypedValue[decimal.Decimal]) Local {
	core := acal.NewLocal(name, original)

	l := Local{
		Local: core,
		opProvider: opProvider{
			tv: core,
		},
	}

	l.WithFormatFn(FormatFn)

	return l
}

// GetTypedValue returns the typed value this Local contains.
func (l Local) GetTypedValue() decimal.Decimal {
	return acal.ExtractTypedValue[decimal.Decimal](l.Local)
}

// SelfReplaceIfNil returns the replacement to represent this Local if it is nil.
func (l Local) SelfReplaceIfNil() acal.Value {
	if l.IsNil() {
		return NilFloat
	}

	return l
}

// Anchor returns a new Simple initialized to the value of this
// Local and anchored with the given name.
func (l Local) Anchor(name string) Simple {
	return Simple{
		Simple: l.DoAnchor(name),
	}
}

// Decimal returns the value of this Local as a decimal.Decimal.
// If it's nil, a decimal.Decimal value of 0 is returned.
func (l Local) Decimal() decimal.Decimal {
	return acal.ExtractTypedValue[decimal.Decimal](l.Local)
}

// Float returns the value of this Local as a float64.
// If it's nil, 0 is returned.
func (l Local) Float() float64 {
	return l.Decimal().InexactFloat64()
}
