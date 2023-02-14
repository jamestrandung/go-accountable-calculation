package acal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTags_AddTags(t *testing.T) {
	testNameValuePairs := []NameValuePair{&MockNameValuePair{}, &MockNameValuePair{}}

	scenarios := []struct {
		desc string
		tags Tags
		want Tags
	}{
		{
			desc: "nil tags",
			tags: nil,
			want: Tags{&MockNameValuePair{}, &MockNameValuePair{}},
		},
		{
			desc: "no existing tags",
			tags: Tags{},
			want: Tags{&MockNameValuePair{}, &MockNameValuePair{}},
		},
		{
			desc: "has existing tags",
			tags: Tags{&MockNameValuePair{}},
			want: Tags{&MockNameValuePair{}, &MockNameValuePair{}, &MockNameValuePair{}},
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				actual := sc.tags.AddTags(testNameValuePairs...)

				assert.Equal(t, sc.want, actual)
			},
		)
	}
}

func TestTags_Tag(t *testing.T) {
	mockNameValuePair := &MockNameValuePair{}

	mockTagger1 := &MockTagger{}
	mockTagger1.On("Tag", mockNameValuePair).Once()

	mockTagger2 := &MockTagger{}
	mockTagger2.On("Tag", mockNameValuePair).Once()

	tags := Tags{mockNameValuePair}

	tags.Tag(mockTagger1, mockTagger2)

	mock.AssertExpectationsForObjects(t, mockTagger1, mockTagger2)
}

func TestTags_MarshalJSON(t *testing.T) {
	mockNameValuePair1 := &MockNameValuePair{}
	mockNameValuePair1.On("GetName").Return("TestName1").Once()
	mockNameValuePair1.On("GetValue").Return(5).Once()

	mockNameValuePair2 := &MockNameValuePair{}
	mockNameValuePair2.On("GetName").Return("TestName2").Once()
	mockNameValuePair2.On("GetValue").Return(true).Once()

	wantedJSON := "{\"TestName1\":{\"Value\":\"5\"},\"TestName2\":{\"Value\":\"true\"}}"

	tags := Tags{mockNameValuePair1, mockNameValuePair2}

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
			want: Tags{&MockNameValuePair{}},
		},
		{
			desc:   "Tagger has existing tags",
			tagger: mockTagger,
			setup: func() {
				mockTagger.On("GetTags").Return(Tags{&MockNameValuePair{}}).Twice()
			},
			want: Tags{&MockNameValuePair{}, &MockNameValuePair{}},
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				if sc.setup != nil {
					sc.setup()
				}

				actual := AppendTags(sc.tagger, &MockNameValuePair{})

				assert.Equal(t, sc.want, actual)
			},
		)
	}
}
