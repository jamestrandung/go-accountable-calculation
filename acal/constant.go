package acal

// Constant ...
type Constant[T any] struct {
	namedValue
	valueFormatter[T]

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

// ToSyntaxOperand returns the SyntaxOperand representation of this Constant.
func (c *Constant[T]) ToSyntaxOperand(nextOp Op) *SyntaxOperand {
	return NewSyntaxOperandWithStaticValue(c.Stringify())
}

// HasFormula returns whether this Constant has a formula.
func (c *Constant[T]) HasFormula() bool {
	return true
}

// GetFormula returns the formula provided by this Constant.
func (c *Constant[T]) GetFormula() *SyntaxNode {
	return NewSyntaxNode(OpCategoryAssignStatic, OpTransparent, c.Stringify(), nil)
}

// ExtractValues does nothing for Constant as it's a constant.
func (c *Constant[T]) ExtractValues(cache IValueCache) IValueCache {
	return cache
}

// SelfReplaceIfNil returns the replacement to represent this Constant if it is nil.
func (c *Constant[T]) SelfReplaceIfNil() Value {
	if c.IsNil() {
		return ZeroSimple[T]("NilConstant")
	}

	return c
}

// Stringify returns the value this Constant contains as a string.
func (c *Constant[T]) Stringify() string {
	return c.formatValue(c.value)
}
