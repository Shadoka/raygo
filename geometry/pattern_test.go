package geometry

import (
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestRingPattern(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)
	rp := CreateRingPattern(white, black)

	assert.Assert(t, white.Equals(rp.ColorAt(math.CreatePoint(0.0, 0.0, 0.0))))
	assert.Assert(t, black.Equals(rp.ColorAt(math.CreatePoint(1.0, 0.0, 0.0))))
	assert.Assert(t, black.Equals(rp.ColorAt(math.CreatePoint(0.0, 0.0, 1.0))))
	assert.Assert(t, black.Equals(rp.ColorAt(math.CreatePoint(0.708, 0.0, 0.708))))
}

func TestCheckerPatternRepeatX(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)
	c := CreateCheckerPattern(white, black)

	assert.Assert(t, white.Equals(c.ColorAt(math.CreatePoint(0.0, 0.0, 0.0))))
	assert.Assert(t, white.Equals(c.ColorAt(math.CreatePoint(0.99, 0.0, 0.0))))
	assert.Assert(t, black.Equals(c.ColorAt(math.CreatePoint(1.01, 0.0, 0.0))))
}

func TestCheckerPatternRepeatZ(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)
	c := CreateCheckerPattern(white, black)

	assert.Assert(t, white.Equals(c.ColorAt(math.CreatePoint(0.0, 0.0, 0.0))))
	assert.Assert(t, white.Equals(c.ColorAt(math.CreatePoint(0.0, 0.0, 0.99))))
	assert.Assert(t, black.Equals(c.ColorAt(math.CreatePoint(0.0, 0.0, 1.01))))
}

func TestCheckerPatternRepeatY(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)
	c := CreateCheckerPattern(white, black)

	assert.Assert(t, white.Equals(c.ColorAt(math.CreatePoint(0.0, 0.0, 0.0))))
	assert.Assert(t, white.Equals(c.ColorAt(math.CreatePoint(0.0, 0.99, 0.0))))
	assert.Assert(t, black.Equals(c.ColorAt(math.CreatePoint(0.0, 1.01, 0.0))))
}
