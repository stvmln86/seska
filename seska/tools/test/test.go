// Package test implements unit testing data and functions.
package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"github.com/stvmln86/seska/seska/tools/sqls"
)

// mockData is additional database data for unit testing.
const mockData = `
	insert into Notes (init, name, hash) values
		(unixepoch()-7200, 'alpha',  'jtP2rWhblZ6tcCJRjhr3bNgW-OjsfM3aHtQBjo8iI_g'),
		(unixepoch()-3600, 'bravo',  '8USmkH3EKE0fn-an2bn_U8AsHQe6aPJNQT1_9_dXp4I');

	insert into Pages (init, note, body, hash) values
		(unixepoch()-7200, 1, 'Alpha one.', '2uX0Ji_QreTPsV3-XPXWwJgR7eA_eC8NFR_kF5z4klU'),
		(unixepoch()-5400, 1, 'Alpha two.', 'eF1U-1JjWcek5mfcB9IsZCXC8SHws7bZrPWJ7YeVSiA'),
		(unixepoch()-3600, 2, 'Bravo one.', 'fRI-7CujV00Kae22wjPk0E2vUbWmCk_53skTuPzPqVQ');
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

// MockTx returns a mocked database and transaction.
func MockTx(t *testing.T) (*sqlx.DB, *sqlx.Tx) {
	db := MockDB(t)
	tx, err := db.Beginx()
	require.NoError(t, err)
	return db, tx
}
