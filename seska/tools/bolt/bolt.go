// Package bolt implements bolt database handling functions.
package bolt

import (
	"go.etcd.io/bbolt"
)

// Delete deletes an existing database bucket
func Delete(db *bbolt.DB, buck string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(buck))
	})
}

// Exists returns true if a database bucket exists.
func Exists(db *bbolt.DB, name string) (bool, error) {
	var okay bool
	return okay, db.View(func(tx *bbolt.Tx) error {
		okay = tx.Bucket([]byte(name)) != nil
		return nil
	})
}

// Get returns an existing database bucket as a map.
func Get(db *bbolt.DB, name string) (map[string]string, error) {
	var pairs map[string]string
	return pairs, db.View(func(tx *bbolt.Tx) error {
		if buck := tx.Bucket([]byte(name)); buck != nil {
			pairs = make(map[string]string)
			return buck.ForEach(func(attr, valu []byte) error {
				pairs[string(attr)] = string(valu)
				return nil
			})
		}

		return nil
	})
}

// List returns the names of all existing buckets in a database.
func List(db *bbolt.DB) ([]string, error) {
	var names []string
	return names, db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			names = append(names, string(name))
			return nil
		})
	})
}

// Set overwrites a new or existing database bucket from a map.
func Set(db *bbolt.DB, name string, pairs map[string]string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		buck, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}

		for attr, valu := range pairs {
			if err := buck.Put([]byte(attr), []byte(valu)); err != nil {
				return err
			}
		}

		return nil
	})
}
