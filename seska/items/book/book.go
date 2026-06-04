// Package book implements the Book type and methods.
package book

import (
	"github.com/stvmln86/seska/seska/items/note"
	"github.com/stvmln86/seska/seska/tools/bolt"
	"go.etcd.io/bbolt"
)

// Book is a database of recorded Notes.
type Book struct {
	DB *bbolt.DB
}

// New returns a new Book.
func New(db *bbolt.DB) *Book {
	return &Book{db}
}

// Open returns a new Book from a database path.
func Open(path string) (*Book, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	return New(db), nil
}

// Create creates and returns a new Note in the Book.
func (b *Book) Create(name, body string) (*note.Note, error) {
	return note.Create(b.DB, name, body)
}

// Filter returns all existing Notes in the Book matching a filter function.
func (b *Book) Filter(ffun func(*note.Note) bool) ([]*note.Note, error) {
	var notes []*note.Note
	names, err := bolt.List(b.DB)
	if err != nil {
		return nil, err
	}

	for _, name := range names {
		note, err := note.Get(b.DB, name)
		switch {
		case err != nil:
			return nil, err
		case ffun(note):
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// Get returns an existing Note from the Book.
func (b *Book) Get(name string) (*note.Note, error) {
	return note.Get(b.DB, name)
}
