// Package neat implements value sanitisation and conversion functions.
package neat

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// Body returns a whitespace-trimmed newline-ended body string and its hash.
func Body(body string) (string, string) {
	body = strings.TrimSpace(body) + "\n"
	return body, Hash(body)
}

// Hash returns a SHA256 Hash of a string.
func Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%x", hash)
}

// Name returns a lowercase whitespace-trimmed name string and its hash.
func Name(name string) (string, string) {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name, Hash(name)
}
