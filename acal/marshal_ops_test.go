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
	valueOpsMock, cleanup := MockValueOps(t)
	defer cleanup()

	defer func(original func(v any) ([]byte, error)) {
		// restore (after test)
		jsonMarshalFn = original
	}(jsonMarshalFn)

	testMap := make(map[string]Value)

	cacheMock := NewMockIValueCache(t)
	cacheMock.On("Flatten").Return(testMap).Once()

	anchordedAValMock := newMockValueWithAllFeatures(t)
	anchordedAValMock.On("SetAlias", UnknownValueName).Maybe()
	anchordedAValMock.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

	unanchoredAValMock1 := newMockValueWithAllFeatures(t)
	unanchoredAValMock1.On("SetAlias", UnknownValueName).Once()
	unanchoredAValMock1.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

	unanchoredAValMock2 := newMockValueWithAllFeatures(t)
	unanchoredAValMock2.On("SetAlias", UnknownValueName).Once()
	unanchoredAValMock2.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

	valueOpsMock.On("IsNilValue", nil).Return(true).Once()
	valueOpsMock.On("IsNilValue", anchordedAValMock).Return(false).Once()
	valueOpsMock.On("IsNilValue", unanchoredAValMock1).Return(false).Once()
	valueOpsMock.On("IsNilValue", unanchoredAValMock2).Return(false).Once()

	valueOpsMock.On("HasIdentity", anchordedAValMock).Return(true).Once()
	valueOpsMock.On("HasIdentity", unanchoredAValMock1).Return(false).Once()
	valueOpsMock.On("HasIdentity", unanchoredAValMock2).Return(false).Once()

	jsonMarshalFn = func(v any) ([]byte, error) {
		assert.Equal(t, testMap, v)

		return json.Marshal(v)
	}

	ops := marshalOpsImpl{}

	_, err := ops.MarshalJSON(nil, anchordedAValMock, unanchoredAValMock1, unanchoredAValMock2)
	assert.Nil(t, err, "error should be nil")

	anchordedAValMock.AssertNotCalled(t, "SetAlias", UnknownValueName)
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

				valueOpsMock, cleanup := MockValueOps(t)
				defer cleanup()

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
				aValMock1 := newMockValueWithAllFeatures(t)
				aValMock2 := newMockValueWithAllFeatures(t)

				cacheMock := NewMockIValueCache(t)
				cacheMock.On("Flatten").Return(
					map[string]Value{
						"Unknown": aValMock1,
					},
				).Once()

				aValMock1.On("SetAlias", UnknownValueName).Once()
				aValMock1.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

				aValMock2.On("SetAlias", UnknownValueName).Once()
				aValMock2.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

				valueOpsMock, cleanup := MockValueOps(t)
				defer cleanup()

				valueOpsMock.On("IsNilValue", aValMock1).Return(false).Once()
				valueOpsMock.On("HasIdentity", aValMock1).Return(false).Once()

				valueOpsMock.On("IsNilValue", aValMock2).Return(false).Once()
				valueOpsMock.On("HasIdentity", aValMock2).Return(false).Once()

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
			},
		},
		{
			desc: "data is a struct with an exported field that is a non-nil anchored Value",
			test: func(t *testing.T) {
				aValMock := newMockValueWithAllFeatures(t)

				cacheMock := NewMockIValueCache(t)
				cacheMock.On("Flatten").Return(
					map[string]Value{
						"TestName": aValMock,
					},
				).Once()

				aValMock.On("SetAlias", UnknownValueName).Maybe()
				aValMock.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

				valueOpsMock, cleanup := MockValueOps(t)
				defer cleanup()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("HasIdentity", aValMock).Return(true).Once()

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
			},
		},
		{
			desc: "data is a struct with an exported field that is a non-nil anchored Value whose alias is the same with another field in the struct",
			test: func(t *testing.T) {
				aValMock := newMockValueWithAllFeatures(t)

				cacheMock := NewMockIValueCache(t)
				cacheMock.On("Flatten").Return(
					map[string]Value{
						"TestName": aValMock,
					},
				).Once()

				aValMock.On("SetAlias", "TestName.2").Once()
				aValMock.On("ExtractValues", mock.Anything).Return(cacheMock).Once()

				valueOpsMock, cleanup := MockValueOps(t)
				defer cleanup()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("HasIdentity", aValMock).Return(true).Once()

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
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
