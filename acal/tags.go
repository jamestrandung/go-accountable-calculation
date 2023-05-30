package acal

import (
	"encoding/json"
	"fmt"
)

// Tag is a simple pair of name and value that can be attached to a
// Value to provide extra information.
type Tag struct {
	Name  string
	Value any
	aVal  Value
}

// NewTagFrom returns a Tag from the given Value.
func NewTagFrom(v Value) Tag {
	return Tag{
		Name:  v.GetName(),
		Value: v.GetValue(),
		aVal:  v,
	}
}

// GetName returns the name of this Tag.
func (t Tag) GetName() string {
	return t.Name
}

// GetValue returns the untyped value this Tag contains.
func (t Tag) GetValue() any {
	return t.Value
}

// Tagger represents any struct that can have tags attached to it.
//
//go:generate mockery --name=Tagger --case underscore --inpackage
type Tagger interface {
	// Tag append the given Tag to the existing tags of this Tagger.
	Tag(...Tag)
	// GetTags returns the current tags of this Tagger.
	GetTags() Tags
}

// Tags is the collection of tags that was applied on a Value.
type Tags []Tag

// AddTags adds the provided Tag to this Tags.
func (t Tags) AddTags(tags ...Tag) Tags {
	if t == nil {
		return tags
	}

	return append(t, tags...)
}

// Tag applies this collection of tags on the provided Tagger.
func (t Tags) Tag(taggers ...Tagger) {
	if len(t) == 0 || len(taggers) == 0 {
		return
	}

	for _, tagger := range taggers {
		tagger.Tag(t...)
	}
}

// MarshalJSON returns the JSON encoding of this Tags.
func (t Tags) MarshalJSON() ([]byte, error) {
	type tagJSON struct {
		Value   string
		IsValue bool `json:",omitempty"`
	}

	tags := make(map[string]tagJSON)
	for _, tag := range t {
		tags[tag.GetName()] = tagJSON{
			Value:   fmt.Sprintf("%v", tag.GetValue()),
			IsValue: tag.aVal != nil,
		}
	}

	return json.Marshal(&tags)
}

// AppendTags appends the provided Tag to collection of
// tags of the given Tagger.
var AppendTags = func(t Tagger, tags ...Tag) Tags {
	if t == nil {
		return nil
	}

	curTags := t.GetTags()
	if curTags == nil {
		return tags
	}

	return curTags.AddTags(tags...)
}
