package acal

import (
	"encoding/json"
)

// Progressive ...
type Progressive[T any] struct {
	namedValue
	tagger
	valueFormatter[T]

	curStage     *Stage[T]
	curCondition *Condition
}

// NewProgressive ...
func NewProgressive[T any](name string) *Progressive[T] {
	return &Progressive[T]{
		namedValue: namedValue{
			name: name,
		},
	}
}

// IsNil returns whether this Progressive is nil.
func (p *Progressive[T]) IsNil() bool {
	return p == nil
}

// GetTypedValue returns the typed value this Progressive contains.
func (p *Progressive[T]) GetTypedValue() T {
	if p.curStage == nil {
		var temp T
		return temp
	}

	return p.curStage.value
}

// GetValue returns the untyped value this Progressive contains.
func (p *Progressive[T]) GetValue() any {
	return p.GetTypedValue()
}

// ToSyntaxOperand returns the SyntaxOperand representation of this Progressive.
func (p *Progressive[T]) ToSyntaxOperand(nextOp Op) *SyntaxOperand {
	return NewSyntaxOperandWithStageIdx(p, p.getCurrentStageIdx())
}

// ExtractValues extracts this Progressive and all Value that were used to calculate it.
func (p *Progressive[T]) ExtractValues(cache IValueCache) IValueCache {
	if p.IsNil() || !cache.Take(p) {
		return cache
	}

	curStage := p.curStage
	for curStage != nil {
		curStage.source.ExtractValues(cache)

		curStage = curStage.prevStage
	}

	curCondition := p.curCondition
	for curCondition != nil {
		curCondition.criteria.ExtractValues(cache)

		curCondition = curCondition.prevCondition
	}

	for _, tag := range p.tags {
		if tag.aVal != nil {
			tag.aVal.ExtractValues(cache)
		}
	}

	return cache
}

// AsTag returns a Tag represented by this Progressive.
func (p *Progressive[T]) AsTag() Tag {
	return NewTagFrom(p)
}

// AddCondition attaches the given condition to this Progressive, returning
// a  CloseIfFunc that must be triggered when an if clause ends so that the
// framework can record at which stage this condition ends.
func (p *Progressive[T]) AddCondition(criteria TypedValue[bool]) CloseIfFunc {
	condition := NewProgressiveCondition(criteria, p.curCondition, p.getCurrentStageIdx()+1)

	p.curCondition = condition

	return func() {
		condition.setCloseIfStageIdx(p.getCurrentStageIdx())
	}
}

// DoAnchor returns a new Simple initialized to the current value of this
// Progressive and anchored with the given name.
func (p *Progressive[T]) DoAnchor(name string) *Simple[T] {
	s := NewSimpleFrom[T](p)
	s.DoAnchor(name)

	return s
}

// Update adds a new Stage to this Progressive to record its new value.
func (p *Progressive[T]) Update(value TypedValue[T]) {
	if IsNilValue(value) {
		value = ZeroSimple[T]("NilProgressiveUpdate")
	}

	if fp, ok := value.(*Progressive[T]); ok {
		value = fp.getSnapshot()
	}

	p.curStage = &Stage[T]{
		self:      p,
		idx:       p.getCurrentStageIdx() + 1,
		prevStage: p.curStage,
		value:     value.GetTypedValue(),
		source:    value,
	}
}

// getCurrentStageIdx returns the index of the current Stage.
func (p *Progressive[T]) getCurrentStageIdx() int {
	if p.curStage == nil {
		return -1
	}

	return p.curStage.idx
}

// getSnapshot returns the current Stage as a snapshot of this Progressive.
func (p *Progressive[T]) getSnapshot() *Stage[T] {
	if p.IsNil() {
		return nil
	}

	if p.curStage == nil {
		p.Update(ZeroSimple[T]("Default" + p.name))
	}

	return p.curStage
}

// getStage returns the Stage at the given index.
func (p *Progressive[T]) getStage(stageIdx int) *Stage[T] {
	if stageIdx < 0 || p.curStage == nil || stageIdx > p.curStage.idx {
		return nil
	}

	curStage := p.curStage
	for curStage != nil {
		if curStage.idx == stageIdx {
			return curStage
		}

		curStage = curStage.prevStage
	}

	return nil
}

type jsonStage struct {
	Value   string
	Formula *SyntaxNode
}

// MarshalJSON returns the JSON encoding of this FloatP
func (p *Progressive[T]) MarshalJSON() ([]byte, error) {
	if p.IsNil() {
		return nil, nil
	}

	stages := make([]jsonStage, p.getCurrentStageIdx()+1)

	curStage := p.curStage
	for curStage != nil {
		stages[curStage.idx] = jsonStage{
			Value:   curStage.Stringify(),
			Formula: DescribeValueAsFormula(curStage.source)(),
		}

		curStage = curStage.prevStage
	}

	conds := make(map[int][]*Condition)

	curCondition := p.curCondition
	for curCondition != nil {
		if !curCondition.isValidProgressiveCondition() {
			curCondition = curCondition.prevCondition
			continue
		}

		conds[curCondition.openIfStageIdx] = append(conds[curCondition.openIfStageIdx], curCondition)

		curCondition = curCondition.prevCondition
	}

	// Reverse the condition slices to get the correct order
	for _, conditions := range conds {
		for i, j := 0, len(conditions)-1; i < j; i, j = i+1, j-1 {
			conditions[i], conditions[j] = conditions[j], conditions[i]
		}
	}

	return json.Marshal(
		&struct {
			Source     string
			Stages     []jsonStage
			Conditions map[int][]*Condition `json:",omitempty"`
			Tags       Tags                 `json:",omitempty"`
		}{
			Source:     sourceProgressiveCalculation.String(),
			Stages:     stages,
			Conditions: conds,
			Tags:       p.tags,
		},
	)
}

// Stringify returns the value this Progressive contains as a string.
func (p *Progressive[T]) Stringify() string {
	return p.formatValue(p.GetTypedValue())
}
