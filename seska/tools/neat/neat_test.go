package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seska/seska/tools/errs"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.", body)
}

func TestHash(t *testing.T) {
	// success
	hash := Hash("text")
	assert.Equal(t, "mC2ePrmW9VnmM_TRlN7zdh2Qn1o7ZH0ahR_q1nwyydE", hash)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME\n")
	assert.Equal(t, "name", name)
}

func TestStrf(t *testing.T) {
	// setup
	tobj := time.Unix(0, 0).UTC()

	// success
	strf := Strf(tobj)
	assert.Equal(t, "1970-01-01 00:00:00 UTC", strf)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(0, 0).UTC()

	// success
	tobj, err := Time("1970-01-01 00:00:00 UTC")
	assert.Equal(t, want, tobj)
	assert.NoError(t, err)

	// failure - invalid time
	tobj, err = Time("")
	assert.Equal(t, time.Time{}, tobj)
	assert.ErrorIs(t, err, errs.InvalidTime)
}
