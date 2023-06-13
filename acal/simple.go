package acal

import (
	"encoding/json"
)

// Simple ...
type Simple[T any] struct {
    namedValue
    tagger
    staticConditioner
    valueFormatter[T]

    value     T
    source    Source
    formulaFn func() *SyntaxNode
}

// NewSimple ...
func NewSimple[T any](name string, value T) *Simple[T] {
    return &Simple[T]{
        namedValue: namedValue{
            name: name,
        },
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
    return &Simple[T]{
        value:     value.GetTypedValue(),
        formulaFn: DescribeValueAsFormula(value),
    }
}

// ZeroSimple ...
func ZeroSimple[T any](name string) *Simple[T] {
    var temp T
    return &Simple[T]{
        namedValue: namedValue{
            name: name,
        },
        value:  temp,
        source: SourceHardcode,
    }
}

// IsNil returns whether this Simple is nil.
func (s *Simple[T]) IsNil() bool {
    return s == nil
}

// GetTypedValue returns the typed value this Simple contains.
func (s *Simple[T]) GetTypedValue() T {
    return s.value
}

// GetValue returns the untyped value this Simple contains.
func (s *Simple[T]) GetValue() any {
    return s.GetTypedValue()
}

// ToSyntaxOperand returns the SyntaxOperand representation of this Simple.
func (s *Simple[T]) ToSyntaxOperand(nextOp Op) *SyntaxOperand {
    if HasIdentity(s) {
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

    if HasIdentity(s) {
        toAnchor = NewSimpleWithFormula(s.GetTypedValue(), DescribeValueAsFormula[T](s))
        isNew = true
    }

    toAnchor.name = name

    return toAnchor, isNew
}

// HasFormula returns whether this Simple has a formula.
func (s *Simple[T]) HasFormula() bool {
    return s.formulaFn != nil
}

// GetFormulaFn returns the function to build a formula of this Simple.
func (s *Simple[T]) GetFormulaFn() func() *SyntaxNode {
    return s.formulaFn
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
                Value:     s.Stringify(),
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
            Value:     s.Stringify(),
            Source:    string(s.source),
            Tags:      s.tags,
            Condition: s.condition,
        },
    )
}

// Stringify returns the value this Simple contains as a string.
func (s *Simple[T]) Stringify() string {
    return s.formatValue(s.value)
}
