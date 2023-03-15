package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
)

// Simple ...
type Simple struct {
	*acal.Simple[bool]
}

// Bool ...
func New(name string, value bool) *Simple {
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

// GetTypedValue returns the typed value this Simple contains.
func (s *Simple) GetTypedValue() bool {
	return s.Simple.GetTypedValue()
}

// ToSyntaxOperand returns the acal.SyntaxOperand representation of this Simple.
func (s *Simple) ToSyntaxOperand(nextOp acal.Op) *acal.SyntaxOperand {
	if result := s.Simple.ToSyntaxOperand(nextOp); result != nil {
		return result
	}

	formula := s.GetFormulaFn()()
	lastOp := formula.GetOp()

	return acal.NewSyntaxOperandWithFormula(
		formula,
		!nextOp.IsTransparent() && !lastOp.IsTransparent() && nextOp != lastOp,
	)
}

// SelfReplaceIfNil returns the replacement to represent this Simple if it is nil.
func (s *Simple) SelfReplaceIfNil() acal.Value {
	if s.IsNil() {
		return acal.ZeroSimple[bool]("NilBool")
	}

	return s
}

// Bool returns the value of this Simple as a bool.
// If it's nil, false is returned.
func (s *Simple) Bool() bool {
	return s.GetTypedValue()
}
