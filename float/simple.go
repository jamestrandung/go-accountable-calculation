package float

import (
    "github.com/jamestrandung/go-accountable-calculation/acal"
    "github.com/shopspring/decimal"
)

// Simple ...
type Simple struct {
    *acal.Simple[decimal.Decimal]
    opProvider
}

// MakeSimple ...
func MakeSimple(name string, value float64) Simple {
    return MakeSimpleFromFloat(name, value)
}

// MakeSimpleFromFloat ...
func MakeSimpleFromFloat(name string, value float64) Simple {
    return makeSimple(acal.NewSimple(name, decimal.NewFromFloat(value)))
}

// MakeSimpleFromFloat32 ...
func MakeSimpleFromFloat32(name string, value float32) Simple {
    return makeSimple(acal.NewSimple(name, decimal.NewFromFloat32(value)))
}

// MakeSimpleFromInt ...
func MakeSimpleFromInt(name string, value int64) Simple {
    return makeSimple(acal.NewSimple(name, decimal.NewFromInt(value)))
}

// MakeSimpleFromInt32 ...
func MakeSimpleFromInt32(name string, value int32) Simple {
    return makeSimple(acal.NewSimple(name, decimal.NewFromInt32(value)))
}

// MakeSimpleFromDecimal ...
func MakeSimpleFromDecimal(name string, value decimal.Decimal) Simple {
    return makeSimple(acal.NewSimple(name, value))
}

// MakeSimpleWithFormula returns a new Simple with the given value and formula.
func MakeSimpleWithFormula(value decimal.Decimal, formulaFn func() *acal.SyntaxNode) Simple {
    return makeSimple(acal.NewSimpleWithFormula(value, formulaFn))
}

// MakeSimpleFrom returns a new Simple using the given value as formula.
func MakeSimpleFrom(value acal.TypedValue[decimal.Decimal]) Simple {
    return makeSimple(acal.NewSimpleFrom(value))
}

func makeSimple(core *acal.Simple[decimal.Decimal]) Simple {
    s := Simple{
        Simple: core,
    }

    s.opProvider = opProvider{
        tv: s,
    }

    s.WithFormatFn(FormatFn)

    return s
}

// GetTypedValue returns the typed value this Simple contains.
func (s Simple) GetTypedValue() decimal.Decimal {
    return acal.ExtractTypedValue[decimal.Decimal](s.Simple)
}

// ToSyntaxOperand returns the acal.SyntaxOperand representation of this Simple.
func (s Simple) ToSyntaxOperand(nextOp acal.Op) *acal.SyntaxOperand {
    if result := toBaseSyntaxOperand(s, nextOp); result != nil {
        return result
    }

    formula := s.GetFormulaFn()()
    lastOp := formula.GetOp()

    return acal.NewSyntaxOperandWithFormula(
        formula,
        !nextOp.IsTransparent() && !lastOp.IsTransparent() && opLevel[lastOp] < opLevel[nextOp],
    )
}

var toBaseSyntaxOperand = func(s Simple, nextOp acal.Op) *acal.SyntaxOperand {
    return s.Simple.ToSyntaxOperand(nextOp)
}

// Anchor updates the name of this Simple to the provided string.
func (s Simple) Anchor(name string) Simple {
    if s.IsNil() {
        return MakeSimpleFromInt(name, 0)
    }

    anchored, isNew := s.Simple.DoAnchor(name)
    if !isNew {
        return s
    }

    return Simple{
        Simple: anchored,
    }
}

// Decimal returns the value of this Simple as a decimal.Decimal.
// If it's nil, a decimal.Decimal value of 0 is returned.
func (s Simple) Decimal() decimal.Decimal {
    return acal.ExtractTypedValue[decimal.Decimal](s.Simple)
}

// Float returns the value of this Simple as a float64.
// If it's nil, 0 is returned.
func (s Simple) Float() float64 {
    return s.Decimal().InexactFloat64()
}

// Then does nothing and returns this Simple as-is. It's meant
// for separating code into more readable chunk.
func (s Simple) Then() Simple {
    return s
}
