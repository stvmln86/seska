// Package test implements unit testing data and functions.
package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stvmln86/seska/seska/tools/sqls"
)

// mockData is additional database data for unit testing.
const mockData = ``

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
