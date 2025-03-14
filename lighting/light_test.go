package lighting

import (
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateLight(t *testing.T) {
	p := math.CreatePoint(0.0, 0.0, 0.0)
	intensity := math.CreateColor(1.0, 1.0, 1.0)

	light := CreateLight(p, intensity)

	assert.Assert(t, p.Equals(light.Position))
	assert.Assert(t, intensity.Equals(light.Intensity))
}
