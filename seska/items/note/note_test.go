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
	_, tx := test.MockTx(t)

	// success
	note, err := Create(tx, "name", "body")
	assert.Equal(t, tx, note.Tx)
	assert.Equal(t, int64(3), note.ID)
	test.AssertInit(t, note.Init, 0)
	assert.Equal(t, "name", note.Name)
	assert.Equal(t, "gqNTf_Dbzn7sNdae3DoYnubxfYLzU6VT-aqWywvjzok", note.Hash)
	assert.NoError(t, err)

	// confirm - transaction
	err = tx.Commit()
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	_, tx := test.MockTx(t)

	// success
	note, err := Get(tx, "alpha")
	assert.Equal(t, tx, note.Tx)
	assert.Equal(t, int64(1), note.ID)
	test.AssertInit(t, note.Init, 2*time.Hour)
	assert.Equal(t, "alpha", note.Name)
	assert.Equal(t, "jtP2rWhblZ6tcCJRjhr3bNgW-OjsfM3aHtQBjo8iI_g", note.Hash)
	assert.NoError(t, err)

	// failure - note does not exist
	note, err = Get(tx, "nope")
	assert.Nil(t, note)
	assert.ErrorIs(t, err, sql.ErrNoRows)

	// confirm - transaction
	err = tx.Commit()
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	_, tx := test.MockTx(t)
	note, _ := Get(tx, "alpha")

	// success - true
	okay, err := note.Exists()
	assert.True(t, okay)
	assert.NoError(t, err)

	// setup
	note = &Note{Tx: tx, ID: -1}

	// success - false
	okay, err = note.Exists()
	assert.False(t, okay)
	assert.NoError(t, err)

	// confirm - transaction
	err = tx.Commit()
	assert.NoError(t, err)
}
