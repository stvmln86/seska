package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAssertInit(t *testing.T) {
	// setup
	tnow := time.Now().Unix()

	// success
	AssertInit(t, tnow, 0)
}

func TestMockDB(t *testing.T) {
	// success
	db := MockDB(t)
	assert.NotNil(t, db)
}

func TestMockTx(t *testing.T) {
	// success
	db, tx := MockTx(t)
	assert.NotNil(t, db)
	assert.NotNil(t, tx)
}
