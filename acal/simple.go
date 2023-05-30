package acal

import (
	"encoding/json"
	"fmt"
)

// Simple ...
type Simple[T any] struct {
	name  string
	alias string
	value T

	source    Source
	tags      Tags
	condition *Condition
	formulaFn func() *SyntaxNode
}

// NewSimple ...
func NewSimple[T any](name string, value T) *Simple[T] {
	return &Simple[T]{
		name:   name,
		value:  value,
		source: SourceUnknown,
	}
}

// NewSimpleWithFormula returns a new Simple with the given value and formula.
func NewSimpleWithFormula[T any](value T, formulaFn func() *SyntaxNode) *Simple[T] {
	return &Simple[T]{
		value:     value,
		formulaFn: formulaFn,
	}
}

// NewSimpleFrom returns a new Simple using the given value as formula.
func NewSimpleFrom[T any](value TypedValue[T]) *Simple[T] {
	p, ok := value.(*Progressive[T])
	if ok {
		snapshot := p.GetSnapshot()

		return &Simple[T]{
			value:     snapshot.GetTypedValue(),
			formulaFn: DescribeValueAsFormula(snapshot),
		}
	}

	return &Simple[T]{
		value:     value.GetTypedValue(),
		formulaFn: DescribeValueAsFormula(value),
	}
}

// ZeroSimple ...
func ZeroSimple[T any](name string) *Simple[T] {
	var temp T
	return &Simple[T]{
		name:   name,
		value:  temp,
		source: SourceHardcode,
	}
}

// IsNil returns whether this Simple is nil.
func (s *Simple[T]) IsNil() bool {
	return s == nil
}

// GetName returns the name of this Simple.
func (s *Simple[T]) GetName() string {
	return s.name
}

// GetAlias returns the alias of this Simple.
func (s *Simple[T]) GetAlias() string {
	return s.alias
}

// SetAlias updates the alias of this Simple to the provided string.
func (s *Simple[T]) SetAlias(alias string) {
	s.alias = alias
}

// GetTypedValue returns the typed value this Simple contains.
func (s *Simple[T]) GetTypedValue() T {
	if s.IsNil() {
		var temp T
		return temp
	}

	return s.value
}

// GetValue returns the untyped value this Simple contains.
func (s *Simple[T]) GetValue() any {
	return s.GetTypedValue()
}

// ToSyntaxOperand returns the SyntaxOperand representation of this Simple.
func (s *Simple[T]) ToSyntaxOperand(nextOp Op) *SyntaxOperand {
	if IsAnchored(s) {
		return NewSyntaxOperand(s)
	}

	if IsNilValue(s) || !s.HasFormula() {
		return NewSyntaxOperandWithStaticValue(Describe(s))
	}

	return nil
}

// ExtractValues extracts this Simple and all Value that were used to calculate it.
func (s *Simple[T]) ExtractValues(cache IValueCache) IValueCache {
	return PerformStandardValueExtraction(s, cache)
}

// SelfReplaceIfNil returns the replacement to represent this Simple if it is nil.
func (s *Simple[T]) SelfReplaceIfNil() Value {
	if s.IsNil() {
		return ZeroSimple[T]("NilSimple")
	}

	return s
}

// From updates the source of this Simple to the provided Source.
func (s *Simple[T]) From(source Source) {
	s.source = source
}

// GetSource returns the source of this Simple.
func (s *Simple[T]) GetSource() Source {
	return s.source
}

// DoAnchor updates the name of this Simple to the provided string.
func (s *Simple[T]) DoAnchor(name string) (*Simple[T], bool) {
	toAnchor := s
	isNew := false

	if IsAnchored(s) {
		toAnchor = NewSimpleWithFormula(s.GetTypedValue(), DescribeValueAsFormula(s))
		isNew = true
	}

	toAnchor.name = name

	return toAnchor, isNew
}

// GetTags returns the current tags of this Simple.
func (s *Simple[T]) GetTags() Tags {
	return s.tags
}

// Tag append the given Tag to the existing tags of this Simple.
func (s *Simple[T]) Tag(tags ...Tag) {
	s.tags = AppendTags(s, tags...)
}

// HasFormula returns whether this Simple has a formula.
func (s *Simple[T]) HasFormula() bool {
	return s.formulaFn != nil
}

// GetFormulaFn returns the function to build a formula of this Simple.
func (s *Simple[T]) GetFormulaFn() func() *SyntaxNode {
	return s.formulaFn
}

// IsConditional returns whether a Condition is attached to this Simple.
func (s *Simple[T]) IsConditional() bool {
	return s.condition != nil
}

// GetCondition returns the Condition attached to this Simple.
func (s *Simple[T]) GetCondition() *Condition {
	return s.condition
}

// AddCondition attaches the given Condition to this Simple.
func (s *Simple[T]) AddCondition(condition *Condition) {
	s.condition = condition
}

// MarshalJSON returns the JSON encoding of this Simple.
func (s *Simple[T]) MarshalJSON() ([]byte, error) {
	if s.IsNil() {
		return nil, nil
	}

	if s.HasFormula() {
		return json.Marshal(
			&struct {
				Value     string
				Source    string
				Tags      Tags       `json:",omitempty"`
				Condition *Condition `json:",omitempty"`
				Formula   *SyntaxNode
			}{
				Value:     fmt.Sprintf("%v", s.value),
				Source:    sourceStaticCalculation.String(),
				Tags:      s.tags,
				Condition: s.condition,
				Formula:   s.formulaFn(),
			},
		)
	}

	return json.Marshal(
		&struct {
			Value     string
			Source    string
			Tags      Tags       `json:",omitempty"`
			Condition *Condition `json:",omitempty"`
		}{
			Value:     fmt.Sprintf("%v", s.value),
			Source:    string(s.source),
			Tags:      s.tags,
			Condition: s.condition,
		},
	)
}
