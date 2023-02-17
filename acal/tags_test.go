package acal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTags_AddTags(t *testing.T) {
	testTags := []Tag{{Name: "TestTag1"}, {Name: "TestTag2"}}

	scenarios := []struct {
		desc string
		tags Tags
		want Tags
	}{
		{
			desc: "nil tags",
			tags: nil,
			want: testTags,
		},
		{
			desc: "no existing tags",
			tags: Tags{},
			want: testTags,
		},
		{
			desc: "has existing tags",
			tags: Tags{{Name: "TestTag3"}},
			want: Tags{{Name: "TestTag3"}, {Name: "TestTag1"}, {Name: "TestTag2"}},
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				actual := sc.tags.AddTags(testTags...)

				assert.Equal(t, sc.want, actual)
			},
		)
	}
}

func TestTags_Tag(t *testing.T) {
	tag := Tag{Name: "TestTag"}

	mockTagger1 := &MockTagger{}
	mockTagger1.On("Tag", tag).Once()

	mockTagger2 := &MockTagger{}
	mockTagger2.On("Tag", tag).Once()

	tags := Tags{tag}

	tags.Tag(mockTagger1, mockTagger2)

	mock.AssertExpectationsForObjects(t, mockTagger1, mockTagger2)
}

func TestTags_MarshalJSON(t *testing.T) {
	tag1 := Tag{Name: "TestName1", Value: 5}
	tag2 := Tag{Name: "TestName2", Value: true}
	tag3 := Tag{Name: "TestName3", Value: 6.5, aVal: &MockValue{}}

	wantedJSON := "{\"TestName1\":{\"Value\":\"5\"},\"TestName2\":{\"Value\":\"true\"},\"TestName3\":{\"Value\":\"6.5\",\"IsValue\":true}}"

	tags := Tags{tag1, tag2, tag3}

	actualJSON, err := tags.MarshalJSON()
	assert.Equal(t, wantedJSON, string(actualJSON), "marshal result should be %v", wantedJSON)
	assert.Nil(t, err, "error should be nil")
}

func TestAppendTags(t *testing.T) {
	mockTagger := &MockTagger{}

	scenarios := []struct {
		desc   string
		tagger Tagger
		setup  func()
		want   Tags
	}{
		{
			desc:   "nil Tagger",
			tagger: nil,
			want:   nil,
		},
		{
			desc:   "Tagger has no existing tags",
			tagger: mockTagger,
			setup: func() {
				mockTagger.On("GetTags").Return(nil).Once()
			},
			want: Tags{Tag{Name: "TestName2", Value: true}},
		},
		{
			desc:   "Tagger has existing tags",
			tagger: mockTagger,
			setup: func() {
				mockTagger.On("GetTags").Return(Tags{Tag{Name: "TestName1", Value: 5}}).Twice()
			},
			want: Tags{Tag{Name: "TestName1", Value: 5}, Tag{Name: "TestName2", Value: true}},
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				if sc.setup != nil {
					sc.setup()
				}

				actual := AppendTags(sc.tagger, Tag{Name: "TestName2", Value: true})

				assert.Equal(t, sc.want, actual)
			},
		)
	}
}
