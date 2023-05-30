package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple_And(t *testing.T) {
	tSimple1 := NewSimple("True1", true)
	tSimple2 := NewSimple("True2", true)
	fSimple := NewSimple("False", false)

	and1 := tSimple1.And(tSimple2)

	assert.True(t, and1.Bool())
	assert.Equal(
		t,
		`{"True1":{"Value":"true","Source":"unknown"},"True2":{"Value":"true","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"AND","Operands":[{"Name":"True1"},{"Name":"True2"}]}}}`,
		acal.ToString(and1),
	)

	and2 := tSimple1.And(fSimple)

	assert.False(t, and2.Bool())
	assert.Equal(
		t,
		`{"False":{"Value":"false","Source":"unknown"},"True1":{"Value":"true","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"AND","Operands":[{"Name":"True1"},{"Name":"False"}]}}}`,
		acal.ToString(and2),
	)
}

func TestSimple_Or(t *testing.T) {
	tSimple1 := NewSimple("True1", true)
	tSimple2 := NewSimple("True2", true)
	fSimple := NewSimple("False", false)

	or1 := tSimple1.Or(tSimple2)

	assert.True(t, or1.Bool())
	assert.Equal(
		t,
		`{"True1":{"Value":"true","Source":"unknown"},"True2":{"Value":"true","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"OR","Operands":[{"Name":"True1"},{"Name":"True2"}]}}}`,
		acal.ToString(or1),
	)

	or2 := tSimple1.Or(fSimple)

	assert.True(t, or2.Bool())
	assert.Equal(
		t,
		`{"False":{"Value":"false","Source":"unknown"},"True1":{"Value":"true","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"OR","Operands":[{"Name":"True1"},{"Name":"False"}]}}}`,
		acal.ToString(or2),
	)
}

func TestSimple_Not(t *testing.T) {
	tSimple := NewSimple("True1", true)
	fSimple := NewSimple("False", false)

	not1 := tSimple.Not()

	assert.False(t, not1.Bool())
	assert.Equal(
		t,
		`{"True1":{"Value":"true","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"NOT","Operands":[{"Name":"True1"}]}}}`,
		acal.ToString(not1),
	)

	not2 := fSimple.Not()

	assert.True(t, not2.Bool())
	assert.Equal(
		t,
		`{"False":{"Value":"false","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"NOT","Operands":[{"Name":"False"}]}}}`,
		acal.ToString(not2),
	)
}
