// Package test implements unit testing data and functions.
package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stvmln86/seska/seska/tools/sqls"
)

// mockData is additional database data for unit testing.
const mockData = `
	insert into Notes (init, name, hash) values
		(unixepoch()-3600, 'alpha', '8ed3f6ad685b959ead7022518e1af76cd816f8e8ec7ccdda1ed4018e8f2223f8'),
		(unixepoch(),      'bravo', 'f144a6907dc4284d1f9fe6a7d9b9ff53c02c1d07ba68f24d413d7ff7f757a782');
`

// MockDB returns an in-memory database populated with mock data.
func MockDB(t *testing.T) *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	db.MustExec(sqls.Pragma + sqls.Schema + mockData)

	t.Helper()
	t.Cleanup(func() { db.Close() })
	return db
}
