// Package main implements the Seska command-line program.
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/kong"
	"github.com/stvmln86/seska/seska/comms/list"
	"github.com/stvmln86/seska/seska/tools/clui"
	"github.com/stvmln86/seska/seska/tools/dbse"
)

// Seska is the top-level application controller.
type Seska struct {
	List list.List `cmd:"" help:"list existing notes"`
}

// try fatally prints a non-nil error.
func try(err error) {
	if err != nil {
		fmt.Printf("Error: %s.\n", err.Error())
		os.Exit(1)
	}
}

// main runs the main Seska program.
func main() {
	core := new(Seska)
	ktxt := kong.Parse(core,
		kong.Name("seska"),
		kong.Description("Stephen's Eternal Scrap Keeper Application."),
		kong.ShortUsageOnError(),
	)

	path := clui.Path()
	db, err := dbse.Open(path)
	try(err)

	tx, err := db.Beginx()
	try(err)

	ktxt.Bind(tx)
	ktxt.BindTo(os.Stdout, (*io.Writer)(nil))
	try(ktxt.Run())
	try(tx.Commit())
	try(db.Close())
}
