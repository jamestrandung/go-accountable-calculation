package acal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNil(t *testing.T) {
	scenarios := []struct {
		desc  string
		value any
		want  bool
	}{
		{
			desc:  "nil any",
			value: nil,
			want:  true,
		},
		{
			desc:  "nil pointer",
			value: (*float64)(nil),
			want:  true,
		},
		{
			desc:  "nil map",
			value: map[string]int64(nil),
			want:  true,
		},
		{
			desc:  "nil func",
			value: (func())(nil),
			want:  true,
		},
		{
			desc:  "nil chan",
			value: (chan string)(nil),
			want:  true,
		},
		{
			desc:  "nil slice",
			value: []bool(nil),
			want:  true,
		},
		{
			desc:  "nil interface",
			value: Value(nil),
			want:  true,
		},
		{
			desc:  "nil concrete interface implementation",
			value: fmt.Stringer((*Source)(nil)),
			want:  true,
		},
		{
			desc:  "non-nil interface",
			value: fmt.Stringer(SourceHardcode),
			want:  false,
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				actual := isNil(sc.value)

				assert.Equal(t, sc.want, actual)
			},
		)
	}
}
