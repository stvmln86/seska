// Package main implements the Seska command-line program.
package main

import (
	"io"
	"os"

	"github.com/alecthomas/kong"
	"github.com/stvmln86/seska/seska/comms/list"
	"github.com/stvmln86/seska/seska/tools/dbse"
	"github.com/stvmln86/seska/seska/tools/test"
)

type Core struct {
	List list.List `cmd:"" help:"list existing notes"`
}

func main() {
	ktxt := kong.Parse(
		new(Core),
		kong.Name("seska"),
		kong.Description("Stephen's Eternal Scrap Keeper Application."),
		kong.ShortUsageOnError(),
	)

	db, err := dbse.Open(":memory:")
	ktxt.FatalIfErrorf(err)
	defer db.Close()

	_, err = db.Exec(test.MockData)
	ktxt.FatalIfErrorf(err)

	tx, err := db.Beginx()
	ktxt.FatalIfErrorf(err)
	defer tx.Rollback()

	ktxt.BindTo(os.Stdout, (*io.Writer)(nil))
	err = ktxt.Run(tx)
	ktxt.FatalIfErrorf(err)

	err = tx.Commit()
	ktxt.FatalIfErrorf(err)
}
