package errs

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTry(t *testing.T) {
	// setup
	var exit bool
	var buff = bytes.NewBuffer(nil)
	exitFunc = func(int) { exit = true }

	// success
	Try(buff, errors.New("message"))
	assert.Equal(t, "Error: message.\n", buff.String())
	assert.True(t, exit)
}
