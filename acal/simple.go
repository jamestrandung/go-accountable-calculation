package acal

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
        value:  value,
        name:   name,
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

// Anchor updates the name of this FloatV to the provided string
func (s *Simple[T]) Anchor(name string) *Simple[T] {
    if s.IsNil() {
        var temp T
        return NewSimple(name, temp)
    }

    toAnchor := s
    if IsAnchored(s) {
        toAnchor = NewSimpleWithFormula(s.GetTypedValue(), DescribeValueAsFormula(s))
    }

    toAnchor.name = name

    return toAnchor
}
