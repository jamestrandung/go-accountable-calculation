package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/my-shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemote_GetTypedValue(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Remote is nil",
			test: func(t *testing.T) {
				r := Remote{
					Remote: nil,
				}

				actual := r.GetTypedValue()
				assert.Equal(t, decimal.Decimal{}, actual)
			},
		},
		{
			desc: "inner Remote is not nil",
			test: func(t *testing.T) {
				r := MakeRemote("Remote", decimal.NewFromInt(10), "RemoteField", "RemoteKey")

				actual := r.GetTypedValue()
				assert.Equal(t, decimal.NewFromInt(10), actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestRemote_Anchor(t *testing.T) {
	r := MakeRemote("Remote", decimal.NewFromInt(10), "RemoteField", "RemoteKey")

	actual := r.Anchor("Simple")

	assert.Equal(t, "Simple", actual.GetName())
	assert.Equal(t, decimal.NewFromInt(10), actual.Decimal())
	assert.Equal(
		t,
		`{"Remote":{"Value":"10","Source":"remote_calculation","DependentField":"RemoteField","LogKey":"RemoteKey"},"Simple":{"Value":"10","Source":"static_calculation","Formula":{"Category":"AssignVariable","Operands":[{"Name":"Remote"}]}}}`,
		acal.ToString(actual),
	)
}

func TestRemote_Decimal(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Remote is nil",
			test: func(t *testing.T) {
				r := Remote{
					Remote: nil,
				}

				actual := r.Decimal()
				assert.Equal(t, decimal.Decimal{}, actual)
			},
		},
		{
			desc: "inner Remote is not nil",
			test: func(t *testing.T) {
				r := MakeRemote("Remote", decimal.NewFromInt(10), "RemoteField", "RemoteKey")

				actual := r.Decimal()
				assert.Equal(t, decimal.NewFromInt(10), actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestRemote_Float(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "inner Remote is nil",
			test: func(t *testing.T) {
				r := Remote{
					Remote: nil,
				}

				actual := r.Float()
				assert.Equal(t, 0.0, actual)
			},
		},
		{
			desc: "inner Remote is not nil",
			test: func(t *testing.T) {
				r := MakeRemote("Remote", decimal.NewFromInt(10), "RemoteField", "RemoteKey")

				actual := r.Float()
				assert.Equal(t, 10.0, actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
