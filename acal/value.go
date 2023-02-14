package acal

const (
	UnknownValueName = "Unknown"
)

// Value is a wrapper that gives an identity to any variables that are used in our code.
//go:generate mockery --name=Value --case underscore --inpackage
type Value interface {
	// IsNil returns whether this Value is nil.
	IsNil() bool
	// GetName return the name of this Value.
	GetName() string
	// GetAlias return the alias of this Value.
	GetAlias() string
	// SetAlias updates the alias of this Value to the provided string.
	SetAlias(string)
	// GetValue returns the untyped value this Value contains.
	GetValue() any
	// ToSyntaxOperand returns the SyntaxOperand representation of this Value.
	ToSyntaxOperand(nextOp Op) *SyntaxOperand
	// ExtractValues extracts this Value and all Value that were used to calculate it.
	ExtractValues(cache IValueCache) IValueCache
	// SelfReplaceIfNil returns the replacement to represent this Value if it is nil.
	SelfReplaceIfNil() Value
}

// TypedValue is a Value bounded by a particular type.
//go:generate mockery --name=TypedValue --case underscore --inpackage
type TypedValue[T any] interface {
	Value
	// GetTypedValue returns the typed value this TypedValue contains.
	GetTypedValue() T
}
