package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
)

const (
	opAnd acal.Op = iota
	opOr
)

// If executes the provided doIfFn if the given BoolI is true and
// return whether the function was executed
var If = func(b acal.TypedValue[bool], doIfFn func(criteria acal.TypedValue[bool])) bool {
	b = BoolOps.ReplaceNilBoolI(b)
	if b.Bool() {
		doIfFn(b)
		return true
	}

	return false
}

// IfNot executes the provided doIfNotFn if the given BoolI is false
// and return whether the function was executed
var IfNot = func(b BoolI, doIfNotFn func(criteria BoolI)) bool {
	b = BoolOps.ReplaceNilBoolI(b)
	if !b.Bool() {
		doIfNotFn(b.Not())
		return true
	}

	return false
}

// IfElse executes the provided doIfFn if the given BoolI is true. Otherwise,
// doElseFn is executed instead. Finally, the value of BoolI will be returned.
var IfElse = func(b BoolI, doIfFn func(criteria BoolI), doElseFn func(elseCriteria BoolI)) bool {
	b = BoolOps.ReplaceNilBoolI(b)
	if b.Bool() {
		doIfFn(b)
		return true
	}

	doElseFn(b.Not())
	return false
}
