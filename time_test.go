package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	ts, e := parseTime("now")

	assert.Nil(t, e)
	assert.InDelta(t, time.Now().UnixMilli(), ts.UnixMilli(), 1000)
}

func TestToday(t *testing.T) {
	now := time.Now()

	if now.Hour() == 23 && now.Minute() == 59 && now.Second() == 59 {
		time.Sleep(time.Second * 1)
		now = time.Now()
	}

	ts, e := parseTime("today")

	assert.Nil(t, e)

	assert.Equal(t, now.Year(), ts.Year())
	assert.Equal(t, now.Month(), ts.Month())
	assert.Equal(t, now.Day(), ts.Day())

	assert.Equal(t, 0, ts.Hour())
	assert.Equal(t, 0, ts.Minute())
	assert.Equal(t, 0, ts.Second())
	assert.Equal(t, 0, ts.Nanosecond())
}

func TestNoTZ(t *testing.T) {
	tsLocal, eLocal := parseTime("1991-08-06T16:56:20")

	tsGMT, eGMT := parseTime("1991-08-06T16:56:20Z")

	assert.Nil(t, eLocal)
	assert.Nil(t, eGMT)

	_, offset := tsLocal.Zone()
	assert.Equal(t, int64(offset*1000), tsGMT.Sub(tsLocal).Milliseconds())
}

func TestGMT(t *testing.T) {
	ts, e := parseTime("1991-08-06T14:56:20+00:00")

	assert.Nil(t, e)
	assert.Equal(t, int64(681490580000), ts.UnixMilli())
}

func TestGMT2(t *testing.T) {
	ts, e := parseTime("1991-08-06T16:56:20+02:00")

	assert.Nil(t, e)
	assert.Equal(t, int64(681490580000), ts.UnixMilli())
}
