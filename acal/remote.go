package acal

import (
	"encoding/json"
)

// Remote ...
type Remote[T any] struct {
	namedValue
	tagger
	staticConditioner
	valueFormatter[T]

	value     T
	fieldName string
	logKey    string
}

// NewRemote ...
func NewRemote[T any](name string, value T, remoteFieldName string, remoteLogKey string) *Remote[T] {
	return &Remote[T]{
		namedValue: namedValue{
			name: name,
		},
		value:     value,
		fieldName: remoteFieldName,
		logKey:    remoteLogKey,
	}
}

// IsNil returns whether this Remote is nil.
func (r *Remote[T]) IsNil() bool {
	return r == nil
}

// GetTypedValue returns the typed value this Remote contains.
func (r *Remote[T]) GetTypedValue() T {
	return r.value
}

// GetValue returns the untyped value this Remote contains.
func (r *Remote[T]) GetValue() any {
	return r.GetTypedValue()
}

// ToSyntaxOperand returns the SyntaxOperand representation of this Remote.
func (r *Remote[T]) ToSyntaxOperand(nextOp Op) *SyntaxOperand {
	return NewSyntaxOperand(r)
}

// ExtractValues extracts this Remote and all Value that were used to calculate it.
func (r *Remote[T]) ExtractValues(cache IValueCache) IValueCache {
	return PerformStandardValueExtraction(r, cache)
}

// DoAnchor returns a new Simple initialized to the value of this
// Remote and anchored with the given name.
func (r *Remote[T]) DoAnchor(name string) *Simple[T] {
	s := NewSimpleFrom[T](r)
	s.DoAnchor(name)

	return s
}

// MarshalJSON returns the JSON encoding of this Remote.
func (r *Remote[T]) MarshalJSON() ([]byte, error) {
	if r.IsNil() {
		return nil, nil
	}

	return json.Marshal(
		&struct {
			Value          string
			Source         string
			DependentField string
			LogKey         string
			Tags           Tags       `json:",omitempty"`
			Condition      *Condition `json:",omitempty"`
		}{
			Value:          r.Stringify(),
			Source:         sourceRemoteCalculation.String(),
			DependentField: r.fieldName,
			LogKey:         r.logKey,
			Tags:           r.tags,
			Condition:      r.condition,
		},
	)
}

// Stringify returns the value this Remote contains as a string.
func (r *Remote[T]) Stringify() string {
	return r.formatValue(r.GetTypedValue())
}
