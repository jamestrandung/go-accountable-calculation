package anything

import (
	"fmt"
	"github.com/jamestrandung/go-accountable-calculation/acal"
)

// Simple ...
type Simple[T comparable] struct {
	*acal.Simple[T]
}

// NewSimple ...
func NewSimple[T comparable](name string, value T) *Simple[T] {
	return &Simple[T]{
		Simple: acal.NewSimple(name, value),
	}
}

// SelfReplaceIfNil returns the replacement to represent this Simple if it is nil.
func (s *Simple[T]) SelfReplaceIfNil() acal.Value {
	if s.IsNil() {
		return acal.ZeroSimple[T]("NilAny")
	}

	return s
}

// String returns the value of this Simple as a string.
// If it's nil, an empty string is returned.
func (s *Simple[T]) String() string {
	if s.IsNil() {
		return ""
	}

	return fmt.Sprintf("%v", s.GetTypedValue())
}
