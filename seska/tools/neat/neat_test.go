package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body, hash := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
	assert.Equal(t, "44261ce242e1b99d52c7d2a4cb6dbcb5a4ab507bed9b9b303062a969fafe1d1e", hash)
}

func TestHash(t *testing.T) {
	// success
	hash := Hash("text")
	assert.Equal(t, "982d9e3eb996f559e633f4d194def3761d909f5a3b647d1a851fead67c32c9d1", hash)
}

func TestName(t *testing.T) {
	// success
	name, hash := Name("\tNAME\n")
	assert.Equal(t, "name", name)
	assert.Equal(t, "82a3537ff0dbce7eec35d69edc3a189ee6f17d82f353a553f9aa96cb0be3ce89", hash)
}
