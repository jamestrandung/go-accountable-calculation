package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSource_Apply(t *testing.T) {
	var dummySource Source = "dummy"

	mockSourcer1 := newMockSourcer(t)
	mockSourcer1.On("From", dummySource).Once()

	mockSourcer2 := newMockSourcer(t)
	mockSourcer2.On("From", dummySource).Once()

	dummySource.Apply(mockSourcer1, mockSourcer2)
}

func TestSource_String(t *testing.T) {
	var dummySource Source = "dummy"

	assert.Equal(t, "dummy", dummySource.String())
}
