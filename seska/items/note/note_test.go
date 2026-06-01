package note

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seska/seska/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	note, err := Create(db, "name")
	assert.Equal(t, db, note.DB)
	assert.Equal(t, int64(3), note.ID)
	assert.Equal(t, time.Now().Unix(), note.Init)
	assert.Equal(t, "name", note.Name)
	assert.Equal(t, "gqNTf_Dbzn7sNdae3DoYnubxfYLzU6VT-aqWywvjzok", note.Hash)
	assert.NoError(t, err)

	// confirm - database
	var okay bool
	err = db.Get(&okay, "select exists (select 1 from Notes where name='name')")
	assert.True(t, okay)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	note, err := Get(db, "alpha")
	assert.Equal(t, db, note.DB)
	assert.Equal(t, int64(1), note.ID)
	assert.Equal(t, time.Now().Unix()-7200, note.Init)
	assert.Equal(t, "alpha", note.Name)
	assert.Equal(t, "jtP2rWhblZ6tcCJRjhr3bNgW-OjsfM3aHtQBjo8iI_g", note.Hash)
	assert.NoError(t, err)

	// failure - note does not exist
	note, err = Get(db, "nope")
	assert.Nil(t, note)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}

func TestExists(t *testing.T) {
	// setup
	db := test.MockDB(t)
	note, _ := Get(db, "alpha")

	// success - true
	okay, err := note.Exists()
	assert.True(t, okay)
	assert.NoError(t, err)

	// setup
	note = &Note{DB: db, ID: -1}

	// success - false
	okay, err = note.Exists()
	assert.False(t, okay)
	assert.NoError(t, err)
}
