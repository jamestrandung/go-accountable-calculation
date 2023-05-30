package acal

import (
	"encoding/json"
	"fmt"
)

// Remote ...
type Remote[T any] struct {
	namedValue
	tagger
	staticConditioner
	iValueFormatter[T]

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
	if r.IsNil() {
		var temp T
		return temp
	}

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

// SelfReplaceIfNil returns the replacement to represent this Remote if it is nil.
func (r *Remote[T]) SelfReplaceIfNil() Value {
	if r.IsNil() {
		return ZeroSimple[T]("NilRemote")
	}

	return r
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

	v := func() string {
		if r.iValueFormatter != nil {
			return r.formatValue(r.GetTypedValue())
		}

		return fmt.Sprintf("%v", r.GetTypedValue())
	}()

	return json.Marshal(
		&struct {
			Value          string
			Source         string
			DependentField string
			LogKey         string
			Tags           Tags       `json:",omitempty"`
			Condition      *Condition `json:",omitempty"`
		}{
			Value:          v,
			Source:         sourceRemoteCalculation.String(),
			DependentField: r.fieldName,
			LogKey:         r.logKey,
			Tags:           r.tags,
			Condition:      r.condition,
		},
	)
}
