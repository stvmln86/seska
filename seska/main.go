// Package main implements the Seska command-line program.
package main

import (
	"container/list"
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/kong"
	"github.com/stvmln86/seska/seska/tools/dbse"
	"github.com/stvmln86/seska/seska/tools/test"
)

type Seska struct {
	List list.List `cmd:"" help:"list existing notes"`
}

func try(err error) {
	if err != nil {
		fmt.Printf("Error: %s.\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	db, err := dbse.Open(":memory:") // temp
	db.MustExec(test.MockData)       // temp
	try(err)

	core := new(Seska)
	ktxt := kong.Parse(core,
		kong.Name("seska"),
		kong.Description("Stephen's Eternal Scrap Keeper Application."),
		kong.ShortUsageOnError(),
	)

	tx, err := db.Beginx()
	try(err)

	ktxt.Bind(tx)
	ktxt.BindTo(os.Stdout, (*io.Writer)(nil))
	try(ktxt.Run())
	try(tx.Commit())
}
