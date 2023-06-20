package akin

import (
    "github.com/jamestrandung/go-accountable-calculation/acal"
    "github.com/jamestrandung/go-accountable-calculation/boolean"
)

// Value represents an acal.Value of any kind.
//
//go:generate mockery --name=Value --case underscore --inpackage
type Value[T any] interface {
    acal.TypedValue[T]

    // Comparable returns the comparable value this Value contains.
    Comparable() T
    // EqualsRaw returns whether the value of this Value equals to the raw input value.
    EqualsRaw(v T) boolean.Simple
    // Equals returns whether the value of this Value equals to the input value.
    Equals(T) boolean.Simple
    // NotEqualsRaw returns whether the value of this Value does not equal to the raw input value.
    NotEqualsRaw(v T) boolean.Simple
    // NotEquals returns whether the value of this Value does not equal to the input value.
    NotEquals(T) boolean.Simple
}
