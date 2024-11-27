package tim

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmpty(t *testing.T) {
	_, _, err := parseDuration("")
	if err == nil {
		t.Fail()
	}

}

func TestDay(t *testing.T) {
	d, ds, err := parseDuration("1d")

	assert.Nil(t, err)
	assert.Equal(t, int64(0), d.Milliseconds())
	assert.Equal(t, 1, ds.days)
}

func TestDays(t *testing.T) {
	d, ds, err := parseDuration("8d")

	assert.Nil(t, err)
	assert.Equal(t, int64(0), d.Milliseconds())
	assert.Equal(t, 8, ds.days)
}

func TestMonths(t *testing.T) {
	d, ds, err := parseDuration("9M")

	assert.Nil(t, err)
	assert.Equal(t, int64(0), d.Milliseconds())
	assert.Equal(t, 9, ds.months)
}

func TestWeek(t *testing.T) {
	d, ds, err := parseDuration("7y")

	assert.Nil(t, err)
	assert.Equal(t, int64(0), d.Milliseconds())
	assert.Equal(t, 7, ds.years)
}
