package acal

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockMarshalOps can be used in tests to mock IMarshalOps.
func MockMarshalOps(t *testing.T) (*MockIMarshalOps, func()) {
	old := marshalOps
	mock := NewMockIMarshalOps(t)

	marshalOps = mock
	return mock, func() {
		marshalOps = old
	}
}

func TestMarshalOpsImpl_MarshalJSON(t *testing.T) {
	mockValueOps, cleanup := MockValueOps(t)
	defer cleanup()

	defer func(original func(v any) ([]byte, error)) {
		// restore (after test)
		jsonMarshalFn = original
	}(jsonMarshalFn)

	testMap := make(map[string]Value)

	mockCache := NewMockIValueCache(t)
	mockCache.On("Flatten").Return(testMap).Once()

	anchoredMockValue := newMockValueWithAllFeatures(t)
	anchoredMockValue.On("SetAlias", UnknownValueName).Maybe()
	anchoredMockValue.On("ExtractValues", mock.Anything).Return(mockCache).Once()

	unanchoredMockValue1 := newMockValueWithAllFeatures(t)
	unanchoredMockValue1.On("SetAlias", UnknownValueName).Once()
	unanchoredMockValue1.On("ExtractValues", mock.Anything).Return(mockCache).Once()

	unanchoredMockValue2 := newMockValueWithAllFeatures(t)
	unanchoredMockValue2.On("SetAlias", UnknownValueName).Once()
	unanchoredMockValue2.On("ExtractValues", mock.Anything).Return(mockCache).Once()

	mockValueOps.On("IsNilValue", nil).Return(true).Once()
	mockValueOps.On("IsNilValue", anchoredMockValue).Return(false).Once()
	mockValueOps.On("IsNilValue", unanchoredMockValue1).Return(false).Once()
	mockValueOps.On("IsNilValue", unanchoredMockValue2).Return(false).Once()

	mockValueOps.On("HasIdentity", anchoredMockValue).Return(true).Once()
	mockValueOps.On("HasIdentity", unanchoredMockValue1).Return(false).Once()
	mockValueOps.On("HasIdentity", unanchoredMockValue2).Return(false).Once()

	jsonMarshalFn = func(v any) ([]byte, error) {
		assert.Equal(t, testMap, v)

		return json.Marshal(v)
	}

	ops := marshalOpsImpl{}

	_, err := ops.MarshalJSON(nil, anchoredMockValue, unanchoredMockValue1, unanchoredMockValue2)
	assert.Nil(t, err, "error should be nil")

	anchoredMockValue.AssertNotCalled(t, "SetAlias", UnknownValueName)
}

