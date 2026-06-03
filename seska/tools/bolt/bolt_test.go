package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seska/seska/tools/test"
	"go.etcd.io/bbolt"
)

func TestDelete(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	err := Delete(db, "alpha")
	assert.NoError(t, err)

	// confirm - database
	assert.NoError(t, db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("alpha"))
		assert.Nil(t, buck)
		return nil
	}))
}

func TestExists(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success - true
	okay, err := Exists(db, "alpha")
	assert.True(t, okay)
	assert.NoError(t, err)

	// success - false
	okay, err = Exists(db, "nope")
	assert.False(t, okay)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	pairs, err := Get(db, "alpha")
	assert.Equal(t, test.MockData["alpha"], pairs)
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	names, err := List(db)
	assert.Equal(t, []string{"alpha"}, names)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	err := Set(db, "name", map[string]string{"attr": "valu"})
	assert.NoError(t, err)

	// confirm - database
	assert.NoError(t, db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("name"))
		assert.NoError(t, err)

		valu := buck.Get([]byte("attr"))
		assert.Equal(t, []byte("valu"), valu)
		return nil
	}))
}
