package clui

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	// setup
	os.Clearenv()
	t.Setenv("HOME", "/home/user")

	// success - HOME
	path := Path()
	assert.Equal(t, "/home/user/.seska", path)

	// setup
	t.Setenv("XDG_CONFIG_HOME", "/home/user/.config")

	// success - XDG_CONFIG_HOME
	path = Path()
	assert.Equal(t, "/home/user/.config/seska.db", path)

	// setup
	t.Setenv("SESKA_DB", "./seska.db")

	// success - SESKA_DB
	path = Path()
	assert.Equal(t, "./seska.db", path)
}
