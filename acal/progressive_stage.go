package acal

// Stage ...
type Stage[T any] struct {
	self      *Progressive[T]
	idx       int
	prevStage *Stage[T]

	value T

	sourceStageIdx int
	source         TypedValue[T]
}

// IsNil returns whether this Stage is nil.
func (s *Stage[T]) IsNil() bool {
	return s == nil
}

// GetName returns the name of the underlying Progressive.
func (s *Stage[T]) GetName() string {
	return s.self.GetName()
}

// GetAlias returns the alias of the underlying Progressive.
func (s *Stage[T]) GetAlias() string {
	return s.self.GetAlias()
}

// SetAlias updates the alias of the underlying Progressive to the provided string.
func (s *Stage[T]) SetAlias(alias string) {
	s.self.SetAlias(alias)
}

// HasIdentity returns whether the underlying Progressive was given an identity.
func (s *Stage[T]) HasIdentity() bool {
	return s.self.HasIdentity()
}

// Identify returns the identity that was given to the underlying Progressive.
func (s *Stage[T]) Identify() string {
	return s.self.Identify()
}

// GetTypedValue returns the typed value this Stage contains.
func (s *Stage[T]) GetTypedValue() T {
	return s.value
}

// GetValue returns the untyped value this Stage contains.
func (s *Stage[T]) GetValue() any {
	return s.GetTypedValue()
}

// ToSyntaxOperand returns the SyntaxOperand representation of this Stage.
func (s *Stage[T]) ToSyntaxOperand(nextOp Op) *SyntaxOperand {
	return NewSyntaxOperandWithStageIdx(s, s.idx)
}

// ExtractValues extracts this Stage and all Value that were used to calculate it.
func (s *Stage[T]) ExtractValues(cache IValueCache) IValueCache {
	return s.self.ExtractValues(cache)
}

// SelfReplaceIfNil returns the replacement to represent this Stage if it is nil.
func (s *Stage[T]) SelfReplaceIfNil() Value {
	if s.IsNil() {
		return ZeroSimple[T]("NilStage")
	}

	return s
}

// Stringify returns the value this Stage contains as a string.
func (s *Stage[T]) Stringify() string {
	return s.self.formatValue(s.value)
}
