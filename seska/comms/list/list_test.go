package list

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seska/seska/tools/test"
)

func TestRun(t *testing.T) {
	// setup
	w := bytes.NewBuffer(nil)
	_, tx := test.MockTx(t)
	list := &List{Text: "ALPH"}

	// success
	err := list.Run(w, tx)
	assert.Equal(t, "alpha\n", w.String())
	assert.NoError(t, err)
}
