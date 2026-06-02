// Package clui implements command-line user interface functions.
package clui

import (
	"os"
	"path/filepath"
)

// Path returns the system database path from SESKA_DB, XDG_CONFIG or HOME.
func Path() string {
	if evar, okay := os.LookupEnv("SESKA_DB"); okay {
		return evar
	}

	if evar, okay := os.LookupEnv("XDG_CONFIG_HOME"); okay {
		return filepath.Join(evar, "seska.db")
	}

	evar := os.Getenv("HOME")
	return filepath.Join(evar, ".seska")
}
