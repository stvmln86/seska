package sqls

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	// success
	db := sqlx.MustConnect("sqlite3", ":memory:"+Params)
	assert.NotNil(t, db)

	// confirm - parameters were executed
	var okay bool
	err := db.Get(&okay, "pragma foreign_keys")
	assert.True(t, okay)
	assert.NoError(t, err)
}

func TestSchema(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(Schema)
	assert.NoError(t, err)
}
