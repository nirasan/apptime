package apptime

import (
	"log"
	"time"
)

type Apptime struct {
	timeFunc func() time.Time
}

const TimeFormat = "2006-01-02 15:04:05"

func New() *Apptime {
	return &Apptime{
		timeFunc: time.Now,
	}
}

func (a *Apptime) Now() time.Time {
	return a.timeFunc()
}

func (a *Apptime) Set(t time.Time) {
	a.timeFunc = func() time.Time { return t }
}

func (a *Apptime) SetByString(s string) {
	t, e := time.Parse(TimeFormat, s)
	if e != nil {
		log.Panic("invalid string: " + s)
	}
	a.Set(t)
}

func (a *Apptime) Reset() {
	a.timeFunc = time.Now
}
