// Package page implements the Page type and methods.
package page

import (
	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/seska/seska/tools/neat"
)

// Page is a single recorded version of a Note.
type Page struct {
	Tx   *sqlx.Tx `db:"-"`
	ID   int64    `db:"id"`
	Init int64    `db:"init"`
	Note int64    `db:"note"`
	Body string   `db:"body"`
	Hash string   `db:"hash"`
}

// Create creates and returns a new Page.
func Create(tx *sqlx.Tx, note int64, body string) (*Page, error) {
	page := &Page{Tx: tx}
	code := "insert into Pages (note, body, hash) values (?, ?, ?) returning *"
	body, hash := neat.Body(body)
	if err := tx.Get(page, code, note, body, hash); err != nil {
		return nil, err
	}

	return page, nil
}

// Latest returns a Note's latest Page.
func Latest(tx *sqlx.Tx, note int64) (*Page, error) {
	page := &Page{Tx: tx}
	code := "select * from Pages where note=? order by id desc limit 1"
	if err := tx.Get(page, code, note); err != nil {
		return nil, err
	}

	return page, nil
}

// Exists returns true if the Page exists.
func (p *Page) Exists() (bool, error) {
	var okay bool
	code := "select exists (select 1 from Pages where id=?)"
	if err := p.Tx.Get(&okay, code, p.ID); err != nil {
		return false, err
	}

	return okay, nil
}
