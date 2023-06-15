package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
)

const (
	opAnd acal.Op = iota
	opOr
)

// If executes the provided doIfFn if the given Interface is true and
// return whether the function was executed.
func If(b Interface, doIfFn func(criteria Interface)) bool {
	if b == nil || b.IsNil() {
		b = NilBool
	}

	if b.Bool() {
		doIfFn(b)
		return true
	}

	return false
}

// IfNot executes the provided doIfNotFn if the given Interface is false
// and return whether the function was executed.
func IfNot(b Interface, doIfNotFn func(criteria Interface)) bool {
	if b == nil || b.IsNil() {
		b = NilBool
	}

	if !b.Bool() {
		doIfNotFn(b.Not())
		return true
	}

	return false
}

// IfElse executes the provided doIfFn if the given Interface is true. Otherwise,
// doElseFn is executed instead. Finally, the value of Interface will be returned.
func IfElse(
	b Interface,
	doIfFn func(criteria Interface),
	doElseFn func(elseCriteria Interface),
) bool {
	if b == nil || b.IsNil() {
		b = NilBool
	}

	if b.Bool() {
		doIfFn(b)
		return true
	}

	doElseFn(b.Not())
	return false
}
