package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	// success
	db := DB(t)
	assert.NotNil(t, db)
}
