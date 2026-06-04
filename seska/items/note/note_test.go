package note

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stvmln86/seska/seska/tools/bolt"
	"github.com/stvmln86/seska/seska/tools/errs"
	"github.com/stvmln86/seska/seska/tools/neat"
	"github.com/stvmln86/seska/seska/tools/test"
)

func mockNote(t *testing.T) *Note {
	db := test.MockDB(t)
	note, err := Get(db, "alpha")
	require.NoError(t, err)
	return note
}

func TestCreate(t *testing.T) {
	// setup
	db := test.MockDB(t)
	strf := neat.Strf(time.Now())

	// success
	note, err := Create(db, "name", "Body.")
	assert.Equal(t, db, note.DB)
	assert.Equal(t, "name", note.Name)
	assert.Equal(t, "Body.", note.Body)
	assert.Equal(t, "UhslzEWGhOtQSnyIWtzNdIzy-XQp_4ChSIbQgE1iyGI", note.Hash)
	assert.Equal(t, strf, note.Init)
	assert.Equal(t, strf, note.Last)
	assert.NoError(t, err)

	// confirm - database
	pairs, err := bolt.Get(db, "name")
	assert.Equal(t, map[string]string{
		"body": "Body.",
		"hash": "UhslzEWGhOtQSnyIWtzNdIzy-XQp_4ChSIbQgE1iyGI",
		"init": strf,
		"last": strf,
	}, pairs)
	assert.NoError(t, err)

	// failure - note exists
	note, err = Create(db, "name", "Body.")
	assert.Nil(t, note)
	assert.ErrorIs(t, err, errs.Exists)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	note, err := Get(db, "alpha")
	assert.Equal(t, db, note.DB)
	assert.Equal(t, "alpha", note.Name)
	assert.Equal(t, "Alpha note.", note.Body)
	assert.Equal(t, "wUbNdwPXVxhUClSw1w_nL01Xud_xfmmtuMnZeoyiFrY", note.Hash)
	assert.Equal(t, "1970-01-01 00:00:00 UTC", note.Init)
	assert.Equal(t, "1970-01-01 01:00:00 UTC", note.Last)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	err := note.Delete()
	assert.NoError(t, err)

	// confirm - database
	okay, err := bolt.Exists(note.DB, "alpha")
	assert.False(t, okay)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	okay, err := note.Exists()
	assert.True(t, okay)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	note := mockNote(t)
	strf := neat.Strf(time.Now())

	// success
	err := note.Update("Body.")
	assert.Equal(t, "Body.", note.Body)
	assert.Equal(t, "UhslzEWGhOtQSnyIWtzNdIzy-XQp_4ChSIbQgE1iyGI", note.Hash)
	assert.Equal(t, "1970-01-01 00:00:00 UTC", note.Init)
	assert.Equal(t, strf, note.Last)
	assert.NoError(t, err)

	// confirm - database
	pairs, err := bolt.Get(note.DB, "alpha")
	assert.Equal(t, map[string]string{
		"body": "Body.",
		"hash": "UhslzEWGhOtQSnyIWtzNdIzy-XQp_4ChSIbQgE1iyGI",
		"init": "1970-01-01 00:00:00 UTC",
		"last": strf,
	}, pairs)
	assert.NoError(t, err)
}
