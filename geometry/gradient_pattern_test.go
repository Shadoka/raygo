package geometry

import (
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestColorAtGradient(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)
	p := CreateGradientPattern(white, black)
	expected := math.CreateColor(0.75, 0.75, 0.75)
	expected2 := math.CreateColor(0.5, 0.5, 0.5)
	expected3 := math.CreateColor(0.25, 0.25, 0.25)

	assert.Assert(t, white.Equals(p.ColorAt(math.CreatePoint(0.0, 0.0, 0.0))))
	assert.Assert(t, expected.Equals(p.ColorAt(math.CreatePoint(0.25, 0.0, 0.0))))
	assert.Assert(t, expected2.Equals(p.ColorAt(math.CreatePoint(0.5, 0.0, 0.0))))
	assert.Assert(t, expected3.Equals(p.ColorAt(math.CreatePoint(0.75, 0.0, 0.0))))
}
