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
		Value: v.Stringify(),
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

// ITagger represents any struct that can have tags attached to it.
//
//go:generate mockery --name=ITagger --case underscore --inpackage
type ITagger interface {
	// Tag append the given Tag to the existing tags of this ITagger.
	Tag(...Tag)
	// GetTags returns the current tags of this ITagger.
	GetTags() Tags
}

type tagger struct {
	tags Tags
}

// GetTags returns the current tags of this tagger.
func (t *tagger) GetTags() Tags {
	return t.tags
}

// Tag append the given Tag to the existing tags of this tagger.
func (t *tagger) Tag(tags ...Tag) {
	t.tags = AppendTags(t, tags...)
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

// Tag applies this collection of tags on the provided ITagger.
func (t Tags) Tag(taggers ...ITagger) {
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
// tags of the given ITagger.
var AppendTags = func(t ITagger, tags ...Tag) Tags {
	if t == nil {
		return nil
	}

	curTags := t.GetTags()
	if curTags == nil {
		return tags
	}

	return curTags.AddTags(tags...)
}
