package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/shopspring/decimal"
)

// Remote ...
type Remote struct {
	*acal.Remote[decimal.Decimal]
	opProvider
}

// MakeRemote ...
func MakeRemote(name string, value decimal.Decimal, remoteFieldName string, remoteLogKey string) Remote {
	core := acal.NewRemote(name, value, remoteFieldName, remoteLogKey)

	r := Remote{
		Remote: core,
		opProvider: opProvider{
			tv: core,
		},
	}

	r.WithFormatFn(FormatFn)

	return r
}

// GetTypedValue returns the typed value this Remote contains.
func (r Remote) GetTypedValue() decimal.Decimal {
	return acal.ExtractTypedValue[decimal.Decimal](r.Remote)
}

// SelfReplaceIfNil returns the replacement to represent this Remote if it is nil.
func (r Remote) SelfReplaceIfNil() acal.Value {
	if r.IsNil() {
		return NilFloat
	}

	return r
}

// Anchor returns a new Simple initialized to the value of this
// Remote and anchored with the given name.
func (r Remote) Anchor(name string) Simple {
	return Simple{
		Simple: r.DoAnchor(name),
	}
}

// Decimal returns the value of this Remote as a decimal.Decimal.
// If it's nil, a decimal.Decimal value of 0 is returned.
func (r Remote) Decimal() decimal.Decimal {
	return acal.ExtractTypedValue[decimal.Decimal](r.Remote)
}

// Float returns the value of this Remote as a float64.
// If it's nil, 0 is returned.
func (r Remote) Float() float64 {
	return r.Decimal().InexactFloat64()
}
