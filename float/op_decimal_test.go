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

func TestOpProvider_Neg(t *testing.T) {
    op1 := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}

    actual := op1.Neg()

    assert.Equal(t, -1.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple1":{"Value":"1","Source":"unknown"},"Unknown":{"Value":"-1","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Negate","Operands":[{"Name":"Simple1"}]}}}`,
        acal.ToString(actual),
    )

    op2 := opProvider{tv: MakeSimpleFromInt("Simple2", -1)}

    actual = op2.Neg()

    assert.Equal(t, 1.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple2":{"Value":"-1","Source":"unknown"},"Unknown":{"Value":"1","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Negate","Operands":[{"Name":"Simple2"}]}}}`,
        acal.ToString(actual),
    )
}

func TestOpProvider_Inv(t *testing.T) {
    op := opProvider{tv: MakeSimpleFromInt("Simple", 2)}

    actual := op.Inv()

    assert.Equal(t, 0.5, actual.Float())
    assert.Equal(
        t,
        `{"Simple":{"Value":"2","Source":"unknown"},"Unknown":{"Value":"0.5","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Inverse","Operands":[{"Name":"Simple"}]}}}`,
        acal.ToString(actual),
    )
}

func TestOpProvider_Abs(t *testing.T) {
    op1 := opProvider{tv: MakeSimpleFromInt("Simple1", 1)}

    actual := op1.Abs()

    assert.Equal(t, 1.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple1":{"Value":"1","Source":"unknown"},"Unknown":{"Value":"1","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Absolute","Operands":[{"Name":"Simple1"}]}}}`,
        acal.ToString(actual),
    )

    op2 := opProvider{tv: MakeSimpleFromInt("Simple2", -1)}

    actual = op2.Abs()

    assert.Equal(t, 1.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple2":{"Value":"-1","Source":"unknown"},"Unknown":{"Value":"1","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Absolute","Operands":[{"Name":"Simple2"}]}}}`,
        acal.ToString(actual),
    )
}

func TestOpProvider_Ceil(t *testing.T) {
    op1 := opProvider{tv: MakeSimpleFromFloat("Simple1", 1.4)}

    actual := op1.Ceil()

    assert.Equal(t, 2.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple1":{"Value":"1.4","Source":"unknown"},"Unknown":{"Value":"2","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Ceil","Operands":[{"Name":"Simple1"}]}}}`,
        acal.ToString(actual),
    )

    op2 := opProvider{tv: MakeSimpleFromFloat("Simple2", -1.6)}

    actual = op2.Ceil()

    assert.Equal(t, -1.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple2":{"Value":"-1.6","Source":"unknown"},"Unknown":{"Value":"-1","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Ceil","Operands":[{"Name":"Simple2"}]}}}`,
        acal.ToString(actual),
    )
}

func TestOpProvider_Floor(t *testing.T) {
    op1 := opProvider{tv: MakeSimpleFromFloat("Simple1", 1.4)}

    actual := op1.Floor()

    assert.Equal(t, 1.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple1":{"Value":"1.4","Source":"unknown"},"Unknown":{"Value":"1","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Floor","Operands":[{"Name":"Simple1"}]}}}`,
        acal.ToString(actual),
    )

    op2 := opProvider{tv: MakeSimpleFromFloat("Simple2", -1.6)}

    actual = op2.Floor()

    assert.Equal(t, -2.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple2":{"Value":"-1.6","Source":"unknown"},"Unknown":{"Value":"-2","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Floor","Operands":[{"Name":"Simple2"}]}}}`,
        acal.ToString(actual),
    )
}

func TestOpProvider_Round(t *testing.T) {
    op1 := opProvider{tv: MakeSimpleFromFloat("Simple1", 5.45)}

    actual := op1.Round(One)

    assert.Equal(t, 5.5, actual.Float())
    assert.Equal(
        t,
        `{"Simple1":{"Value":"5.45","Source":"unknown"},"Unknown":{"Value":"5.5","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Round","Operands":[{"Name":"Simple1"},{"StaticValue":"1"}]}}}`,
        acal.ToString(actual),
    )

    op2 := opProvider{tv: MakeSimpleFromFloat("Simple2", 545)}

    actual = op2.Round(MinusOne)

    assert.Equal(t, 550.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple2":{"Value":"545","Source":"unknown"},"Unknown":{"Value":"550","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Round","Operands":[{"Name":"Simple2"},{"StaticValue":"-1"}]}}}`,
        acal.ToString(actual),
    )
}

