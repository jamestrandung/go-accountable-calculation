package comparable

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
)

// Interface governs the methods that Value should provide.
//
//go:generate mockery --name=Interface --case underscore --inpackage
type Interface[T any] interface {
	acal.Value
	// EqualsRaw returns whether the value of this Interface equals to the raw input value.
	EqualsRaw(v T) *boolean.Simple
	// Equals returns whether the value of this Interface equals to the input value.
	Equals(T) *boolean.Simple
	// NotEqualsRaw returns whether the value of this Interface does not equal to the raw input value.
	NotEqualsRaw(v T) *boolean.Simple
	// NotEquals returns whether the value of this Interface does not equal to the input value.
	NotEquals(T) *boolean.Simple
}
