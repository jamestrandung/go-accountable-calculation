package boolean

import "github.com/jamestrandung/go-accountable-calculation/acal"

// Value represents an acal.Value of boolean kind.
//
//go:generate mockery --name=Value --case underscore --inpackage
type Value interface {
	acal.TypedValue[bool]
	// Bool returns the value of this Value as a bool.
	// If it's nil, false is returned.
	Bool() bool

	// And applies AND operation on the value of this Interface and the given acal.TypedValue.
	And(acal.TypedValue[bool]) Simple
	// Or applies OR operation on the value of this Interface and the given acal.TypedValue.
	Or(acal.TypedValue[bool]) Simple
	// Not returns the inverse value of this Interface.
	Not() Simple
}
