package acal

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
	iNamedValue

	// IsNil returns whether this Value is nil.
	IsNil() bool
	// GetValue returns the untyped value this Value contains.
	GetValue() any
	// Stringify returns the value this Value contains as a string.
	Stringify() string
	// ToSyntaxOperand returns the SyntaxOperand representation of this Value.
	ToSyntaxOperand(nextOp Op) *SyntaxOperand
	// ExtractValues extracts this Value and all Value that were used to calculate it.
	ExtractValues(cache IValueCache) IValueCache
	// SelfReplaceIfNil returns the replacement to represent this Value if it is nil.
	SelfReplaceIfNil() Value
}

type iNamedValue interface {
	// GetName return the name of this iNamedValue.
	GetName() string
	// GetAlias return the alias of this iNamedValue.
	GetAlias() string
	// SetAlias updates the alias of this iNamedValue to the provided string.
	SetAlias(string)
	// HasIdentity returns whether this iNamedValue was given an identity.
	HasIdentity() bool
	// Identify returns the identity that was given to this iNamedValue.
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
