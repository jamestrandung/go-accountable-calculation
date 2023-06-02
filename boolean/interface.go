package boolean

import "github.com/jamestrandung/go-accountable-calculation/acal"

// Value represents an acal.Value of boolean kind.
//
//go:generate mockery --name=Value --case underscore --inpackage
type Value interface {
	acal.Value
	// Bool returns the value of this Value as a bool.
	// If it's nil, false is returned.
	Bool() bool
}

// Interface governs the methods that Value should provide.
//
//go:generate mockery --name=Interface --case underscore --inpackage
type Interface interface {
	Value
	// And applies AND operation on the value of this Interface and the given acal.TypedValue.
	And(acal.TypedValue[bool]) Simple
	// Or applies OR operation on the value of this Interface and the given acal.TypedValue.
	Or(acal.TypedValue[bool]) Simple
	// Not returns the inverse value of this Interface.
	Not() Simple
	// Then does nothing and returns this Interface as-is. It's meant for separating code
	// into more readable chunk.
	Then() Interface
}
