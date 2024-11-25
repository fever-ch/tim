package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimNoTZNoShift(t *testing.T) {
	ts, err := parseTim("1991-08-06T16:56:20")

	assert.Nil(t, err)
	assert.Equal(t, int64(681490580000), ts.UnixMilli())
}

func TestTimWithTZNoShift(t *testing.T) {
	ts, err := parseTim("1991-08-06T16:56:20+02:00")

	assert.Nil(t, err)
	assert.Equal(t, int64(681490580000), ts.UnixMilli())
}

func TestTimNoTZWithShift(t *testing.T) {
	ts, err := parseTim("1991-06-06T16:56:20+02:00+2d")

	assert.Nil(t, err)
	assert.Equal(t, int64(676392980000), ts.UnixMilli())
}

func TestTimWithTZWithShift(t *testing.T) {
	ts, err := parseTim("1991-08-06T20:56:20+02:00-4h")

	assert.Nil(t, err)
	assert.Equal(t, int64(681490580000), ts.UnixMilli())
}

func TestTimNowNoShift(t *testing.T) {
	ts, e := parseTim("now")

	assert.Nil(t, e)
	assert.InDelta(t, time.Now().UnixMilli(), ts.UnixMilli(), 1000)
}

func TestTimNowWithShift(t *testing.T) {
	ts, err := parseTim("now-2w")

	assert.Nil(t, err)
	assert.InDelta(t, time.Now().UnixMilli()-2*7*24*3600*1000, ts.UnixMilli(), 1000)
}

func TestBadTime(t *testing.T) {
	_, err := parseTim("2024-02-31T12:34:67-4d")

	assert.NotNil(t, err)
}

func TestBadDelay(t *testing.T) {
	_, err := parseTim("1991-06-06T16:56:20+02:00+2x")

	assert.NotNil(t, err)
}

func TestWithLocation(t *testing.T) {
	ts, err := parseTim("1991-08-06T12:56:20@Europe/Zurich+4h")

	assert.Nil(t, err)
	assert.Equal(t, int64(681490580000), ts.UnixMilli())
}
