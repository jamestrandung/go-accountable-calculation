package acal

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockMarshalOps can be used in tests to mock IMarshalOps.
func MockMarshalOps() (*MockIMarshalOps, func()) {
	old := marshalOps
	mock := &MockIMarshalOps{}

	marshalOps = mock
	return mock, func() {
		marshalOps = old
	}
}

func TestMarshalOpsImpl_MarshalJSON(t *testing.T) {
	valueOpsMock, cleanup := MockValueOps()
	defer cleanup()

	defer func(original func(v any) ([]byte, error)) {
		// restore (after test)
		jsonMarshalFn = original
	}(jsonMarshalFn)

	testMap := make(map[string]Value)

	cacheMock := &MockIValueCache{}
	cacheMock.On("Flatten").Return(testMap).Once()

	anchordedAValMock := &mockValueWithAllFeatures{}
	anchordedAValMock.On("SetAlias", UnknownValueName).Maybe()
	anchordedAValMock.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

	unanchoredAValMock1 := &mockValueWithAllFeatures{}
	unanchoredAValMock1.On("SetAlias", UnknownValueName).Once()
	unanchoredAValMock1.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

	unanchoredAValMock2 := &mockValueWithAllFeatures{}
	unanchoredAValMock2.On("SetAlias", UnknownValueName).Once()
	unanchoredAValMock2.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

	valueOpsMock.On("IsNilValue", nil).Return(true).Once()
	valueOpsMock.On("IsNilValue", anchordedAValMock).Return(false).Once()
	valueOpsMock.On("IsNilValue", unanchoredAValMock1).Return(false).Once()
	valueOpsMock.On("IsNilValue", unanchoredAValMock2).Return(false).Once()

	valueOpsMock.On("IsAnchored", anchordedAValMock).Return(true).Once()
	valueOpsMock.On("IsAnchored", unanchoredAValMock1).Return(false).Once()
	valueOpsMock.On("IsAnchored", unanchoredAValMock2).Return(false).Once()

	jsonMarshalFn = func(v any) ([]byte, error) {
		assert.Equal(t, testMap, v)

		return json.Marshal(v)
	}

	ops := marshalOpsImpl{}

	_, err := ops.MarshalJSON(nil, anchordedAValMock, unanchoredAValMock1, unanchoredAValMock2)
	assert.Nil(t, err, "error should be nil")

	anchordedAValMock.AssertNotCalled(t, "SetAlias", UnknownValueName)
	mock.AssertExpectationsForObjects(t, cacheMock, anchordedAValMock, unanchoredAValMock1, unanchoredAValMock2, valueOpsMock)
}

func TestMarshalOpsImpl_MarshalJSONByFields(t *testing.T) {
	valueOpsMock, cleanup := MockValueOps()
	defer cleanup()

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
				data := &mockValueWithAllFeatures{}

				marshalOpsMock, cleanup := MockMarshalOps()
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
				mock.AssertExpectationsForObjects(t, marshalOpsMock)
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
				}{&MockTypedValue[float64]{}}

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

				valueOpsMock.On("IsNilValue", (*MockTypedValue[float64])(nil)).Return(true).Once()

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("dummy jsonMarshalFn was executed"), err)
			},
		},
		{
			desc: "data is a struct with an exported field that is a non-nil un-anchored Value",
			test: func(t *testing.T) {
				aValMock1 := &mockValueWithAllFeatures{}
				aValMock2 := &mockValueWithAllFeatures{}

				cacheMock := &MockIValueCache{}
				cacheMock.On("Flatten").Return(
					map[string]Value{
						"Unknown": aValMock1,
					},
				).Once()

				aValMock1.On("SetAlias", UnknownValueName).Once()
				aValMock1.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

				aValMock2.On("SetAlias", UnknownValueName).Once()
				aValMock2.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

				valueOpsMock.On("IsNilValue", aValMock1).Return(false).Once()
				valueOpsMock.On("IsAnchored", aValMock1).Return(false).Once()

				valueOpsMock.On("IsNilValue", aValMock2).Return(false).Once()
				valueOpsMock.On("IsAnchored", aValMock2).Return(false).Once()

				data := struct {
					ExportedField1 Value
					ExportedField2 Value
				}{aValMock1, aValMock2}

				jsonMarshalFn = func(v any) ([]byte, error) {
					m, ok := v.(map[string]any)
					assert.Equal(t, true, ok, "the object to be marshalled should be a map[string]any")
					assert.Equal(t, 1, len(m))
					assert.Equal(t, aValMock1, m["Unknown"])

					return nil, fmt.Errorf("dummy jsonMarshalFn was executed")
				}

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("dummy jsonMarshalFn was executed"), err)
				mock.AssertExpectationsForObjects(t, cacheMock, aValMock1, aValMock2, valueOpsMock)
			},
		},
		{
			desc: "data is a struct with an exported field that is a non-nil anchored Value",
			test: func(t *testing.T) {
				aValMock := &mockValueWithAllFeatures{}

				cacheMock := &MockIValueCache{}
				cacheMock.On("Flatten").Return(
					map[string]Value{
						"TestName": aValMock,
					},
				).Once()

				aValMock.On("SetAlias", UnknownValueName).Maybe()
				aValMock.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("IsAnchored", aValMock).Return(true).Once()

				data := struct {
					ExportedField Value
				}{aValMock}

				jsonMarshalFn = func(v any) ([]byte, error) {
					m, ok := v.(map[string]any)
					assert.Equal(t, true, ok, "the object to be marshalled should be a map[string]any")
					assert.Equal(t, 1, len(m))
					assert.Equal(t, aValMock, m["TestName"])

					return nil, fmt.Errorf("dummy jsonMarshalFn was executed")
				}

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("dummy jsonMarshalFn was executed"), err)
				aValMock.AssertNotCalled(t, "SetAlias", UnknownValueName)
				mock.AssertExpectationsForObjects(t, cacheMock, aValMock, valueOpsMock)
			},
		},
		{
			desc: "data is a struct with an exported field that is a non-nil anchored Value whose alias is the same with another field in the struct",
			test: func(t *testing.T) {
				aValMock := &mockValueWithAllFeatures{}

				cacheMock := &MockIValueCache{}
				cacheMock.On("Flatten").Return(
					map[string]Value{
						"TestName": aValMock,
					},
				).Once()

				aValMock.On("SetAlias", "TestName.2").Once()
				aValMock.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("IsAnchored", aValMock).Return(true).Once()

				data := struct {
					TestName      float64
					ExportedField Value
				}{5.0, aValMock}

				jsonMarshalFn = func(v any) ([]byte, error) {
					m, ok := v.(map[string]any)
					assert.Equal(t, true, ok, "the object to be marshalled should be a map[string]any")
					assert.Equal(t, 2, len(m))
					assert.Equal(t, aValMock, m["TestName.2"])

					return nil, fmt.Errorf("dummy jsonMarshalFn was executed")
				}

				ops := marshalOpsImpl{}

				json, err := ops.MarshalJSONByFields(data)
				assert.Equal(t, []byte(nil), json)
				assert.Equal(t, fmt.Errorf("dummy jsonMarshalFn was executed"), err)
				mock.AssertExpectationsForObjects(t, cacheMock, aValMock, valueOpsMock)
			},
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				sc.test(t)
			},
		)
	}
}
