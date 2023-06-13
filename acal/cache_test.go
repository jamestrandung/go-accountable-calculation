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
			expect: func(t *testing.T, mockValue *mockValueWithFormula, cache valueCache) {
				assert.Equal(t, cache.values[mockValue.GetName()], []Value{mockValue}, "cache should contain the newly added value")
			},
			taken: true,
		},
		{
			desc: "Cache already contains the value",
			setup: func(mockValue *mockValueWithFormula, cache valueCache) {
				cache.Take(mockValue)
			},
			expect: func(t *testing.T, mockValue *mockValueWithFormula, cache valueCache) {
				assert.Equal(t, cache.values[mockValue.GetName()], []Value{mockValue}, "cache should contain the newly added value")
			},
			taken: false,
		},
		{
			desc: "Duplicate value has no alias, 1 duplicate",
			setup: func(mockValue *mockValueWithFormula, cache valueCache) {
				mockValue.On("GetAlias").Return("").Once()
				mockValue.On("SetAlias", "TestName2").Once()

				mockDuplicate := newMockValueWithFormula(t)
				mockDuplicate.On("GetName").Return("TestName").Maybe()
				cache.Take(mockDuplicate)
			},
			expect: func(t *testing.T, mockValue *mockValueWithFormula, cache valueCache) {
				assert.Equal(t, len(cache.values[mockValue.GetName()]), 2, "cache should contain 2 values under TestName")
			},
			taken: true,
		},
		{
			desc: "Duplicate value has no alias, 2 duplicates",
			setup: func(mockValue *mockValueWithFormula, cache valueCache) {
				mockValue.On("GetAlias").Return("").Once()
				mockValue.On("SetAlias", "TestName3").Once()

				mockDuplicate := newMockValueWithFormula(t)
				mockDuplicate.On("GetName").Return("TestName").Maybe()
				cache.Take(mockDuplicate)

				mockDuplicate2 := newMockValueWithFormula(t)
				mockDuplicate2.On("GetName").Return("TestName").Maybe()
				mockDuplicate2.On("GetAlias").Return("").Once()
				mockDuplicate2.On("SetAlias", "TestName2").Once()
				cache.Take(mockDuplicate2)
			},
			expect: func(t *testing.T, mockValue *mockValueWithFormula, cache valueCache) {
				assert.Equal(t, len(cache.values[mockValue.GetName()]), 3, "cache should contain 3 values under TestName")
			},
			taken: true,
		},
		{
			desc: "Duplicate value has alias",
			setup: func(mockValue *mockValueWithFormula, cache valueCache) {
				mockValue.On("GetAlias").Return("TestAlias").Once()
				mockValue.On("SetAlias", "TestName2").Maybe()

				mockDuplicate := newMockValueWithFormula(t)
				mockDuplicate.On("GetName").Return("TestName").Maybe()
				cache.Take(mockDuplicate)
			},
			expect: func(t *testing.T, mockValue *mockValueWithFormula, cache valueCache) {
				mockValue.AssertNotCalled(t, "SetAlias", "TestName2")
				assert.Equal(t, len(cache.values[mockValue.GetName()]), 2, "cache should contain 2 values under TestName")
			},
			taken: true,
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				cache := NewValueCache()

				mockValue := newMockValueWithFormula(t)
				mockValue.On("GetName").Return("TestName").Maybe()

				if sc.setup != nil {
					sc.setup(mockValue, cache.(valueCache))
				}

				taken := cache.Take(mockValue)
				assert.Equal(t, sc.taken, taken, "value should be taken, expectation: %v", sc.taken)
				sc.expect(t, mockValue, cache.(valueCache))
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
				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValue1 := newMockValueWithFormula(t)
				mockValue2 := newMockValueWithFormula(t)

				mockValueOps.On("Identify", mockValue1).Return("TestIdentity1").Once()
				mockValueOps.On("Identify", mockValue2).Return("TestIdentity2").Once()

				cache := &valueCache{
					values: map[string][]Value{
						"TestName1": {mockValue1},
						"TestName2": {mockValue2},
					},
				}

				want := map[string]Value{
					"TestIdentity1": mockValue1,
					"TestIdentity2": mockValue2,
				}

				actual := cache.Flatten()

				assert.Equal(t, want, actual, "flatten map should meet expectation")
			},
		},
		{
			desc: "Duplicated identities, different names",
			test: func(t *testing.T) {
				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValue1 := newMockValueWithFormula(t)
				mockValue1.On("GetName").Return("TestName1").Once()
				mockValue1.On("SetAlias", "TestIdentity_2").Maybe()

				mockValue2 := newMockValueWithFormula(t)
				mockValue2.On("GetName").Return("TestName2").Once()
				mockValue2.On("SetAlias", "TestIdentity_2").Maybe()

				mockValueOps.On("Identify", mock.Anything).Return("TestIdentity").Twice()

				cache := &valueCache{
					values: map[string][]Value{
						"TestName1": {mockValue1},
						"TestName2": {mockValue2},
					},
				}

				actual := cache.Flatten()

				case1 := actual["TestIdentity"] == mockValue1 && actual["TestIdentity_2"] == mockValue2
				case2 := actual["TestIdentity"] == mockValue2 && actual["TestIdentity_2"] == mockValue1
				assert.True(t, case1 || case2)
			},
		},
		{
			desc: "Duplicated identities, same name",
			test: func(t *testing.T) {
				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValue1 := newMockValueWithFormula(t)
				mockValue1.On("GetName").Return("TestName").Once()

				mockValue2 := newMockValueWithFormula(t)
				mockValue2.On("GetName").Return("TestName").Once()

				mockValueOps.On("Identify", mock.Anything).Return("TestIdentity").Twice()

				cache := &valueCache{
					values: map[string][]Value{
						"TestName1": {mockValue1},
						"TestName2": {mockValue2},
					},
				}

				want := map[string]Value{
					"TestIdentity": mockValue1,
				}

				actual := cache.Flatten()

				assert.Equal(t, want, actual, "flatten map should meet expectation")
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
