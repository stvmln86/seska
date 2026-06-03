// Package note implements the Note type and methods.
package note

import (
	"time"

	"github.com/stvmln86/seska/seska/tools/bolt"
	"github.com/stvmln86/seska/seska/tools/errs"
	"github.com/stvmln86/seska/seska/tools/neat"
	"go.etcd.io/bbolt"
)

// Note is a single recorded note in a Book.
type Note struct {
	DB   *bbolt.DB
	Name string
	Body string
	Hash string
	Init string
	Last string
}

// Create creates and returns a new Note.
func Create(db *bbolt.DB, name, body string) (*Note, error) {
	name = neat.Name(name)
	okay, err := bolt.Exists(db, name)
	switch {
	case okay:
		return nil, errs.Exists
	case err != nil:
		return nil, errs.Database
	}

	body = neat.Body(body)
	pairs := map[string]string{
		"body": body,
		"hash": neat.Hash(body),
		"init": neat.Strf(time.Now()),
		"last": neat.Strf(time.Now()),
	}

	if err := bolt.Set(db, name, pairs); err != nil {
		return nil, errs.Database
	}

	return Get(db, name)
}

// Get returns an existing Note.
func Get(db *bbolt.DB, name string) (*Note, error) {
	name = neat.Name(name)
	pairs, err := bolt.Get(db, name)
	switch {
	case pairs == nil:
		return nil, nil
	case err != nil:
		return nil, err
	}

	return &Note{
		db, name,
		pairs["body"], pairs["hash"],
		pairs["init"], pairs["last"],
	}, nil
}

// Delete deletes the Note.
func (n *Note) Delete() error {
	return bolt.Delete(n.DB, n.Name)
}

// Exists returns true if the Note exists.
func (n *Note) Exists() (bool, error) {
	return bolt.Exists(n.DB, n.Name)
}

// Update overwrites the Note's body.
func (n *Note) Update(body string) error {
	body = neat.Body(body)
	pairs := map[string]string{
		"body": body,
		"hash": neat.Hash(body),
		"last": neat.Strf(time.Now()),
	}

	if err := bolt.Set(n.DB, n.Name, pairs); err != nil {
		return errs.Database
	}

	n.Body = pairs["body"]
	n.Hash = pairs["hash"]
	n.Last = pairs["last"]
	return nil
}
