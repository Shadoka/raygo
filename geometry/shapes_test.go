package geometry

import (
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestFindBoundingBox(t *testing.T) {
	p1 := math.CreatePoint(-1.0, -1.0, -1.0)
	p2 := math.CreatePoint(-1.0, -1.0, 1.0)
	p3 := math.CreatePoint(-2.0, 1.0, -1.0)
	p4 := math.CreatePoint(-2.0, 1.0, 1.0)
	p5 := math.CreatePoint(1.0, -1.0, 1.0)
	p6 := math.CreatePoint(1.0, -1.0, -1.0)
	p7 := math.CreatePoint(2.0, 1.0, 1.0)
	p8 := math.CreatePoint(2.0, 1.0, -1.0)
	corners := make([]math.Point, 0)
	corners = append(corners, p1, p2, p3, p4, p5, p6, p7, p8)

	expected := Bounds{
		Minimum: math.CreatePoint(-2.0, -1.0, -1.0),
		Maximum: math.CreatePoint(2.0, 1.0, 1.0),
	}

	assert.Assert(t, expected.Equals(FindBoundingBox(corners)))
}
