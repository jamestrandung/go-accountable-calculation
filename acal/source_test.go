package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSource_Apply(t *testing.T) {
	var dummySource Source = "dummy"

	sourcerMock1 := newMockSourcer(t)
	sourcerMock1.On("From", dummySource).Once()

	sourcerMock2 := newMockSourcer(t)
	sourcerMock2.On("From", dummySource).Once()

	dummySource.Apply(sourcerMock1, sourcerMock2)
}

func TestSource_String(t *testing.T) {
	var dummySource Source = "dummy"

	assert.Equal(t, "dummy", dummySource.String())
}
