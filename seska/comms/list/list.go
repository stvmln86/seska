// Package list implements the ListComm type and methods.
package list

import (
	"fmt"
	"io"

	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/seska/seska/items/note"
)

// List is a command that lists existing Notes.
type List struct {
	Text string `arg:"" optional:"" help:""`
}

// Run executes the ListComm's command.
func (c *List) Run(w io.Writer, tx *sqlx.Tx) error {
	notes, err := note.Match(tx, c.Text)
	if err != nil {
		return err
	}

	for _, n := range notes {
		fmt.Fprintf(w, "%s\n", n.Name)
	}

	return nil
}
