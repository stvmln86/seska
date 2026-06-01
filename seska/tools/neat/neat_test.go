package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body, hash := Body("\tBody.\n")
	assert.Equal(t, "Body.", body)
	assert.Equal(t, "UhslzEWGhOtQSnyIWtzNdIzy-XQp_4ChSIbQgE1iyGI", hash)
}

func TestHash(t *testing.T) {
	// success
	hash := Hash("text")
	assert.Equal(t, "mC2ePrmW9VnmM_TRlN7zdh2Qn1o7ZH0ahR_q1nwyydE", hash)
}

func TestLike(t *testing.T) {
	// success
	like := Like(`a%b_c\d`)
	assert.Equal(t, `%a\%b\_c\\d%`, like)
}

func TestName(t *testing.T) {
	// success
	name, hash := Name("\tNAME\n")
	assert.Equal(t, "name", name)
	assert.Equal(t, "gqNTf_Dbzn7sNdae3DoYnubxfYLzU6VT-aqWywvjzok", hash)
}
