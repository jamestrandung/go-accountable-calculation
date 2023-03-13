package anything

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
)

// Value represents an acal.Value of any kinds.
type Value interface {
	acal.Value
	// String returns the value of this Value as a string.
	// If it's nil, an empty string is returned.
	String() string
}

//go:generate mockery --name=Interface --case underscore --inpkg
// Interface governs the methods that Value should provide.
type Interface[T any] interface {
	Value
	// EqualsRaw returns whether the value of this Interface equals to the raw input value.
	EqualsRaw(v T) *boolean.Simple
	// Equals returns whether the value of this Interface equals to the input value.
	Equals(T) *boolean.Simple
	// NotEqualsRaw returns whether the value of this Interface does not equal to the raw input value.
	NotEqualsRaw(v T) *boolean.Simple
	// NotEquals returns whether the value of this Interface does not equal to the input value.
	NotEquals(T) *boolean.Simple
}
