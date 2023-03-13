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
func NewSimpleWithFormula(value bool, formulaFn func() *acal.SyntaxNode) *Simple {
	return &Simple{
		Simple: acal.NewSimpleWithFormula(value, formulaFn),
	}
}

// SelfReplaceIfNil returns the replacement to represent this Simple if it is nil.
func (s *Simple) SelfReplaceIfNil() acal.Value {
	if s.IsNil() {
		return acal.ZeroSimple[bool]("NilBool")
	}

	return s
}
