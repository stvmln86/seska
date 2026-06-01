package note

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seska/seska/tools/neat"
	"github.com/stvmln86/seska/seska/tools/test"
)

func assertNote(t *testing.T, note *Note, n_id, secs int64, name string) {
	assert.NotNil(t, note.Tx)
	assert.Equal(t, n_id, note.ID)
	assert.InDelta(t, time.Now().Unix()-secs, note.Init, 1.0)
	assert.Equal(t, name, note.Name)
	assert.Equal(t, neat.Hash(name), note.Hash)
}

func TestCreate(t *testing.T) {
	// setup
	_, tx := test.MockTx(t)

	// success
	note, err := Create(tx, "name", "body")
	assertNote(t, note, 3, 0, "name")
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	_, tx := test.MockTx(t)

	// success
	note, err := Get(tx, "alpha")
	assertNote(t, note, 1, 7200, "alpha")
	assert.NoError(t, err)

	// failure - note does not exist
	note, err = Get(tx, "nope")
	assert.Nil(t, note)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}

func TestMatch(t *testing.T) {
	// setup
	_, tx := test.MockTx(t)

	// success
	notes, err := Match(tx, "ALPH")
	assert.Len(t, notes, 1)
	assertNote(t, notes[0], 1, 7200, "alpha")
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
}
