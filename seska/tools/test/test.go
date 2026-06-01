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
		(unixepoch()-7200, 'alpha',   'jtP2rWhblZ6tcCJRjhr3bNgW-OjsfM3aHtQBjo8iI_g'),
		(unixepoch()-3600, 'bravo',   '8USmkH3EKE0fn-an2bn_U8AsHQe6aPJNQT1_9_dXp4I'),
		(unixepoch(),      'charlie', 'ud2WDBdTRZp4EV08uEWlfZJLaHfoBbCL0BCGzN80Qzw');
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
