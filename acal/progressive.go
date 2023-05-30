package acal

import (
	"encoding/json"
	"fmt"
)

// Progressive ...
type Progressive[T any] struct {
	name  string
	alias string

	curStage     *Stage[T]
	curCondition *Condition
	tags         Tags
}

// NewProgressive ...
func NewProgressive[T any](name string) *Progressive[T] {
	return &Progressive[T]{
		name: name,
	}
}

// IsNil returns whether this Progressive is nil.
func (p *Progressive[T]) IsNil() bool {
	return p == nil
}

// GetName returns the name of this Progressive.
func (p *Progressive[T]) GetName() string {
	return p.name
}

// GetAlias returns the alias of this Progressive.
func (p *Progressive[T]) GetAlias() string {
	return p.alias
}

// SetAlias updates the alias of this Progressive to the provided string.
func (p *Progressive[T]) SetAlias(alias string) {
	p.alias = alias
}

// GetTypedValue returns the typed value this Progressive contains.
func (p *Progressive[T]) GetTypedValue() T {
	if p.IsNil() || p.curStage == nil {
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

// SelfReplaceIfNil returns the replacement to represent this Progressive if it is nil.
func (p *Progressive[T]) SelfReplaceIfNil() Value {
	if p.IsNil() {
		return ZeroSimple[T]("NilProgressive")
	}

	return p
}

// GetTags returns the current tags of this Progressive.
func (p *Progressive[T]) GetTags() Tags {
	return p.tags
}

// Tag append the given Tag to the existing tags of this Progressive.
func (p *Progressive[T]) Tag(tags ...Tag) {
	p.tags = AppendTags(p, tags...)
}

// AsTag returns a Tag represented by this Progressive.
func (p *Progressive[T]) AsTag() Tag {
	return NewTagFrom(p.GetSnapshot())
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
	if fp, ok := value.(*Progressive[T]); ok {
		p.curStage = &Stage[T]{
			self:           p,
			idx:            p.getCurrentStageIdx() + 1,
			prevStage:      p.curStage,
			value:          value.GetTypedValue(),
			sourceStageIdx: fp.getCurrentStageIdx(),
			source:         fp,
		}

		return
	}

	p.curStage = &Stage[T]{
		self:      p,
		idx:       p.getCurrentStageIdx() + 1,
		prevStage: p.curStage,
		value:     value.GetTypedValue(),
		source:    value,
	}
}

// GetSnapshot returns the current Stage as a snapshot of this Progressive.
func (p *Progressive[T]) GetSnapshot() *Stage[T] {
	if p.IsNil() {
		return nil
	}

	if p.curStage == nil {
		p.Update(ZeroSimple[T]("Default"))
	}

	return p.curStage
}

// getStage returns the Stage at the given index.
func (p *Progressive[T]) getStage(stageIdx int) *Stage[T] {
	if stageIdx > p.curStage.idx {
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

// getCurrentStageIdx returns the index of the current Stage.
func (p *Progressive[T]) getCurrentStageIdx() int {
	if p.curStage == nil {
		return -1
	}

	return p.curStage.idx
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
		if fp, ok := curStage.source.(*Progressive[T]); ok {
			stages[curStage.idx] = jsonStage{
				Value:   fmt.Sprintf("%v", curStage.value),
				Formula: DescribeValueAsFormula(fp.getStage(curStage.sourceStageIdx))(),
			}

			curStage = curStage.prevStage

			continue
		}

		stages[curStage.idx] = jsonStage{
			Value:   fmt.Sprintf("%v", curStage.value),
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
