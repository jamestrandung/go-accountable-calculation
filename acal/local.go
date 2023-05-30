package acal

import (
	"encoding/json"
	"fmt"
)

// Local ...
type Local[T any] struct {
	namedValue
	tagger
	staticConditioner
	iValueFormatter[T]

	original TypedValue[T]
}

// NewLocal ...
func NewLocal[T any](name string, original TypedValue[T]) *Local[T] {
	if original.IsNil() {
		original = ZeroSimple[T]("NilOriginal")
	}

	return &Local[T]{
		namedValue: namedValue{
			name: name,
		},
		original: original,
	}
}

// IsNil returns whether this Local is nil.
func (l *Local[T]) IsNil() bool {
	return l == nil
}

// GetTypedValue returns the typed value this Local contains.
func (l *Local[T]) GetTypedValue() T {
	if l.IsNil() {
		var temp T
		return temp
	}

	return l.original.GetTypedValue()
}

// GetValue returns the untyped value this Local contains.
func (l *Local[T]) GetValue() any {
	return l.GetTypedValue()
}

// ToSyntaxOperand returns the SyntaxOperand representation of this Local.
func (l *Local[T]) ToSyntaxOperand(nextOp Op) *SyntaxOperand {
	return NewSyntaxOperand(l)
}

// ExtractValues extracts this Local and all Value that were used to calculate it.
func (l *Local[T]) ExtractValues(cache IValueCache) IValueCache {
	return PerformStandardValueExtraction(l, cache)
}

// SelfReplaceIfNil returns the replacement to represent this Local if it is nil.
func (l *Local[T]) SelfReplaceIfNil() Value {
	if l.IsNil() {
		return ZeroSimple[T]("NilLocal")
	}

	return l
}

// DoAnchor returns a new Simple initialized to the value of this
// Local and anchored with the given name.
func (l *Local[T]) DoAnchor(name string) *Simple[T] {
	s := NewSimpleFrom[T](l)
	s.DoAnchor(name)

	return s
}

// MarshalJSON returns the JSON encoding of this Local.
func (l *Local[T]) MarshalJSON() ([]byte, error) {
	if l.IsNil() {
		return nil, nil
	}

	v := func() string {
		if l.iValueFormatter != nil {
			return l.formatValue(l.GetTypedValue())
		}

		return fmt.Sprintf("%v", l.GetTypedValue())
	}()

	return json.Marshal(
		&struct {
			Value          string
			Source         string
			DependentField string
			Calculation    map[string]Value
			Tags           Tags       `json:",omitempty"`
			Condition      *Condition `json:",omitempty"`
		}{
			Value:          v,
			Source:         sourceDependentCalculation.String(),
			DependentField: l.original.GetName(),
			Calculation:    l.original.ExtractValues(NewValueCache()).Flatten(),
			Tags:           l.tags,
			Condition:      l.condition,
		},
	)
}
