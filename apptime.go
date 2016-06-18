package apptime

import (
    "log"
    "time"
)

var timeNowFunc = time.Now

var TimeFormat = "2006-01-02 15:04:05"

func Now() time.Time {
    return timeNowFunc()
}

func Set(t time.Time) {
    timeNowFunc = func() time.Time { return t }
}

func SetByString(s string) {
    t, e := time.Parse(TimeFormat, s)
    if e != nil {
        log.Panic("invalid string: " + s)
    }
    Set(t)
}

func Reset() {
    timeNowFunc = time.Now
}

