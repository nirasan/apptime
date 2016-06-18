package apptime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	assert.Equal(t, time.Now().Unix(), Now().Unix())

	now, _ := time.Parse("2006-01-02 15:04:05", "2001-02-03 04:05:06")
	Set(now)

	assert.NotEqual(t, time.Now().Unix(), Now().Unix())
	assert.Equal(t, Now(), now)

	Reset()

	assert.Equal(t, time.Now().Unix(), Now().Unix())

	SetByString("2001-02-03 04:05:06")

	assert.NotEqual(t, time.Now().Unix(), Now().Unix())
	assert.Equal(t, Now(), now)
}

