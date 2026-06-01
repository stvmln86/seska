package page

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seska/seska/tools/neat"
	"github.com/stvmln86/seska/seska/tools/test"
)

func assertPage(t *testing.T, page *Page, p_id, secs, note int64, body string) {
	assert.NotNil(t, page.Tx)
	assert.Equal(t, p_id, page.ID)
	assert.InDelta(t, time.Now().Unix()-secs, page.Init, 1.0)
	assert.Equal(t, note, page.Note)
	assert.Equal(t, body, page.Body)
	assert.Equal(t, neat.Hash(body), page.Hash)
}

func TestCreate(t *testing.T) {
	// setup
	_, tx := test.MockTx(t)

	// success
	page, err := Create(tx, 1, "body")
	assertPage(t, page, 4, 0, 1, "body")
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
	assertPage(t, page, 2, 5400, 1, "Alpha two.")
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
