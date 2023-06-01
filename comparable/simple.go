package comparable

import (
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

// NewSimpleFrom returns a new Simple using the given value as formula.
func NewSimpleFrom[T comparable](value acal.TypedValue[T]) *Simple[T] {
	return &Simple[T]{
		Simple: acal.NewSimpleFrom(value),
	}
}

// IsNil returns whether this Simple is nil.
func (s *Simple[T]) IsNil() bool {
	return s == nil || s.Simple.IsNil()
}

// GetTypedValue returns the typed value this Simple contains.
func (s *Simple[T]) GetTypedValue() T {
	if s == nil || s.IsNil() {
		var temp T
		return temp
	}

	return s.Simple.GetTypedValue()
}

// SelfReplaceIfNil returns the replacement to represent this Simple if it is nil.
func (s *Simple[T]) SelfReplaceIfNil() acal.Value {
	if s == nil || s.IsNil() {
		return acal.ZeroSimple[T]("NilComparable")
	}

	return s
}

// Anchor updates the name of this Simple to the provided string.
func (s *Simple[T]) Anchor(name string) *Simple[T] {
	if s.IsNil() {
		var temp T
		return NewSimple(name, temp)
	}

	anchored, isNew := s.Simple.DoAnchor(name)
	if !isNew {
		return s
	}

	return &Simple[T]{
		Simple: anchored,
	}
}
