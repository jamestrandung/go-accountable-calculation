package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/my-shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProgressive_GetTypedValue(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Progressive is nil",
			test: func(t *testing.T) {
				var p Progressive

				actual := p.GetTypedValue()
				assert.Equal(t, decimal.Decimal{}, actual)
			},
		},
		{
			desc: "inner Progressive is not nil",
			test: func(t *testing.T) {
				p := MakeProgressive("Progressive")

				actual := p.GetTypedValue()
				assert.Equal(t, decimal.Decimal{}, actual)

				p.Update(Ten)

				actual = p.GetTypedValue()
				assert.Equal(t, decimal.NewFromInt(10), actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestProgressive_Anchor(t *testing.T) {
	p := MakeProgressive("Progressive")
	p.Update(Ten)
	p.Update(Thousand)

	actual := p.Anchor("Simple")

	assert.Equal(t, "Simple", actual.GetName())
	assert.Equal(t, decimal.NewFromInt(1000), actual.Decimal())
	assert.Equal(
		t,
		`{"Progressive":{"Source":"progressive_calculation","Stages":[{"Value":"10","Formula":{"Category":"AssignStatic","Operation":"10"}},{"Value":"1000","Formula":{"Category":"AssignStatic","Operation":"1000"}}]},"Simple":{"Value":"1000","Source":"static_calculation","Formula":{"Category":"AssignVariable","Operands":[{"Name":"Progressive","StageIdx":1}]}}}`,
		acal.ToString(actual),
	)
}

func TestProgressive_Decimal(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Progressive is nil",
			test: func(t *testing.T) {
				var p Progressive

				actual := p.Decimal()
				assert.Equal(t, decimal.Decimal{}, actual)
			},
		},
		{
			desc: "inner Progressive is not nil",
			test: func(t *testing.T) {
				p := MakeProgressive("Progressive")

				actual := p.GetTypedValue()
				assert.Equal(t, decimal.Decimal{}, actual)

				p.Update(Ten)

				actual = p.Decimal()
				assert.Equal(t, decimal.NewFromInt(10), actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestProgressive_Float(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Progressive is nil",
			test: func(t *testing.T) {
				var p Progressive

				actual := p.Float()
				assert.Equal(t, 0.0, actual)
			},
		},
		{
			desc: "inner Progressive is not nil",
			test: func(t *testing.T) {
				p := MakeProgressive("Progressive")

				actual := p.Float()
				assert.Equal(t, 0.0, actual)

				p.Update(Ten)

				actual = p.Float()
				assert.Equal(t, 10.0, actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
