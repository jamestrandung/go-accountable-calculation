package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpProvider_IsZero(t *testing.T) {
	op1 := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}

	actual := op1.IsZero()

	assert.False(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"==","Operands":[{"Name":"Simple1"},{"StaticValue":"0"}]}}}`,
		acal.ToString(actual),
	)

	op2 := opProvider{tv: MakeSimpleFromInt("Simple2", 0)}

	actual = op2.IsZero()

	assert.True(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple2":{"Value":"0","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"==","Operands":[{"Name":"Simple2"},{"StaticValue":"0"}]}}}`,
		acal.ToString(actual),
	)
}

func TestOpProvider_IsPositive(t *testing.T) {
	op1 := opProvider{tv: MakeSimpleFromInt("Simple1", -1)}

	actual := op1.IsPositive()

	assert.False(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"-1","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003e","Operands":[{"Name":"Simple1"},{"StaticValue":"0"}]}}}`,
		acal.ToString(actual),
	)

	op2 := opProvider{tv: MakeSimpleFromInt("Simple2", 0)}

	actual = op2.IsPositive()

	assert.False(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple2":{"Value":"0","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003e","Operands":[{"Name":"Simple2"},{"StaticValue":"0"}]}}}`,
		acal.ToString(actual),
	)

	op3 := opProvider{tv: MakeSimpleFromInt("Simple3", 1)}

	actual = op3.IsPositive()

	assert.True(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple3":{"Value":"1","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003e","Operands":[{"Name":"Simple3"},{"StaticValue":"0"}]}}}`,
		acal.ToString(actual),
	)
}

func TestOpProvider_IsNegative(t *testing.T) {
	op1 := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}

	actual := op1.IsNegative()

	assert.False(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003c","Operands":[{"Name":"Simple1"},{"StaticValue":"0"}]}}}`,
		acal.ToString(actual),
	)

	op2 := opProvider{tv: MakeSimpleFromInt("Simple2", 0)}

	actual = op2.IsNegative()

	assert.False(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple2":{"Value":"0","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003c","Operands":[{"Name":"Simple2"},{"StaticValue":"0"}]}}}`,
		acal.ToString(actual),
	)

	op3 := opProvider{tv: MakeSimpleFromInt("Simple3", -1)}

	actual = op3.IsNegative()

	assert.True(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple3":{"Value":"-1","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003c","Operands":[{"Name":"Simple3"},{"StaticValue":"0"}]}}}`,
		acal.ToString(actual),
	)
}

func TestOpProvider_Equals(t *testing.T) {
	op := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}
	tv1 := MakeSimpleFromInt("Simple2", 2)

	actual := op.Equals(tv1)

	assert.False(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"==","Operands":[{"Name":"Simple1"},{"Name":"Simple2"}]}}}`,
		acal.ToString(actual),
	)

	tv2 := MakeSimpleFromInt("Simple3", 1)

	actual = op.Equals(tv2)

	assert.True(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple3":{"Value":"1","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"==","Operands":[{"Name":"Simple1"},{"Name":"Simple3"}]}}}`,
		acal.ToString(actual),
	)
}

func TestOpProvider_NotEquals(t *testing.T) {
	op := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}
	tv1 := MakeSimpleFromInt("Simple2", 2)

	actual := op.NotEquals(tv1)

	assert.True(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"!=","Operands":[{"Name":"Simple1"},{"Name":"Simple2"}]}}}`,
		acal.ToString(actual),
	)

	tv2 := MakeSimpleFromInt("Simple3", 1)

	actual = op.NotEquals(tv2)

	assert.False(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple3":{"Value":"1","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"!=","Operands":[{"Name":"Simple1"},{"Name":"Simple3"}]}}}`,
		acal.ToString(actual),
	)
}

func TestOpProvider_LargerThan(t *testing.T) {
	op := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}
	tv1 := MakeSimpleFromInt("Simple2", 0)

	actual := op.LargerThan(tv1)

	assert.True(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"0","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003e","Operands":[{"Name":"Simple1"},{"Name":"Simple2"}]}}}`,
		acal.ToString(actual),
	)

	tv2 := MakeSimpleFromInt("Simple3", 2)

	actual = op.LargerThan(tv2)

	assert.False(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple3":{"Value":"2","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003e","Operands":[{"Name":"Simple1"},{"Name":"Simple3"}]}}}`,
		acal.ToString(actual),
	)
}

func TestOpProvider_LargerThanEquals(t *testing.T) {
	op := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}
	tv1 := MakeSimpleFromInt("Simple2", 0)

	actual := op.LargerThanEquals(tv1)

	assert.True(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"0","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003e=","Operands":[{"Name":"Simple1"},{"Name":"Simple2"}]}}}`,
		acal.ToString(actual),
	)

	tv2 := MakeSimpleFromInt("Simple3", 1)

	actual = op.LargerThanEquals(tv2)

	assert.True(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple3":{"Value":"1","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003e=","Operands":[{"Name":"Simple1"},{"Name":"Simple3"}]}}}`,
		acal.ToString(actual),
	)

	tv3 := MakeSimpleFromInt("Simple4", 2)

	actual = op.LargerThanEquals(tv3)

	assert.False(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple4":{"Value":"2","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003e=","Operands":[{"Name":"Simple1"},{"Name":"Simple4"}]}}}`,
		acal.ToString(actual),
	)
}

func TestOpProvider_SmallerThan(t *testing.T) {
	op := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}
	tv1 := MakeSimpleFromInt("Simple2", 2)

	actual := op.SmallerThan(tv1)

	assert.True(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003c","Operands":[{"Name":"Simple1"},{"Name":"Simple2"}]}}}`,
		acal.ToString(actual),
	)

	tv2 := MakeSimpleFromInt("Simple3", 0)

	actual = op.SmallerThan(tv2)

	assert.False(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple3":{"Value":"0","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003c","Operands":[{"Name":"Simple1"},{"Name":"Simple3"}]}}}`,
		acal.ToString(actual),
	)
}

func TestOpProvider_SmallerThanEquals(t *testing.T) {
	op := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}
	tv1 := MakeSimpleFromInt("Simple2", 2)

	actual := op.SmallerThanEquals(tv1)

	assert.True(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003c=","Operands":[{"Name":"Simple1"},{"Name":"Simple2"}]}}}`,
		acal.ToString(actual),
	)

	tv2 := MakeSimpleFromInt("Simple3", 1)

	actual = op.SmallerThanEquals(tv2)

	assert.True(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple3":{"Value":"1","Source":"unknown"},"Unknown":{"Value":"true","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003c=","Operands":[{"Name":"Simple1"},{"Name":"Simple3"}]}}}`,
		acal.ToString(actual),
	)

	tv3 := MakeSimpleFromInt("Simple4", 0)

	actual = op.SmallerThanEquals(tv3)

	assert.False(t, actual.Bool())
	assert.Equal(
		t,
		`{"Simple1":{"Value":"1","Source":"unknown"},"Simple4":{"Value":"0","Source":"unknown"},"Unknown":{"Value":"false","Source":"static_calculation","Formula":{"Category":"TwoValMiddleOp","Operation":"\u003c=","Operands":[{"Name":"Simple1"},{"Name":"Simple4"}]}}}`,
		acal.ToString(actual),
	)
}
