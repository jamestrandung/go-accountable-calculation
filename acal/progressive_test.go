package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProgressive_IsNil(t *testing.T) {
	var nilProgressive *Progressive[int]

	assert.True(t, nilProgressive.IsNil())

	progressive := NewProgressive[int]("Progressive")

	assert.False(t, progressive.IsNil())
}

func TestProgressive_GetTypedValue(t *testing.T) {
	progressive := NewProgressive[int]("Progressive")

	assert.Equal(t, 0, progressive.GetTypedValue())

	progressive.Update(NewConstant(2))

	assert.Equal(t, 2, progressive.GetTypedValue())
}

func TestProgressive_ToSyntaxOperand(t *testing.T) {
	progressive := NewProgressive[int]("Progressive")

	actual := progressive.ToSyntaxOperand(OpTransparent)

	assert.Equal(
		t, &SyntaxOperand{
			Name:     "Progressive",
			StageIdx: -1,
		}, actual,
	)

	progressive.Update(NewConstant(2))

	actual = progressive.ToSyntaxOperand(OpTransparent)

	assert.Equal(
		t, &SyntaxOperand{
			Name:     "Progressive",
			StageIdx: 0,
		}, actual,
	)
}

func TestProgressive_ExtractValues(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil Progressive",
			test: func(t *testing.T) {
				mockCache := NewMockIValueCache(t)

				var nilProgressive *Progressive[int]

				actual := nilProgressive.ExtractValues(mockCache)

				assert.Equal(t, mockCache, actual)
			},
		},
		{
			desc: "Progressive already taken",
			test: func(t *testing.T) {
				progressive := NewProgressive[int]("Progressive")
				progressive.Update(NewConstant(1))
				progressive.Update(NewConstant(2))

				mockCache := NewMockIValueCache(t)
				mockCache.On("Take", progressive).Return(false).Once()

				actual := progressive.ExtractValues(mockCache)

				assert.Equal(t, mockCache, actual)
			},
		},
		{
			desc: "Progressive not yet taken",
			test: func(t *testing.T) {
				mockCache := NewMockIValueCache(t)

				mockValue1 := NewMockTypedValue[int](t)
				mockValue1.On("GetTypedValue").
					Return(1).
					Once()
				mockValue1.On("ExtractValues", mockCache).
					Return(mockCache).
					Once()

				mockValue2 := NewMockTypedValue[int](t)
				mockValue2.On("GetTypedValue").
					Return(2).
					Once()
				mockValue2.On("ExtractValues", mockCache).
					Return(mockCache).
					Once()

				mockCriteria := NewMockTypedValue[bool](t)
				mockCriteria.On("ExtractValues", mockCache).
					Return(mockCache).
					Once()

				mockTag := NewMockValue(t)
				mockTag.On("GetName").
					Return("Tag").
					Once()
				mockTag.On("Stringify").
					Return("TagContent").
					Once()
				mockTag.On("ExtractValues", mockCache).
					Return(mockCache).
					Once()

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue1).
					Return(false).
					Once()
				mockValueOps.On("IsNilValue", mockValue2).
					Return(false).
					Once()

				progressive := NewProgressive[int]("Progressive")
				progressive.Update(mockValue1)
				progressive.Update(mockValue2)

				progressive.AddCondition(mockCriteria)
				progressive.Tag(NewTagFrom(mockTag))

				mockCache.On("Take", progressive).Return(true).Once()

				actual := progressive.ExtractValues(mockCache)

				assert.Equal(t, mockCache, actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestProgressive_AsTag(t *testing.T) {
	progressive := NewProgressive[int]("Progressive")

	actual := progressive.AsTag()

	assert.Equal(
		t, Tag{
			Name:  "Progressive",
			Value: "0",
			aVal:  progressive,
		}, actual,
	)

	progressive.Update(NewConstant(2))

	actual = progressive.AsTag()

	assert.Equal(
		t, Tag{
			Name:  "Progressive",
			Value: "2",
			aVal:  progressive,
		}, actual,
	)
}

func TestProgressive_AddCondition(t *testing.T) {
	progressive := NewProgressive[int]("Progressive")
	progressive.Update(NewConstant(0))

	criteria := NewMockTypedValue[bool](t)

	closeIfFn := progressive.AddCondition(criteria)
	curCondition := progressive.curCondition

	assert.Equal(
		t, &Condition{
			criteria:       criteria,
			prevCondition:  nil,
			openIfStageIdx: 1,
		}, curCondition,
	)

	progressive.Update(NewConstant(1))
	progressive.Update(NewConstant(2))

	closeIfFn()

	assert.Equal(t, 1, curCondition.openIfStageIdx)
	assert.Equal(t, 2, curCondition.closeIfStageIdx)

	closeIfFn = progressive.AddCondition(criteria)

	assert.Equal(
		t, &Condition{
			criteria:       criteria,
			prevCondition:  curCondition,
			openIfStageIdx: 3,
		}, progressive.curCondition,
	)
}

func TestProgressive_DoAnchor(t *testing.T) {
	progressive := NewProgressive[int]("Progressive")

	progressive.Update(NewConstant(0))
	simple := progressive.DoAnchor("Something")

	assert.Equal(t, "Something", simple.GetName())
	assert.Equal(t, 0, simple.GetTypedValue())

	progressive.Update(NewConstant(1))

	assert.Equal(t, 0, simple.GetTypedValue())
}

func TestProgressive_Update(t *testing.T) {
	progressive := NewProgressive[int]("Progressive")

	var otherProgressive *Progressive[int]

	progressive.Update(otherProgressive)

	assert.Equal(
		t, &Stage[int]{
			self:      progressive,
			idx:       0,
			prevStage: nil,
			value:     0,
			source:    ZeroSimple[int]("NilProgressiveUpdate"),
		}, progressive.curStage,
	)

	prevStage := progressive.curStage

	otherProgressive = NewProgressive[int]("Other")
	otherProgressive.Update(NewConstant(1))

	progressive.Update(otherProgressive)

	assert.Equal(
		t, &Stage[int]{
			self:      progressive,
			idx:       1,
			prevStage: prevStage,
			value:     1,
			source:    otherProgressive.getStage(0),
		}, progressive.curStage,
	)

	prevStage = progressive.curStage

	otherProgressive.Update(NewConstant(2))
	progressive.Update(otherProgressive)

	assert.Equal(
		t, &Stage[int]{
			self:      progressive,
			idx:       2,
			prevStage: prevStage,
			value:     2,
			source:    otherProgressive.getStage(1),
		}, progressive.curStage,
	)

	prevStage = progressive.curStage

	simple := NewSimple("Something", 3)
	progressive.Update(simple)

	assert.Equal(
		t, &Stage[int]{
			self:      progressive,
			idx:       3,
			prevStage: prevStage,
			value:     3,
			source:    simple,
		}, progressive.curStage,
	)
}

func TestProgressive_GetSnapshot(t *testing.T) {
	var progressive *Progressive[int]

	actual := progressive.getSnapshot()

	assert.Nil(t, actual)

	progressive = NewProgressive[int]("Progressive")

	actual = progressive.getSnapshot()

	assert.Equal(t, "DefaultProgressive", actual.source.GetName())
	assert.Equal(t, 0, actual.value)

	simple := NewSimple("Something", 3)
	progressive.Update(simple)

	actual = progressive.getSnapshot()

	assert.Equal(t, "Something", actual.source.GetName())
	assert.Equal(t, 3, actual.value)
}

func TestProgressive_GetStage(t *testing.T) {
	progressive := NewProgressive[int]("Progressive")
	progressive.Update(NewConstant(1))
	progressive.Update(NewConstant(2))
	progressive.Update(NewConstant(3))

	stage1 := progressive.getStage(0)
	assert.Equal(t, 0, stage1.idx)

	stage2 := progressive.getStage(1)
	assert.Equal(t, 1, stage2.idx)

	stage3 := progressive.getStage(2)
	assert.Equal(t, 2, stage3.idx)
}

func TestProgressive_GetCurrentStageIdx(t *testing.T) {
	progressive := NewProgressive[int]("Progressive")

	actual := progressive.getCurrentStageIdx()

	assert.Equal(t, -1, actual)

	progressive.Update(NewConstant(1))

	actual = progressive.getCurrentStageIdx()

	assert.Equal(t, 0, actual)
}

func TestProgressive_MarshalJSON(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil Progressive",
			test: func(t *testing.T) {
				var nilProgressive *Progressive[int]

				actual, err := nilProgressive.MarshalJSON()
				assert.Nil(t, actual)
				assert.Nil(t, err)
			},
		},
		{
			desc: "non-nil Progressive",
			test: func(t *testing.T) {
				progressive := NewProgressive[int]("Progressive")
				progressive.Update(NewConstant(1))

				criteria := NewConstant[bool](true)
				closeIfFn := progressive.AddCondition(criteria)

				progressive.Update(NewSimple("Something", 2))
				closeIfFn()

				progressive.Tag(
					Tag{
						Name:  "Tag",
						Value: "MyTag",
					},
				)

				progressive.Tag(NewTagFrom(NewSimple("Text", "very long text")))

				wanted := `{"Source":"progressive_calculation","Stages":[{"Value":"1","Formula":{"Category":"AssignStatic","Operation":"1"}},{"Value":"2","Formula":{"Category":"AssignVariable","Operands":[{"Name":"Something"}]}}],"Conditions":{"1":[{"Formula":{"Category":"AssignStatic","Operation":"true"},"CloseIfStageIdx":1}]},"Tags":{"Tag":{"Value":"MyTag"},"Text":{"Value":"very long text","IsValue":true}}}`

				actual, err := progressive.MarshalJSON()
				assert.Equal(t, wanted, string(actual))
				assert.Nil(t, err)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
