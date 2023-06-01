package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
)

// Simple ...
type Simple struct {
	*acal.Simple[bool]
}

// NewSimple ...
func NewSimple(name string, value bool) *Simple {
	return &Simple{
		Simple: acal.NewSimple(name, value),
	}
}

// NewSimpleWithFormula returns a new Simple with the given value and formula.
func NewSimpleWithFormula(value bool, formula *acal.SyntaxNode) *Simple {
	return &Simple{
		Simple: acal.NewSimpleWithFormula(value, formula),
	}
}

// NewSimpleFrom returns a new Simple using the given value as formula.
func NewSimpleFrom(value acal.TypedValue[bool]) *Simple {
	return &Simple{
		Simple: acal.NewSimpleFrom(value),
	}
}

// IsNil returns whether this Simple is nil.
func (s *Simple) IsNil() bool {
	return s == nil || s.Simple.IsNil()
}

// GetTypedValue returns the typed value this Simple contains.
func (s *Simple) GetTypedValue() bool {
	if s == nil || s.IsNil() {
		return false
	}

	return s.Simple.GetTypedValue()
}

// ToSyntaxOperand returns the acal.SyntaxOperand representation of this Simple.
func (s *Simple) ToSyntaxOperand(nextOp acal.Op) *acal.SyntaxOperand {
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

var toBaseSyntaxOperand = func(s *Simple, nextOp acal.Op) *acal.SyntaxOperand {
	return s.Simple.ToSyntaxOperand(nextOp)
}

// SelfReplaceIfNil returns the replacement to represent this Simple if it is nil.
func (s *Simple) SelfReplaceIfNil() acal.Value {
	if s == nil || s.IsNil() {
		return NilBool
	}

	return s
}

// Anchor updates the name of this Simple to the provided string.
func (s *Simple) Anchor(name string) *Simple {
	if s.IsNil() {
		return NewSimple(name, false)
	}

	anchored, isNew := s.Simple.DoAnchor(name)
	if !isNew {
		return s
	}

	return &Simple{
		Simple: anchored,
	}
}

// Bool returns the value of this Simple as a bool.
// If it's nil, false is returned.
func (s *Simple) Bool() bool {
	return s.GetTypedValue()
}
