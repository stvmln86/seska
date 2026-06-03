// Package test implements unit testing data and functions.
package test

import (
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"
)

// MockData is a map of mock database data for unit testing.
var MockData = map[string]map[string]string{
	"alpha": {
		"body": "Alpha note.\n",
		"flag": "",
		"hash": "zZ/ACc3JXIMP0FffZsDDY9R29NrVNKISX9qTc9NX5wI=",
		"init": strconv.FormatInt(time.Now().Unix(), 10),
		"last": strconv.FormatInt(time.Now().Unix(), 10),
	},
}

// MockDB returns a temporary database populated with mock data.
func MockDB(t *testing.T) *bbolt.DB {
	dire := t.TempDir()
	path := filepath.Join(dire, "bolt.db")
	db, err := bbolt.Open(path, 0600, nil)
	require.NoError(t, err)
	require.NoError(t, db.Update(func(tx *bbolt.Tx) error {
		for name, pairs := range MockData {
			buck, err := tx.CreateBucket([]byte(name))
			require.NoError(t, err)

			for attr, valu := range pairs {
				err := buck.Put([]byte(attr), []byte(valu))
				require.NoError(t, err)
			}
		}

		return nil
	}))

	return db
}
