// Package dbse implements database handling functions.
package dbse

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stvmln86/seska/seska/tools/sqls"
)

// Open returns a new database connection with applied parameters and schema.
func Open(path string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", path+sqls.Params)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	if _, err := db.Exec(sqls.Schema); err != nil {
		return nil, err
	}

	return db, nil
}
