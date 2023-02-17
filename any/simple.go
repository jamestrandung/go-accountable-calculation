package any

import "github.com/jamestrandung/go-accountable-calculation/acal"

// Simple ...
type Simple[T any] struct {
	*acal.Simple[T]
}

// NewSimple ...
func NewSimple[T any](name string, value T) *Simple[T] {
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