func TestMarshalOpsImpl_MarshalJSONByFields(t *testing.T) {
	defer func(original func(v any) ([]byte, error)) {
		// restore (after test)
		jsonMarshalFn = original
	}(jsonMarshalFn)

	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil data",
			test: func(t *testing.T) {
				data := (*struct{})(nil)

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("cannot marshal nil data"), err)
			},
		},
		{
			desc: "data is Value",
			test: func(t *testing.T) {
				data := newMockValueWithAllFeatures(t)

				marshalOpsMock, cleanup := MockMarshalOps(t)
				defer cleanup()

				marshalOpsMock.On("MarshalJSON", mock.Anything).
					Return(
						nil, func(values ...Value) error {
							assert.Equal(t, 1, len(values))
							assert.Equal(t, data, values[0])

							return fmt.Errorf("dummy marshalJSON was executed")
						},
					).Once()

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("dummy marshalJSON was executed"), err)
			},
		},
		{
			desc: "data is not nil but also not a struct",
			test: func(t *testing.T) {
				data := 5

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("cannot marshal non-struct data"), err)
			},
		},
		{
			desc: "data is a struct with an unexported field",
			test: func(t *testing.T) {
				data := struct {
					unexportedField TypedValue[float64]
				}{NewMockTypedValue[float64](t)}

				jsonMarshalFn = func(v any) ([]byte, error) {
					m, ok := v.(map[string]any)
					assert.Equal(t, true, ok, "the object to be marshalled should be a map[string]any")
					assert.Equal(t, 0, len(m))

					return nil, fmt.Errorf("dummy jsonMarshalFn was executed")
				}

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("dummy jsonMarshalFn was executed"), err)
			},
		},
		{
			desc: "data is a struct with an exported field that is not Value",
			test: func(t *testing.T) {
				data := struct {
					ExportedField float64
				}{5.0}

				jsonMarshalFn = func(v any) ([]byte, error) {
					m, ok := v.(map[string]any)
					assert.Equal(t, true, ok, "the object to be marshalled should be a map[string]any")
					assert.Equal(t, 1, len(m))
					assert.Equal(t, 5.0, m["ExportedField"])

					return nil, fmt.Errorf("dummy jsonMarshalFn was executed")
				}

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("dummy jsonMarshalFn was executed"), err)
			},
		},
		{
			desc: "data is a struct with an exported field that is an Value with nil value",
			test: func(t *testing.T) {
				data := struct {
					ExportedField *MockTypedValue[float64]
				}{}

				jsonMarshalFn = func(v any) ([]byte, error) {
					m, ok := v.(map[string]any)
					assert.Equal(t, true, ok, "the object to be marshalled should be a map[string]any")
					assert.Equal(t, 0, len(m))

					return nil, fmt.Errorf("dummy jsonMarshalFn was executed")
				}

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", (*MockTypedValue[float64])(nil)).Return(true).Once()

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("dummy jsonMarshalFn was executed"), err)
			},
		},
		{
			desc: "data is a struct with an exported field that is a non-nil un-anchored Value",
			test: func(t *testing.T) {
				mockValue1 := newMockValueWithAllFeatures(t)
				mockValue2 := newMockValueWithAllFeatures(t)

				mockCache := NewMockIValueCache(t)
				mockCache.On("Flatten").Return(
					map[string]Value{
						"Unknown": mockValue1,
					},
				).Once()

				mockValue1.On("SetAlias", UnknownValueName).Once()
				mockValue1.On("ExtractValues", mock.Anything).Return(mockCache).Once()

				mockValue2.On("SetAlias", UnknownValueName).Once()
				mockValue2.On("ExtractValues", mock.Anything).Return(mockCache).Once()

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue1).Return(false).Once()
				mockValueOps.On("HasIdentity", mockValue1).Return(false).Once()

				mockValueOps.On("IsNilValue", mockValue2).Return(false).Once()
				mockValueOps.On("HasIdentity", mockValue2).Return(false).Once()

				data := struct {
					ExportedField1 Value
					ExportedField2 Value
				}{mockValue1, mockValue2}

				jsonMarshalFn = func(v any) ([]byte, error) {
					m, ok := v.(map[string]any)
					assert.Equal(t, true, ok, "the object to be marshalled should be a map[string]any")
					assert.Equal(t, 1, len(m))
					assert.Equal(t, mockValue1, m["Unknown"])

					return nil, fmt.Errorf("dummy jsonMarshalFn was executed")
				}

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("dummy jsonMarshalFn was executed"), err)
			},
		},
		{
			desc: "data is a struct with an exported field that is a non-nil anchored Value",
			test: func(t *testing.T) {
				mockValue := newMockValueWithAllFeatures(t)

				mockCache := NewMockIValueCache(t)
				mockCache.On("Flatten").Return(
					map[string]Value{
						"TestName": mockValue,
					},
				).Once()

				mockValue.On("SetAlias", UnknownValueName).Maybe()
				mockValue.On("ExtractValues", mock.Anything).Return(mockCache).Once()

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue).Return(false).Once()
				mockValueOps.On("HasIdentity", mockValue).Return(true).Once()

				data := struct {
					ExportedField Value
				}{mockValue}

				jsonMarshalFn = func(v any) ([]byte, error) {
					m, ok := v.(map[string]any)
					assert.Equal(t, true, ok, "the object to be marshalled should be a map[string]any")
					assert.Equal(t, 1, len(m))
					assert.Equal(t, mockValue, m["TestName"])

					return nil, fmt.Errorf("dummy jsonMarshalFn was executed")
				}

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("dummy jsonMarshalFn was executed"), err)
				mockValue.AssertNotCalled(t, "SetAlias", UnknownValueName)
			},
		},
		{
			desc: "data is a struct with an exported field that is a non-nil anchored Value whose alias is the same with another field in the struct",
			test: func(t *testing.T) {
				mockValue := newMockValueWithAllFeatures(t)

				mockCache := NewMockIValueCache(t)
				mockCache.On("Flatten").Return(
					map[string]Value{
						"TestName": mockValue,
					},
				).Once()

				mockValue.On("SetAlias", "TestName.2").Once()
				mockValue.On("ExtractValues", mock.Anything).Return(mockCache).Once()

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue).Return(false).Once()
				mockValueOps.On("HasIdentity", mockValue).Return(true).Once()

				data := struct {
					TestName      float64
					ExportedField Value
				}{5.0, mockValue}

				jsonMarshalFn = func(v any) ([]byte, error) {
					m, ok := v.(map[string]any)
					assert.Equal(t, true, ok, "the object to be marshalled should be a map[string]any")
					assert.Equal(t, 2, len(m))
					assert.Equal(t, mockValue, m["TestName.2"])

					return nil, fmt.Errorf("dummy jsonMarshalFn was executed")
				}

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("dummy jsonMarshalFn was executed"), err)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
