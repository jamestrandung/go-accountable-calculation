package akin

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
)

// Simple ...
type Simple[T comparable] struct {
	*acal.Simple[T]
}

// MakeSimple ...
func MakeSimple[T comparable](name string, value T) Simple[T] {
	return Simple[T]{
		Simple: acal.NewSimple(name, value),
	}
}

// MakeSimpleFrom returns a new Simple using the given value as formula.
func MakeSimpleFrom[T comparable](value acal.TypedValue[T]) Simple[T] {
	return Simple[T]{
		Simple: acal.NewSimpleFrom(value),
	}
}

// Anchor updates the name of this Simple to the provided string.
func (s Simple[T]) Anchor(name string) Simple[T] {
	if s.IsNil() {
		var temp T
		return MakeSimple(name, temp)
	}

	anchored, isNew := s.Simple.DoAnchor(name)
	if !isNew {
		return s
	}

	return Simple[T]{
		Simple: anchored,
	}
}

// Comparable returns the comparable value this Simple contains.
func (s Simple[T]) Comparable() T {
	return acal.ExtractTypedValue[T](s)
}
