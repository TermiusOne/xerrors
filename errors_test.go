package xerrors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseError_Code(t *testing.T) {
	var code int64 = 10

	someErr := New(code, "1", true)

	assert.Equal(t, someErr.Code(), code)
}

func TestXError_IsLogic(t *testing.T) {
	someErr := New(10, "1", true)

	assert.Equal(t, someErr.IsLogic(), true)
}

func TestBaseError_WithError(t *testing.T) {
	err := New(1, "some message 1", true).WithError(errors.New("some message 2"))

	assert.Equal(t, err.Error(), "code: 1, error: some message 1: some message 2")
}

func TestBaseError_WithDetails(t *testing.T) {
	err := New(2, "some message 1", false).WithDetails("some message %d", 2)

	assert.Equal(t, err.Error(), "code: 2, error: some message 1: some message 2")
}
