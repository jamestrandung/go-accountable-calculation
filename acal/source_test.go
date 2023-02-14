package acal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSource_Apply(t *testing.T) {
	var dummySource Source = "dummy"

	sourcerMock1 := &mockSourcer{}
	sourcerMock1.On("From", dummySource).Once()

	sourcerMock2 := &mockSourcer{}
	sourcerMock2.On("From", dummySource).Once()

	dummySource.Apply(sourcerMock1, sourcerMock2)

	mock.AssertExpectationsForObjects(t, sourcerMock1, sourcerMock2)
}

func TestSource_String(t *testing.T) {
	var dummySource Source = "dummy"

	assert.Equal(t, "dummy", dummySource.String())
}
