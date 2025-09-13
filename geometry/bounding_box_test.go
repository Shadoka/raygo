package geometry

import (
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestBoundingBoxIntersectXFaces(t *testing.T) {
	c := CreateCube()
	b := c.Bounds()
	rFromRight := CreateRay(math.CreatePoint(5.0, 0.5, 0.0), math.CreateVector(-1.0, 0.0, 0.0))
	rFromLeft := CreateRay(math.CreatePoint(-5.0, 0.5, 0.0), math.CreateVector(1.0, 0.0, 0.0))

	xsFromRight := BoundingBoxIntersect(rFromRight, c, b.Minimum, b.Maximum)
	xsFromLeft := BoundingBoxIntersect(rFromLeft, c, b.Minimum, b.Maximum)

	assert.Assert(t, xsFromRight[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromRight[1].IntersectionAt == 6.0)
	assert.Assert(t, xsFromLeft[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromLeft[1].IntersectionAt == 6.0)
}

func TestBoundingBoxIntersectYFaces(t *testing.T) {
	c := CreateCube()
	b := c.Bounds()
	rFromTop := CreateRay(math.CreatePoint(0.5, 5.0, 0.0), math.CreateVector(0.0, -1.0, 0.0))
	rFromBottom := CreateRay(math.CreatePoint(0.5, -5.0, 0.0), math.CreateVector(0.0, 1.0, 0.0))

	xsFromTop := BoundingBoxIntersect(rFromTop, c, b.Minimum, b.Maximum)
	xsFromBottom := BoundingBoxIntersect(rFromBottom, c, b.Minimum, b.Maximum)

	assert.Assert(t, xsFromTop[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromTop[1].IntersectionAt == 6.0)
	assert.Assert(t, xsFromBottom[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromBottom[1].IntersectionAt == 6.0)
}

func TestBoundingBoxIntersectZFaces(t *testing.T) {
	c := CreateCube()
	b := c.Bounds()
	rFromBehind := CreateRay(math.CreatePoint(0.5, 0.0, 5.0), math.CreateVector(0.0, 0.0, -1.0))
	rFromFront := CreateRay(math.CreatePoint(0.5, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))

	xsFromBehind := BoundingBoxIntersect(rFromBehind, c, b.Minimum, b.Maximum)
	xsFromFront := BoundingBoxIntersect(rFromFront, c, b.Minimum, b.Maximum)

	assert.Assert(t, xsFromBehind[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromBehind[1].IntersectionAt == 6.0)
	assert.Assert(t, xsFromFront[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromFront[1].IntersectionAt == 6.0)
}

func TestBoundingBoxIntersectFromInside(t *testing.T) {
	c := CreateCube()
	b := c.Bounds()
	r := CreateRay(math.CreatePoint(0.0, 0.5, 0.0), math.CreateVector(0.0, 0.0, 1.0))

	xsFromBehind := BoundingBoxIntersect(r, c, b.Minimum, b.Maximum)

	assert.Assert(t, xsFromBehind[0].IntersectionAt == -1.0)
	assert.Assert(t, xsFromBehind[1].IntersectionAt == 1.0)
}

func TestBoundingBoxIntersectRayMiss(t *testing.T) {
	c := CreateCube()
	b := c.Bounds()
	r1 := CreateRay(math.CreatePoint(-2.0, 0.0, 0.0), math.CreateVector(0.2673, 0.5345, 0.8018))
	r2 := CreateRay(math.CreatePoint(0.0, -2.0, 0.0), math.CreateVector(0.8018, 0.2673, 0.5345))
	r3 := CreateRay(math.CreatePoint(0.0, 0.0, -2.0), math.CreateVector(0.5345, 0.8018, 0.2673))
	r4 := CreateRay(math.CreatePoint(2.0, 0.0, 2.0), math.CreateVector(0.0, 0.0, -1.0))
	r5 := CreateRay(math.CreatePoint(0.0, 2.0, 2.0), math.CreateVector(0.0, -1.0, 0.0))
	r6 := CreateRay(math.CreatePoint(2.0, 2.0, 0.0), math.CreateVector(-1.0, 0.0, 0.0))

	xs := BoundingBoxIntersect(r1, c, b.Minimum, b.Maximum)
	xs = append(xs, BoundingBoxIntersect(r2, c, b.Minimum, b.Maximum)...)
	xs = append(xs, BoundingBoxIntersect(r3, c, b.Minimum, b.Maximum)...)
	xs = append(xs, BoundingBoxIntersect(r4, c, b.Minimum, b.Maximum)...)
	xs = append(xs, BoundingBoxIntersect(r5, c, b.Minimum, b.Maximum)...)
	xs = append(xs, BoundingBoxIntersect(r6, c, b.Minimum, b.Maximum)...)

	assert.Assert(t, len(xs) == 0)
}
