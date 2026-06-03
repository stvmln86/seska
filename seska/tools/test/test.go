// Package test implements unit testing data and functions.
package test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"
)

// MockData is a map of mock database data for unit testing.
var MockData = map[string]map[string]string{
	"alpha": {
		"body": "Alpha note.",
		"hash": "wUbNdwPXVxhUClSw1w_nL01Xud_xfmmtuMnZeoyiFrY",
		"init": "1970-01-01 00:00:00 UTC",
		"last": "1970-01-01 01:00:00 UTC",
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
