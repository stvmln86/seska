package page

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seska/seska/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	_, tx := test.MockTx(t)

	// success
	page, err := Create(tx, 1, "body")
	assert.Equal(t, tx, page.Tx)
	assert.Equal(t, int64(4), page.ID)
	test.AssertInit(t, page.Init, 0)
	assert.Equal(t, int64(1), page.Note)
	assert.Equal(t, "body", page.Body)
	assert.Equal(t, "Iw2DWNyOiJC0xY3utikS7i8gNXrpKlzIYbmOaP4xrLU", page.Hash)
	assert.NoError(t, err)

	// confirm - transaction
	err = tx.Commit()
	assert.NoError(t, err)
}

func TestLatest(t *testing.T) {
	// setup
	_, tx := test.MockTx(t)

	// success
	page, err := Latest(tx, 1)
	assert.Equal(t, tx, page.Tx)
	assert.Equal(t, int64(2), page.ID)
	test.AssertInit(t, page.Init, 90*time.Minute)
	assert.Equal(t, int64(1), page.Note)
	assert.Equal(t, "Alpha two.", page.Body)
	assert.Equal(t, "eF1U-1JjWcek5mfcB9IsZCXC8SHws7bZrPWJ7YeVSiA", page.Hash)
	assert.NoError(t, err)

	// confirm - transaction
	err = tx.Commit()
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	_, tx := test.MockTx(t)
	page, _ := Latest(tx, 1)

	// success - true
	okay, err := page.Exists()
	assert.True(t, okay)
	assert.NoError(t, err)

	// setup
	page = &Page{Tx: tx, ID: -1}

	// success - false
	okay, err = page.Exists()
	assert.False(t, okay)
	assert.NoError(t, err)

	// confirm - transaction
	err = tx.Commit()
	assert.NoError(t, err)
}
