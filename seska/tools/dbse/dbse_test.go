package dbse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	// success
	db, err := Open(":memory:")
	assert.NotNil(t, db)
	assert.NoError(t, err)

	// confirm - pragma was executed
	var okay bool
	err = db.Get(&okay, "pragma foreign_keys")
	assert.True(t, okay)
	assert.NoError(t, err)

	// confirm - schema was executed
	var size int
	err = db.Get(&size, "select count(*) from SQLITE_SCHEMA")
	assert.NotZero(t, size)
	assert.NoError(t, err)
}
