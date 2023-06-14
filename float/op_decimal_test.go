package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpProvider_Plus(t *testing.T) {
	op := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}
	tv := MakeSimpleFromInt("Simple2", 2)

	actual := op.Add(tv)

	assert.Equal(t, decimal.NewFromInt(3), actual.Decimal())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Unknown":{"Value":"3","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"+","Operands":[{"Name":"Simple1"},{"Name":"Simple2"}]}}}`,
		acal.ToString(actual),
	)
}

func TestOpProvider_Minus(t *testing.T) {
	op := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}
	tv := MakeSimpleFromInt("Simple2", 2)

	actual := op.Sub(tv)

	assert.Equal(t, decimal.NewFromInt(-1), actual.Decimal())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Unknown":{"Value":"-1","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"-","Operands":[{"Name":"Simple1"},{"Name":"Simple2"}]}}}`,
		acal.ToString(actual),
	)
}

func TestOpProvider_Multiply(t *testing.T) {
	op := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}
	tv := MakeSimpleFromInt("Simple2", 2)

	actual := op.Mul(tv)

	assert.Equal(t, decimal.NewFromInt(2), actual.Decimal())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Unknown":{"Value":"2","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"*","Operands":[{"Name":"Simple1"},{"Name":"Simple2"}]}}}`,
		acal.ToString(actual),
	)
}

func TestOpProvider_Divide(t *testing.T) {
	op := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}
	tv := MakeSimpleFromInt("Simple2", 2)

	actual := op.Div(tv)

	assert.Equal(t, 0.5, actual.Float())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Unknown":{"Value":"0.5","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"/","Operands":[{"Name":"Simple1"},{"Name":"Simple2"}]}}}`,
		acal.ToString(actual),
	)
}
