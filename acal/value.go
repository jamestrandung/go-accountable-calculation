package acal

const (
    UnknownValueName = "Unknown"
)

type iNamedValue interface {
    // GetName return the name of this iNamedValue.
    GetName() string
    // GetAlias return the alias of this iNamedValue.
    GetAlias() string
    // SetAlias updates the alias of this iNamedValue to the provided string.
    SetAlias(string)
}

type namedValue struct {
    name  string
    alias string
}

// GetName returns the name of this namedValue.
func (v *namedValue) GetName() string {
    return v.name
}

// GetAlias returns the alias of this namedValue.
func (v *namedValue) GetAlias() string {
    return v.alias
}

// SetAlias updates the alias of this namedValue to the provided string.
func (v *namedValue) SetAlias(alias string) {
    v.alias = alias
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
    // ToSyntaxOperand returns the SyntaxOperand representation of this Value.
    ToSyntaxOperand(nextOp Op) *SyntaxOperand
    // ExtractValues extracts this Value and all Value that were used to calculate it.
    ExtractValues(cache IValueCache) IValueCache
    // SelfReplaceIfNil returns the replacement to represent this Value if it is nil.
    SelfReplaceIfNil() Value
}

// TypedValue is a Value bounded by a particular type.
//
//go:generate mockery --name=TypedValue --case underscore --inpackage
type TypedValue[T any] interface {
    Value
    // GetTypedValue returns the typed value this TypedValue contains.
    GetTypedValue() T
}
