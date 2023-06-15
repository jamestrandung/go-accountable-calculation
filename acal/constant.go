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

// NewConstantWithFormat ...
func NewConstantWithFormat[T any](value T, formatFn func(T) string) *Constant[T] {
	c := &Constant[T]{
		value: value,
	}

	c.WithFormatFn(formatFn)

	return c
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

// GetFormulaFn returns the function to build a formula of this Constant.
func (c *Constant[T]) GetFormulaFn() func() *SyntaxNode {
	return func() *SyntaxNode {
		return NewSyntaxNode(OpCategoryAssignStatic, OpTransparent, c.Stringify(), nil)
	}
}

// ExtractValues does nothing for Constant as it's a constant.
func (c *Constant[T]) ExtractValues(cache IValueCache) IValueCache {
	return cache
}

// Stringify returns the value this Constant contains as a string.
func (c *Constant[T]) Stringify() string {
	return c.formatValue(c.value)
}
