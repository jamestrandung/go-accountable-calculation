package float

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatFloat(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "float exceeding digit threshold",
			test: func(t *testing.T) {
				d := decimal.NewFromFloat(912345678.1234567)

				actual := FormatFn(d)
				assert.Equal(t, "9.12345678123457e+08", actual)
			},
		},
		{
			desc: "float below digit threshold",
			test: func(t *testing.T) {
				d := decimal.NewFromFloat(91234567.1234567)

				actual := FormatFn(d)
				assert.Equal(t, "91234567.123457", actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
