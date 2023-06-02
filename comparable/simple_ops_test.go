package comparable

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple_EqualsRaw(t *testing.T) {
	simple := MakeSimple("Simple", "string")

	e1 := simple.EqualsRaw("string")

	assert.True(t, e1.Bool())
	assert.Equal(
		t,
		`{"Simple":{"Value":"string","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"==","Operands":[{"Name":"Simple"},{"StaticValue":"string"}]}}}`,
		acal.ToString(e1),
	)

	e2 := simple.EqualsRaw("text")

	assert.False(t, e2.Bool())
	assert.Equal(
		t,
		`{"Simple":{"Value":"string","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"==","Operands":[{"Name":"Simple"},{"StaticValue":"text"}]}}}`,
		acal.ToString(e2),
	)
}

func TestSimple_Equals(t *testing.T) {
	simple := MakeSimple("Simple", "string")

	e1 := simple.Equals(MakeSimple("String", "string"))

	assert.True(t, e1.Bool())
	assert.Equal(
		t,
		`{"Simple":{"Value":"string","Source":"unknown"},"String":{"Value":"string","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"==","Operands":[{"Name":"Simple"},{"Name":"String"}]}}}`,
		acal.ToString(e1),
	)

	e2 := simple.Equals(MakeSimple("Text", "text"))

	assert.False(t, e2.Bool())
	assert.Equal(
		t,
		`{"Simple":{"Value":"string","Source":"unknown"},"Text":{"Value":"text","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"==","Operands":[{"Name":"Simple"},{"Name":"Text"}]}}}`,
		acal.ToString(e2),
	)
}

func TestSimple_NotEqualsRaw(t *testing.T) {
	simple := MakeSimple("Simple", "string")

	e1 := simple.NotEqualsRaw("string")

	assert.False(t, e1.Bool())
	assert.Equal(
		t,
		`{"Simple":{"Value":"string","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"!=","Operands":[{"Name":"Simple"},{"StaticValue":"string"}]}}}`,
		acal.ToString(e1),
	)

	e2 := simple.NotEqualsRaw("text")

	assert.True(t, e2.Bool())
	assert.Equal(
		t,
		`{"Simple":{"Value":"string","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"!=","Operands":[{"Name":"Simple"},{"StaticValue":"text"}]}}}`,
		acal.ToString(e2),
	)
}

func TestSimple_NotEquals(t *testing.T) {
	simple := MakeSimple("Simple", "string")

	e1 := simple.NotEquals(MakeSimple("String", "string"))

	assert.False(t, e1.Bool())
	assert.Equal(
		t,
		`{"Simple":{"Value":"string","Source":"unknown"},"String":{"Value":"string","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"!=","Operands":[{"Name":"Simple"},{"Name":"String"}]}}}`,
		acal.ToString(e1),
	)

	e2 := simple.NotEquals(MakeSimple("Text", "text"))

	assert.True(t, e2.Bool())
	assert.Equal(
		t,
		`{"Simple":{"Value":"string","Source":"unknown"},"Text":{"Value":"text","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"!=","Operands":[{"Name":"Simple"},{"Name":"Text"}]}}}`,
		acal.ToString(e2),
	)
}
