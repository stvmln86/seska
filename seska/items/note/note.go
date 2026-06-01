// Package note implements the Note type and methods.
package note

import (
	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/seska/seska/tools/neat"
)

// Note is a single recorded note.
type Note struct {
	Tx   *sqlx.Tx `db:"-"`
	ID   int64    `db:"id"`
	Init int64    `db:"init"`
	Name string   `db:"name"`
	Hash string   `db:"hash"`
}

// Create creates and returns a new Note.
func Create(tx *sqlx.Tx, name, body string) (*Note, error) {
	note := &Note{Tx: tx}
	code := "insert into Notes (name, hash) values (?, ?) returning *"
	name, hash := neat.Name(name)
	if err := tx.Get(note, code, name, hash); err != nil {
		return nil, err
	}

	code = "insert into Pages (note, body, hash) values (?, ?, ?)"
	body, hash = neat.Body(body)
	if _, err := tx.Exec(code, note.ID, body, hash); err != nil {
		return nil, err
	}

	return note, nil
}

// Get returns an existing Note by name.
func Get(tx *sqlx.Tx, name string) (*Note, error) {
	note := &Note{Tx: tx}
	code := "select * from Notes where name=? and hash=? limit 1"
	name, hash := neat.Name(name)
	if err := tx.Get(note, code, name, hash); err != nil {
		return nil, err
	}

	return note, nil
}

// Exists returns true if the Note exists.
func (n *Note) Exists() (bool, error) {
	var okay bool
	code := "select exists (select 1 from Notes where id=?)"
	if err := n.Tx.Get(&okay, code, n.ID); err != nil {
		return false, err
	}

	return okay, nil
}
