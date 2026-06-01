// Package neat implements value sanitisation and conversion functions.
package neat

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

// Body returns a whitespace-trimmed body string and its hash.
func Body(body string) (string, string) {
	body = strings.TrimSpace(body)
	return body, Hash(body)
}

// Hash returns a base64-encoded SHA256 hash of a string.
func Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

// Like returns a SQLite "LIKE" pattern with escaped wildcard characters.
func Like(text string) string {
	esca := strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`)
	return "%" + esca.Replace(text) + "%"
}

// Name returns a lowercase whitespace-trimmed name string and its hash.
func Name(name string) (string, string) {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name, Hash(name)
}
