package acal

import "fmt"

const (
	UnknownValueName = "Unknown"
)

// TypedValue is a Value bounded by a particular type.
//
//go:generate mockery --name=TypedValue --case underscore --inpackage
type TypedValue[T any] interface {
	Value
	// GetTypedValue returns the typed value this TypedValue contains.
	GetTypedValue() T
}

// Value is a wrapper that gives an identity to any variables that are used in our code.
//
//go:generate mockery --name=Value --case underscore --inpackage
type Value interface {
	identifiable
	extractable
	syntaxOperandProvider

	// IsNil returns whether this Value is nil.
	IsNil() bool
	// GetValue returns the untyped value this Value contains.
	GetValue() any
	// Stringify returns the value this Value contains as a string.
	Stringify() string
}

type identifiable interface {
	// GetName return the name of this identifiable.
	GetName() string
	// GetAlias return the alias of this identifiable.
	GetAlias() string
	// SetAlias updates the alias of this identifiable to the provided string.
	SetAlias(string)
	// HasIdentity returns whether this identifiable was given an identity.
	HasIdentity() bool
	// Identify returns the identity that was given to this identifiable.
	Identify() string
}

type namedValue struct {
	name  string
	alias string
}

// GetName ...
func (v *namedValue) GetName() string {
	return v.name
}

// GetAlias ...
func (v *namedValue) GetAlias() string {
	return v.alias
}

// SetAlias ...
func (v *namedValue) SetAlias(alias string) {
	v.alias = alias
}

// HasIdentity ...
func (v *namedValue) HasIdentity() bool {
	return v.GetName() != "" || v.GetAlias() != ""
}

// Identify ...
func (v *namedValue) Identify() string {
	if v.GetAlias() != "" {
		return v.GetAlias()
	}

	return v.GetName()
}

type valueFormatter[T any] struct {
	formatFn func(T) string
}

func (f *valueFormatter[T]) formatValue(v T) string {
	if f.formatFn == nil {
		return fmt.Sprintf("%v", v)
	}

	return f.formatFn(v)
}

func (f *valueFormatter[T]) GetFormatFn() func(T) string {
	return f.formatFn
}

func (f *valueFormatter[T]) WithFormatFn(formatFn func(T) string) {
	f.formatFn = formatFn
}
