package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/shopspring/decimal"
)

// Simple ...
type Simple struct {
	*acal.Simple[decimal.Decimal]
}

// MakeSimpleFromFloat ...
func MakeSimpleFromFloat(name string, value float64) Simple {
	s := Simple{
		Simple: acal.NewSimple(name, decimal.NewFromFloat(value)),
	}

	s.WithFormatFn(floatFormatFn)

	return s
}

// MakeSimpleFromFloat32 ...
func MakeSimpleFromFloat32(name string, value float32) Simple {
	s := Simple{
		Simple: acal.NewSimple(name, decimal.NewFromFloat32(value)),
	}

	s.WithFormatFn(floatFormatFn)

	return s
}

// MakeSimpleFromInt ...
func MakeSimpleFromInt(name string, value int64) Simple {
	s := Simple{
		Simple: acal.NewSimple(name, decimal.NewFromInt(value)),
	}

	s.WithFormatFn(floatFormatFn)

	return s
}

// MakeSimpleFromInt32 ...
func MakeSimpleFromInt32(name string, value int32) Simple {
	s := Simple{
		Simple: acal.NewSimple(name, decimal.NewFromInt32(value)),
	}

	s.WithFormatFn(floatFormatFn)

	return s
}

// MakeSimpleFromDecimal ...
func MakeSimpleFromDecimal(name string, value decimal.Decimal) Simple {
	s := Simple{
		Simple: acal.NewSimple(name, value),
	}

	s.WithFormatFn(floatFormatFn)

	return s
}

// MakeSimpleWithFormula returns a new Simple with the given value and formula.
func MakeSimpleWithFormula(value decimal.Decimal, formula *acal.SyntaxNode) Simple {
	s := Simple{
		Simple: acal.NewSimpleWithFormula(value, formula),
	}

	s.WithFormatFn(floatFormatFn)

	return s
}

// MakeSimpleFrom returns a new Simple using the given value as formula.
func MakeSimpleFrom(value acal.TypedValue[decimal.Decimal]) Simple {
	s := Simple{
		Simple: acal.NewSimpleFrom(value),
	}

	s.WithFormatFn(floatFormatFn)

	return s
}

// GetTypedValue returns the typed value this Simple contains.
func (s Simple) GetTypedValue() decimal.Decimal {
	return acal.ExtractTypedValue[decimal.Decimal](s.Simple)
}

// ToSyntaxOperand returns the acal.SyntaxOperand representation of this Simple.
func (s Simple) ToSyntaxOperand(nextOp acal.Op) *acal.SyntaxOperand {
	if result := toBaseSyntaxOperand(s, nextOp); result != nil {
		return result
	}

	formula := s.GetFormula()
	lastOp := formula.GetOp()

	return acal.NewSyntaxOperandWithFormula(
		formula,
		!nextOp.IsTransparent() && !lastOp.IsTransparent() && opLevel[lastOp] < opLevel[nextOp],
	)
}

var toBaseSyntaxOperand = func(s Simple, nextOp acal.Op) *acal.SyntaxOperand {
	return s.Simple.ToSyntaxOperand(nextOp)
}

// SelfReplaceIfNil returns the replacement to represent this Simple if it is nil.
func (s Simple) SelfReplaceIfNil() acal.Value {
	if s.IsNil() {
		return NilFloat
	}

	return s
}

// Anchor updates the name of this Simple to the provided string.
func (s Simple) Anchor(name string) Simple {
	if s.IsNil() {
		return MakeSimpleFromFloat(name, 0)
	}

	anchored, isNew := s.Simple.DoAnchor(name)
	if !isNew {
		return s
	}

	return Simple{
		Simple: anchored,
	}
}

// Decimal returns the value of this Simple as a decimal.Decimal.
// If it's nil, a decimal.Decimal value of 0 is returned.
func (s Simple) Decimal() decimal.Decimal {
	return acal.ExtractTypedValue[decimal.Decimal](s.Simple)
}

// Float returns the value of this Simple as a float64.
// If it's nil, 0 is returned.
func (s Simple) Float() float64 {
	return s.Decimal().InexactFloat64()
}
