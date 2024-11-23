package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEmpty(t *testing.T) {
	_, e := time.ParseDuration("")
	if e == nil {
		t.Fail()
	}

}

func TestMultiplication(t *testing.T) {
	d, e := parseDuration("7*24h")

	assert.Nil(t, e)

	assert.Equal(t, 7*24*time.Hour, d)
}

func TestMultiplication2(t *testing.T) {
	d, e := parseDuration("7*24h3*2h")

	assert.Nil(t, e)
	assert.Equal(t, 7*24*time.Hour+3*2*time.Hour, d)
}

func TestDay(t *testing.T) {
	d, e := parseDuration("1d")
	if e != nil || d.Seconds() != 24*3600 {
		t.Fail()
	}
	assert.Nil(t, e)
	assert.Equal(t, 24*time.Hour, d)

}

func TestDays(t *testing.T) {
	d, e := parseDuration("8d")

	assert.Nil(t, e)
	assert.Equal(t, 8*24*time.Hour, d)
}

func TestWeek(t *testing.T) {
	d, e := parseDuration("1w")

	assert.Nil(t, e)
	assert.Equal(t, 7*24*time.Hour, d)
}
