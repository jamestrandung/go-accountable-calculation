package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/base"
)

type floatC struct {
	*base.Constant[float64]
}

// floatConstant returns a new floatC with the provided fields
func floatConstant(value float64) *floatC {
	return &floatC{
		Constant: base.NewConstant(value),
	}
}

// ToSyntaxOperand returns the core.SyntaxOperand representation of this Constant
func (f *floatC) ToSyntaxOperand(nextOp acal.Op) *acal.SyntaxOperand {
	return acal.NewSyntaxOperandWithStaticValue(acal.FormatFloatForMarshalling(roundToDecimalPlace(f.GetTypedValue(), 4)))
}

// GetFormulaFn returns the function to build a formula of this Constant
func (f *floatC) GetFormulaFn() func() *acal.SyntaxNode {
	return func() *acal.SyntaxNode {
		return acal.NewSyntaxNode(
			acal.OpCategoryAssignStatic,
			acal.OpTransparent,
			acal.FormatFloatForMarshalling(roundToDecimalPlace(f.GetTypedValue(), 4)),
			nil,
		)
	}
}

// SelfReplaceIfNil returns the replacement to represent this Constant if it is nil
func (f *floatC) SelfReplaceIfNil() acal.Value {
	if f.IsNil() {
		return NilFloat
	}

	return f
}
