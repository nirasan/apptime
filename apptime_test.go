package apptime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestApptime_Now(t *testing.T) {
	a := New()

	assert.Equal(t, time.Now().Unix(), a.Now().Unix())

	now, _ := time.Parse("2006-01-02 15:04:05", "2001-02-03 04:05:06")
	a.Set(now)

	assert.NotEqual(t, time.Now().Unix(), a.Now().Unix())
	assert.Equal(t, a.Now(), now)

	a.Reset()

	assert.Equal(t, time.Now().Unix(), a.Now().Unix())

	a.SetByString("2001-02-03 04:05:06")

	assert.NotEqual(t, time.Now().Unix(), a.Now().Unix())
	assert.Equal(t, a.Now(), now)
}
