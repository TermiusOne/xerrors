package xerrors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseError_Code(t *testing.T) {
	var code = 10

	someErr := New(code, "1")

	assert.Equal(t, someErr.Code(), code)
}

func TestBaseError_Wrap(t *testing.T) {
	err := New(1, "some message 1").Wrap("some message 2")

	assert.Equal(t, err.Error(), "some message 1: some message 2")
}

func TestBaseError_Wrapf(t *testing.T) {
	err := New(1, "some message 1").Wrapf("some message %v", 2)

	assert.Equal(t, err.Error(), "some message 1: some message 2")
}
