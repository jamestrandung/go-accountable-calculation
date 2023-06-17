package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/my-shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocal_GetTypedValue(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Local is nil",
			test: func(t *testing.T) {
				l := Local{
					Local: nil,
				}

				actual := l.GetTypedValue()
				assert.Equal(t, decimal.Decimal{}, actual)
			},
		},
		{
			desc: "inner Local is not nil",
			test: func(t *testing.T) {
				l := MakeLocal("Local", MakeSimpleFromFloat("Simple", 10))

				actual := l.GetTypedValue()
				assert.Equal(t, decimal.NewFromFloat(10), actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestLocal_Anchor(t *testing.T) {
	l := MakeLocal("Local", MakeSimpleFromFloat("Simple", 10))

	actual := l.Anchor("AnotherSimple")

	assert.Equal(t, "AnotherSimple", actual.GetName())
	assert.Equal(t, decimal.NewFromFloat(10), actual.Decimal())
	assert.Equal(
		t,
		`{"AnotherSimple":{"Value":"10","Source":"static_calculation","Formula":{"Category":"AssignVariable","Operands":[{"Name":"Local"}]}},"Local":{"Value":"10","Source":"dependent_calculation","DependentField":"Simple","Calculation":{"Simple":{"Value":"10","Source":"unknown"}}}}`,
		acal.ToString(actual),
	)
}

func TestLocal_Decimal(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Local is nil",
			test: func(t *testing.T) {
				l := Local{
					Local: nil,
				}

				actual := l.Decimal()
				assert.Equal(t, decimal.Decimal{}, actual)
			},
		},
		{
			desc: "inner Local is not nil",
			test: func(t *testing.T) {
				l := MakeLocal("Local", MakeSimpleFromFloat("Simple", 10))

				actual := l.Decimal()
				assert.Equal(t, decimal.NewFromFloat(10), actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestLocal_Float(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Local is nil",
			test: func(t *testing.T) {
				l := Local{
					Local: nil,
				}

				actual := l.Float()
				assert.Equal(t, 0.0, actual)
			},
		},
		{
			desc: "inner Local is not nil",
			test: func(t *testing.T) {
				l := MakeLocal("Local", MakeSimpleFromFloat("Simple", 10))

				actual := l.Float()
				assert.Equal(t, 10.0, actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