func TestMax(t *testing.T) {
    assert.Equal(t, NilFloat, Max())

    s1 := MakeSimpleFromInt("Simple1", 1)
    s2 := MakeSimpleFromInt("Simple2", 2)
    s3 := MakeSimpleFromInt("Simple3", 3)

    actual := Max(s1, s2, s3)

    assert.Equal(t, 3.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Simple3":{"Value":"3","Source":"unknown"},"Unknown":{"Value":"3","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Max","Operands":[{"Name":"Simple1"},{"Name":"Simple2"},{"Name":"Simple3"}]}}}`,
        acal.ToString(actual),
    )
}

func TestMin(t *testing.T) {
    assert.Equal(t, NilFloat, Min())

    s1 := MakeSimpleFromInt("Simple1", 1)
    s2 := MakeSimpleFromInt("Simple2", 2)
    s3 := MakeSimpleFromInt("Simple3", 3)

    actual := Min(s1, s2, s3)

    assert.Equal(t, 1.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Simple3":{"Value":"3","Source":"unknown"},"Unknown":{"Value":"1","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Min","Operands":[{"Name":"Simple1"},{"Name":"Simple2"},{"Name":"Simple3"}]}}}`,
        acal.ToString(actual),
    )
}

func TestAverage(t *testing.T) {
    assert.Equal(t, NilFloat, Average())

    s1 := MakeSimpleFromInt("Simple1", 1)
    s2 := MakeSimpleFromInt("Simple2", 2)
    s3 := MakeSimpleFromInt("Simple3", 3)

    actual := Average(s1, s2, s3)

    assert.Equal(t, 2.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Simple3":{"Value":"3","Source":"unknown"},"Unknown":{"Value":"2","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Average","Operands":[{"Name":"Simple1"},{"Name":"Simple2"},{"Name":"Simple3"}]}}}`,
        acal.ToString(actual),
    )
}

func TestMedian(t *testing.T) {
    assert.Equal(t, NilFloat, Median())

    s1 := MakeSimpleFromInt("Simple1", 1)
    s2 := MakeSimpleFromInt("Simple2", 2)
    s3 := MakeSimpleFromInt("Simple3", 3)
    s4 := MakeSimpleFromInt("Simple4", 4)

    actual := Median(s1, s2, s3, s4)

    assert.Equal(t, 2.5, actual.Float())
    assert.Equal(
        t,
        `{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Simple3":{"Value":"3","Source":"unknown"},"Simple4":{"Value":"4","Source":"unknown"},"Unknown":{"Value":"2.5","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Median","Operands":[{"Name":"Simple1"},{"Name":"Simple2"},{"Name":"Simple3"},{"Name":"Simple4"}]}}}`,
        acal.ToString(actual),
    )
}

func TestBoundBetween(t *testing.T) {
    s1 := MakeSimpleFromInt("Simple1", 1)
    s2 := MakeSimpleFromInt("Simple2", 2)
    s3 := MakeSimpleFromInt("Simple3", 3)
    s4 := MakeSimpleFromInt("Simple4", 4)

    actual := BoundBetween(s1, s2, s3)

    assert.Equal(t, 2.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple1":{"Value":"1","Source":"unknown"},"Simple2":{"Value":"2","Source":"unknown"},"Simple3":{"Value":"3","Source":"unknown"},"Unknown":{"Value":"2","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"BoundBetween","Operands":[{"Name":"Simple1"},{"Name":"Simple2"},{"Name":"Simple3"}]}}}`,
        acal.ToString(actual),
    )

    actual = BoundBetween(s4, s2, s3)

    assert.Equal(t, 3.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple2":{"Value":"2","Source":"unknown"},"Simple3":{"Value":"3","Source":"unknown"},"Simple4":{"Value":"4","Source":"unknown"},"Unknown":{"Value":"3","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"BoundBetween","Operands":[{"Name":"Simple4"},{"Name":"Simple2"},{"Name":"Simple3"}]}}}`,
        acal.ToString(actual),
    )
}

func TestPow(t *testing.T) {
    s1 := MakeSimpleFromInt("Simple1", 2)
    s2 := MakeSimpleFromInt("Simple2", 3)

    actual := Pow(s1, s2)

    assert.Equal(t, 8.0, actual.Float())
    assert.Equal(
        t,
        `{"Simple1":{"Value":"2","Source":"unknown"},"Simple2":{"Value":"3","Source":"unknown"},"Unknown":{"Value":"8","Source":"static_calculation","Formula":{"Category":"FunctionCall","Operation":"Pow","Operands":[{"Name":"Simple1"},{"Name":"Simple2"}]}}}`,
        acal.ToString(actual),
    )
}
