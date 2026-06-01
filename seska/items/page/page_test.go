package page

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seska/seska/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	page, err := Create(db, 1, "body")
	assert.Equal(t, db, page.DB)
	assert.Equal(t, int64(4), page.ID)
	assert.Equal(t, time.Now().Unix(), page.Init)
	assert.Equal(t, int64(1), page.Note)
	assert.Equal(t, "body", page.Body)
	assert.Equal(t, "Iw2DWNyOiJC0xY3utikS7i8gNXrpKlzIYbmOaP4xrLU", page.Hash)
	assert.NoError(t, err)

	// confirm - database
	var okay bool
	err = db.Get(&okay, "select exists (select 1 from Pages where body='body')")
	assert.True(t, okay)
	assert.NoError(t, err)
}

func TestLatest(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	page, err := Latest(db, 1)
	assert.Equal(t, db, page.DB)
	assert.Equal(t, int64(2), page.ID)
	assert.Equal(t, time.Now().Unix()-5400, page.Init)
	assert.Equal(t, int64(1), page.Note)
	assert.Equal(t, "Alpha two.", page.Body)
	assert.Equal(t, "eF1U-1JjWcek5mfcB9IsZCXC8SHws7bZrPWJ7YeVSiA", page.Hash)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	db := test.MockDB(t)
	page, _ := Latest(db, 1)

	// success - true
	okay, err := page.Exists()
	assert.True(t, okay)
	assert.NoError(t, err)

	// setup
	page = &Page{DB: db, ID: -1}

	// success - false
	okay, err = page.Exists()
	assert.False(t, okay)
	assert.NoError(t, err)
}
