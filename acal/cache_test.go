package acal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestValueCache_Take(t *testing.T) {
	scenarios := []struct {
		desc   string
		setup  func(*mockValueWithFormula, valueCache)
		expect func(*testing.T, *mockValueWithFormula, valueCache)
		taken  bool
	}{
		{
			desc: "Cache is empty OR doesn't contain other variables with the same name",
			expect: func(t *testing.T, aValMock *mockValueWithFormula, cache valueCache) {
				assert.Equal(t, cache.values[aValMock.GetName()], []Value{aValMock}, "cache should contain the newly added value")
			},
			taken: true,
		},
		{
			desc: "Cache already contains the value",
			setup: func(aValMock *mockValueWithFormula, cache valueCache) {
				cache.Take(aValMock)
			},
			expect: func(t *testing.T, aValMock *mockValueWithFormula, cache valueCache) {
				assert.Equal(t, cache.values[aValMock.GetName()], []Value{aValMock}, "cache should contain the newly added value")
			},
			taken: false,
		},
		{
			desc: "Duplicate value has no alias, 1 duplicate",
			setup: func(aValMock *mockValueWithFormula, cache valueCache) {
				aValMock.On("GetAlias").Return("").Once()
				aValMock.On("SetAlias", "TestName2").Once()

				duplicateMock := &mockValueWithFormula{}
				duplicateMock.On("GetName").Return("TestName").Maybe()
				cache.Take(duplicateMock)
			},
			expect: func(t *testing.T, aValMock *mockValueWithFormula, cache valueCache) {
				assert.Equal(t, len(cache.values[aValMock.GetName()]), 2, "cache should contain 2 values under TestName")
			},
			taken: true,
		},
		{
			desc: "Duplicate value has no alias, 2 duplicates",
			setup: func(aValMock *mockValueWithFormula, cache valueCache) {
				aValMock.On("GetAlias").Return("").Once()
				aValMock.On("SetAlias", "TestName3").Once()

				duplicateMock := &mockValueWithFormula{}
				duplicateMock.On("GetName").Return("TestName").Maybe()
				cache.Take(duplicateMock)

				duplicateMock2 := &mockValueWithFormula{}
				duplicateMock2.On("GetName").Return("TestName").Maybe()
				duplicateMock2.On("GetAlias").Return("").Once()
				duplicateMock2.On("SetAlias", "TestName2").Once()
				cache.Take(duplicateMock2)
			},
			expect: func(t *testing.T, aValMock *mockValueWithFormula, cache valueCache) {
				assert.Equal(t, len(cache.values[aValMock.GetName()]), 3, "cache should contain 3 values under TestName")
			},
			taken: true,
		},
		{
			desc: "Duplicate value has alias",
			setup: func(aValMock *mockValueWithFormula, cache valueCache) {
				aValMock.On("GetAlias").Return("TestAlias").Once()
				aValMock.On("SetAlias", "TestName2").Maybe()

				duplicateMock := &mockValueWithFormula{}
				duplicateMock.On("GetName").Return("TestName").Maybe()
				cache.Take(duplicateMock)
			},
			expect: func(t *testing.T, aValMock *mockValueWithFormula, cache valueCache) {
				aValMock.AssertNotCalled(t, "SetAlias", "TestName2")
				assert.Equal(t, len(cache.values[aValMock.GetName()]), 2, "cache should contain 2 values under TestName")
			},
			taken: true,
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				cache := NewValueCache()

				aValMock := &mockValueWithFormula{}
				aValMock.On("GetName").Return("TestName").Maybe()

				if sc.setup != nil {
					sc.setup(aValMock, cache.(valueCache))
				}

				taken := cache.Take(aValMock)
				assert.Equal(t, sc.taken, taken, "value should be taken, expectation: %v", sc.taken)
				mock.AssertExpectationsForObjects(t, aValMock)
				sc.expect(t, aValMock, cache.(valueCache))
			},
		)
	}
}

func TestValueCache_Flatten(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "Unique identities",
			test: func(t *testing.T) {
				valueOpsMock, cleanup := MockValueOps()
				defer cleanup()

				aValMock1 := &mockValueWithFormula{}
				aValMock2 := &mockValueWithFormula{}

				valueOpsMock.On("Identify", aValMock1).Return("TestIdentity1").Once()
				valueOpsMock.On("Identify", aValMock2).Return("TestIdentity2").Once()

				cache := &valueCache{
					values: map[string][]Value{
						"TestName1": {aValMock1},
						"TestName2": {aValMock2},
					},
				}

				want := map[string]Value{
					"TestIdentity1": aValMock1,
					"TestIdentity2": aValMock2,
				}

				actual := cache.Flatten()

				assert.Equal(t, want, actual, "flatten map should meet expectation")
				mock.AssertExpectationsForObjects(t, aValMock1, aValMock2, valueOpsMock)
			},
		},
		{
			desc: "Duplicated identities, different names",
			test: func(t *testing.T) {
				valueOpsMock, cleanup := MockValueOps()
				defer cleanup()

				aValMock1 := &mockValueWithFormula{}
				aValMock1.On("GetName").Return("TestName1").Once()

				aValMock2 := &mockValueWithFormula{}
				aValMock2.On("GetName").Return("TestName2").Once()
				aValMock2.On("SetAlias", "TestIdentity_2").Once()

				valueOpsMock.On("Identify", mock.Anything).Return("TestIdentity").Twice()

				cache := &valueCache{
					values: map[string][]Value{
						"TestName1": {aValMock1},
						"TestName2": {aValMock2},
					},
				}

				want := map[string]Value{
					"TestIdentity":   aValMock1,
					"TestIdentity_2": aValMock2,
				}

				actual := cache.Flatten()

				assert.Equal(t, want, actual, "flatten map should meet expectation")
				mock.AssertExpectationsForObjects(t, aValMock1, aValMock2, valueOpsMock)
			},
		},
		{
			desc: "Duplicated identities, same name",
			test: func(t *testing.T) {
				valueOpsMock, cleanup := MockValueOps()
				defer cleanup()

				aValMock1 := &mockValueWithFormula{}
				aValMock1.On("GetName").Return("TestName").Once()

				aValMock2 := &mockValueWithFormula{}
				aValMock2.On("GetName").Return("TestName").Once()

				valueOpsMock.On("Identify", mock.Anything).Return("TestIdentity").Twice()

				cache := &valueCache{
					values: map[string][]Value{
						"TestName1": {aValMock1},
						"TestName2": {aValMock2},
					},
				}

				want := map[string]Value{
					"TestIdentity": aValMock1,
				}

				actual := cache.Flatten()

				assert.Equal(t, want, actual, "flatten map should meet expectation")
				mock.AssertExpectationsForObjects(t, aValMock1, aValMock2, valueOpsMock)
			},
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(sc.desc, sc.test)
	}
}
