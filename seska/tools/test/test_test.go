package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockDB(t *testing.T) {
	// success
	db := MockDB(t)
	assert.NotNil(t, db)
}

func TestMockTx(t *testing.T) {
	// success
	tx := MockTx(t)
	assert.NotNil(t, tx)
}
