package book

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seska/seska/items/note"
	"github.com/stvmln86/seska/seska/tools/test"
)

func mockBook(t *testing.T) *Book {
	db := test.MockDB(t)
	return New(db)
}

func TestNew(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	book := New(db)
	assert.Equal(t, db, book.DB)
}

func TestOpen(t *testing.T) {
	// setup
	dire := t.TempDir()
	path := filepath.Join(dire, "bolt.db")

	// success
	book, err := Open(path)
	assert.NotNil(t, book.DB)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	note, err := book.Create("name", "Body.")
	assert.Equal(t, "name", note.Name)
	assert.Equal(t, "Body.", note.Body)
	assert.NoError(t, err)
}

func TestFilter(t *testing.T) {
	// setup
	book := mockBook(t)
	ffun := func(note *note.Note) bool { return note.Name == "alpha" }

	// success
	notes, err := book.Filter(ffun)
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	note, err := book.Get("alpha")
	assert.Equal(t, "alpha", note.Name)
	assert.Equal(t, "Alpha note.", note.Body)
	assert.NoError(t, err)
}
