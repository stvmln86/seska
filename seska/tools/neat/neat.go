// Package neat implements data sanitisation and conversion functions.
package neat

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"

	"github.com/stvmln86/seska/seska/tools/errs"
)

// Body returns a whitespace-trimmed body string.
func Body(body string) string {
	return strings.TrimSpace(body)
}

// Hash returns a base64-encoded SHA256 hash of a string.
func Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

// Name returns a lowercase whitespace-trimmed name string.
func Name(name string) string {
	name = strings.TrimSpace(name)
	return strings.ToLower(name)
}

// Strf returns a datetime string from a Time object.
func Strf(tobj time.Time) string {
	return tobj.Format("2006-01-02 15:04:05 UTC")
}

// Time returns a Time object from a datetime string.
func Time(strf string) (time.Time, error) {
	tobj, err := time.Parse("2006-01-02 15:04:05 UTC", strf)
	if err != nil {
		return time.Time{}, errs.InvalidTime
	}

	return tobj, nil
}
