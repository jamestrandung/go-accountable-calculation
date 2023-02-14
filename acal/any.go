package acal

import (
	"encoding/json"
)

// Any ...
type Any[T any] struct {
	name  string
	alias string
	value T

	source    Source
	tags      Tags
	condition *Condition
}

// NewAny ...
func NewAny[T any](name string, value T) *Any[T] {
	return &Any[T]{
		name:   name,
		value:  value,
		source: SourceUnknown,
	}
}

// ZeroAny ...
func ZeroAny[T any](name string) *Any[T] {
	var temp T
	return &Any[T]{
		name:   name,
		value:  temp,
		source: SourceHardcode,
	}
}

// IsNil returns whether this Any is nil.
func (a *Any[T]) IsNil() bool {
	return a == nil
}

// GetName returns the name of this Any.
func (a *Any[T]) GetName() string {
	return a.name
}

// GetAlias returns the alias of this Any.
func (a *Any[T]) GetAlias() string {
	return a.alias
}

// SetAlias updates the alias of this Any to the provided string.
func (a *Any[T]) SetAlias(alias string) {
	a.alias = alias
}

// GetTypedValue returns the typed value this Any contains.
func (a *Any[T]) GetTypedValue() T {
	if a.IsNil() {
		var temp T
		return temp
	}

	return a.value
}

// GetValue returns the value of this Any.
func (a *Any[T]) GetValue() any {
	return a.GetTypedValue()
}

// ToSyntaxOperand returns the acal.SyntaxOperand representation of this Any.
func (a *Any[T]) ToSyntaxOperand(nextOp Op) *SyntaxOperand {
	return NewSyntaxOperand(a)
}

// From updates the source of this Any to the provided acal.Source.
func (a *Any[T]) From(source Source) {
	a.source = source
}

// GetTags returns the current tags of this Any.
func (a *Any[T]) GetTags() Tags {
	return a.tags
}

// Tag append the given acal.NameValuePair to the existing tags of this Any.
func (a *Any[T]) Tag(tags ...Tag) {
	a.tags = AppendTags(a, tags...)
}

// IsConditional returns whether a condition is attached to this Any.
func (a *Any[T]) IsConditional() bool {
	return a.condition != nil
}

// GetCondition returns the condition attached to this Any.
func (a *Any[T]) GetCondition() *Condition {
	return a.condition
}

// AddCondition attaches the given condition to this Any.
func (a *Any[T]) AddCondition(condition *Condition) {
	if a.IsNil() {
		return
	}

	a.condition = condition
}

// ExtractValues extracts this Any and all acal.Value that were used to calculate it.
func (a *Any[T]) ExtractValues(cache IValueCache) IValueCache {
	return PerformStandardValueExtraction(a, cache)
}

// SelfReplaceIfNil returns the replacement to represent this Any if it is nil.
func (a *Any[T]) SelfReplaceIfNil() Value {
	if a.IsNil() {
		return ZeroAny[T]("NilAny")
	}

	return a
}

// MarshalJSON returns the JSON encoding of this Any.
func (a *Any[T]) MarshalJSON() ([]byte, error) {
	if a.IsNil() {
		return nil, nil
	}

	return json.Marshal(
		&struct {
			Value     any
			Source    string
			Tags      Tags       `json:",omitempty"`
			Condition *Condition `json:",omitempty"`
		}{
			Value:     a.value,
			Source:    string(a.source),
			Tags:      a.tags,
			Condition: a.condition,
		},
	)
}
