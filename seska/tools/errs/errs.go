// Package errs implements system error definitions and functions.
package errs

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// exitFunc is the function used to exit the program.
var exitFunc func(int) = os.Exit

// Database is an error for failed database operations.
var Database = errors.New("database operation failed")

// EmptyBody is an error for empty body strings.
var EmptyBody = errors.New("body is empty")

// EmptyName is an error for empty name strings.
var EmptyName = errors.New("name is empty")

// InvalidTime is an error for invalid time strings.
var InvalidTime = errors.New("invalid time format")

// Try prints a non-nil error and exits.
func Try(w io.Writer, err error) {
	if err != nil {
		fmt.Fprintf(w, "Error: %s.\n", err.Error())
		exitFunc(1)
	}
}
