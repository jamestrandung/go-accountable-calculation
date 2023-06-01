package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCondition_IsValidProgressiveCondition(t *testing.T) {
	scenarios := []struct {
		desc string
		cond *Condition
		want bool
	}{
		{
			desc: "openIfStageIdx > closeIfStageIdx",
			cond: &Condition{
				openIfStageIdx:  2,
				closeIfStageIdx: 1,
			},
			want: false,
		},
		{
			desc: "openIfStageIdx == closeIfStageIdx",
			cond: &Condition{
				openIfStageIdx:  2,
				closeIfStageIdx: 2,
			},
			want: true,
		},
		{
			desc: "openIfStageIdx < closeIfStageIdx",
			cond: &Condition{
				openIfStageIdx:  2,
				closeIfStageIdx: 3,
			},
			want: true,
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				assert.Equal(t, sc.want, sc.cond.isValidProgressiveCondition())
			},
		)
	}
}

func TestCondition_MarshalJSON(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "simple condition",
			test: func(t *testing.T) {
				valueOpsMock, cleanup := MockValueOps(t)
				defer cleanup()

				criteria := NewMockTypedValue[bool](t)
				valueOpsMock.On("DescribeValueAsFormula", criteria).
					Return(
						&SyntaxNode{
							category: OpCategoryAssignVariable,
							opDesc:   "ABC",
						},
					).
					Once()

				c := NewCondition(criteria)

				want := "{\"Formula\":{\"Category\":\"AssignVariable\",\"Operation\":\"ABC\"}}"

				actualJSON, err := c.MarshalJSON()
				assert.Equal(t, want, string(actualJSON))
				assert.Nil(t, err, "error should be nil")
			},
		},
		{
			desc: "progressive condition",
			test: func(t *testing.T) {
				valueOpsMock, cleanup := MockValueOps(t)
				defer cleanup()

				criteria := NewMockTypedValue[bool](t)
				valueOpsMock.On("DescribeValueAsFormula", criteria).
					Return(
						&SyntaxNode{
							category: OpCategoryAssignVariable,
							opDesc:   "ABC",
						},
					).
					Once()

				c := NewProgressiveCondition(criteria, nil, 2)
				c.setCloseIfStageIdx(5)

				want := "{\"Formula\":{\"Category\":\"AssignVariable\",\"Operation\":\"ABC\"},\"CloseIfStageIdx\":5}"

				actualJSON, err := c.MarshalJSON()
				assert.Equal(t, want, string(actualJSON))
				assert.Nil(t, err, "error should be nil")
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
