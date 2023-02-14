package acal

import (
	"fmt"
)

// Constant ...
type Constant[T any] struct {
	value T
}

// NewConstant ...
func NewConstant[T any](value T) *Constant[T] {
	return &Constant[T]{
		value: value,
	}
}

// IsNil returns whether this Constant is nil.
func (c *Constant[T]) IsNil() bool {
	return c == nil
}

// GetName always returns empty string for Constant as constant values
// will be used directly in formula without any reference to name.
func (c *Constant[T]) GetName() string {
	return ""
}

// GetAlias always returns empty string for Constant as constant values
// will be used directly in formula without any reference to alias.
func (c *Constant[T]) GetAlias() string {
	return ""
}

// SetAlias does nothing for Constant as it's a constant.
func (c *Constant[T]) SetAlias(alias string) {

}

// GetTypedValue returns the typed value this Constant contains.
func (c *Constant[T]) GetTypedValue() T {
	if c.IsNil() {
		var temp T
		return temp
	}

	return c.value
}

// GetValue returns the untyped value this Constant contains.
func (c *Constant[T]) GetValue() any {
	return c.GetTypedValue()
}

// ToSyntaxOperand returns the acal.SyntaxOperand representation of this Constant.
func (c *Constant[T]) ToSyntaxOperand(nextOp Op) *SyntaxOperand {
	return NewSyntaxOperandWithStaticValue(fmt.Sprintf("%v", c.GetTypedValue()))
}

// HasFormula returns whether this Constant has a formula.
func (c *Constant[T]) HasFormula() bool {
	return true
}

// GetFormulaFn returns the function to build a formula of this Constant.
func (c *Constant[T]) GetFormulaFn() func() *SyntaxNode {
	return func() *SyntaxNode {
		return NewSyntaxNode(OpCategoryAssignStatic, OpTransparent, fmt.Sprintf("%v", c.GetTypedValue()), nil)
	}
}

// ExtractValues does nothing for Constant as it's a constant.
func (c *Constant[T]) ExtractValues(cache IValueCache) IValueCache {
	return cache
}

// SelfReplaceIfNil returns the replacement to represent this Constant if it is nil.
func (c *Constant[T]) SelfReplaceIfNil() Value {
	if c.IsNil() {
		return ZeroAny[T]("NilConstant")
	}

	return c
}
