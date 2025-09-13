package groups

import (
	"raygo/geometry"
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestHexagonCorner(t *testing.T) {
	corner := hexagonCorner()

	expected := geometry.Bounds{
		Minimum: math.CreatePoint(-0.25, -0.25, -1.25),
		Maximum: math.CreatePoint(0.25, 0.25, 0.0),
	}

	actual := corner.Bounds().ApplyTransform(corner.Transform)
	assert.Assert(t, expected.Equals(actual))
}
