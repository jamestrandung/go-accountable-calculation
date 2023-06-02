package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
)

// Simple ...
type Simple struct {
	*acal.Simple[bool]
}

// MakeSimple ...
func MakeSimple(name string, value bool) Simple {
	return Simple{
		Simple: acal.NewSimple(name, value),
	}
}

// MakeSimpleWithFormula returns a new Simple with the given value and formula.
func MakeSimpleWithFormula(value bool, formula *acal.SyntaxNode) Simple {
	return Simple{
		Simple: acal.NewSimpleWithFormula(value, formula),
	}
}

// MakeSimpleFrom returns a new Simple using the given value as formula.
func MakeSimpleFrom(value acal.TypedValue[bool]) Simple {
	return Simple{
		Simple: acal.NewSimpleFrom(value),
	}
}

// GetTypedValue returns the typed value this Simple contains.
func (s Simple) GetTypedValue() bool {
	return acal.ExtractTypedValue[bool](s.Simple)
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
		!nextOp.IsTransparent() && !lastOp.IsTransparent() && nextOp != lastOp,
	)
}

var toBaseSyntaxOperand = func(s Simple, nextOp acal.Op) *acal.SyntaxOperand {
	return s.Simple.ToSyntaxOperand(nextOp)
}

// SelfReplaceIfNil returns the replacement to represent this Simple if it is nil.
func (s Simple) SelfReplaceIfNil() acal.Value {
	if s.IsNil() {
		return NilBool
	}

	return s
}

// Anchor updates the name of this Simple to the provided string.
func (s Simple) Anchor(name string) Simple {
	if s.IsNil() {
		return MakeSimple(name, false)
	}

	anchored, isNew := s.Simple.DoAnchor(name)
	if !isNew {
		return s
	}

	return Simple{
		Simple: anchored,
	}
}

// Bool returns the value of this Simple as a bool.
// If it's nil, false is returned.
func (s Simple) Bool() bool {
	return acal.ExtractTypedValue[bool](s.Simple)
}
