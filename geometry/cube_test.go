package geometry

import (
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCubeIntersectXFaces(t *testing.T) {
	c := CreateCube()
	rFromRight := CreateRay(math.CreatePoint(5.0, 0.5, 0.0), math.CreateVector(-1.0, 0.0, 0.0))
	rFromLeft := CreateRay(math.CreatePoint(-5.0, 0.5, 0.0), math.CreateVector(1.0, 0.0, 0.0))

	xsFromRight := c.localCubeIntersect(rFromRight)
	xsFromLeft := c.localCubeIntersect(rFromLeft)

	assert.Assert(t, xsFromRight[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromRight[1].IntersectionAt == 6.0)
	assert.Assert(t, xsFromLeft[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromLeft[1].IntersectionAt == 6.0)
}

func TestCubeIntersectYFaces(t *testing.T) {
	c := CreateCube()
	rFromTop := CreateRay(math.CreatePoint(0.5, 5.0, 0.0), math.CreateVector(0.0, -1.0, 0.0))
	rFromBottom := CreateRay(math.CreatePoint(0.5, -5.0, 0.0), math.CreateVector(0.0, 1.0, 0.0))

	xsFromTop := c.localCubeIntersect(rFromTop)
	xsFromBottom := c.localCubeIntersect(rFromBottom)

	assert.Assert(t, xsFromTop[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromTop[1].IntersectionAt == 6.0)
	assert.Assert(t, xsFromBottom[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromBottom[1].IntersectionAt == 6.0)
}

func TestCubeIntersectZFaces(t *testing.T) {
	c := CreateCube()
	rFromBehind := CreateRay(math.CreatePoint(0.5, 0.0, 5.0), math.CreateVector(0.0, 0.0, -1.0))
	rFromFront := CreateRay(math.CreatePoint(0.5, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))

	xsFromBehind := c.localCubeIntersect(rFromBehind)
	xsFromFront := c.localCubeIntersect(rFromFront)

	assert.Assert(t, xsFromBehind[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromBehind[1].IntersectionAt == 6.0)
	assert.Assert(t, xsFromFront[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromFront[1].IntersectionAt == 6.0)
}

func TestCubeIntersectFromInside(t *testing.T) {
	c := CreateCube()
	r := CreateRay(math.CreatePoint(0.0, 0.5, 0.0), math.CreateVector(0.0, 0.0, 1.0))

	xsFromBehind := c.localCubeIntersect(r)

	assert.Assert(t, xsFromBehind[0].IntersectionAt == -1.0)
	assert.Assert(t, xsFromBehind[1].IntersectionAt == 1.0)
}

func TestCubeIntersectRayMiss(t *testing.T) {
	c := CreateCube()
	r1 := CreateRay(math.CreatePoint(-2.0, 0.0, 0.0), math.CreateVector(0.2673, 0.5345, 0.8018))
	r2 := CreateRay(math.CreatePoint(0.0, -2.0, 0.0), math.CreateVector(0.8018, 0.2673, 0.5345))
	r3 := CreateRay(math.CreatePoint(0.0, 0.0, -2.0), math.CreateVector(0.5345, 0.8018, 0.2673))
	r4 := CreateRay(math.CreatePoint(2.0, 0.0, 2.0), math.CreateVector(0.0, 0.0, -1.0))
	r5 := CreateRay(math.CreatePoint(0.0, 2.0, 2.0), math.CreateVector(0.0, -1.0, 0.0))
	r6 := CreateRay(math.CreatePoint(2.0, 2.0, 0.0), math.CreateVector(-1.0, 0.0, 0.0))

	xs := c.localCubeIntersect(r1)
	xs = append(xs, c.localCubeIntersect(r2)...)
	xs = append(xs, c.localCubeIntersect(r3)...)
	xs = append(xs, c.localCubeIntersect(r4)...)
	xs = append(xs, c.localCubeIntersect(r5)...)
	xs = append(xs, c.localCubeIntersect(r6)...)

	assert.Assert(t, len(xs) == 0)
}

func TestCubeLocalNormalAt(t *testing.T) {
	c := CreateCube()
	p1 := math.CreatePoint(1.0, 0.5, -0.8)
	p2 := math.CreatePoint(-1.0, -0.2, 0.9)
	p3 := math.CreatePoint(-0.4, 1.0, -0.1)
	p4 := math.CreatePoint(0.3, -1.0, -0.7)
	p5 := math.CreatePoint(-0.6, 0.3, 1.0)
	p6 := math.CreatePoint(0.4, 0.4, -1.0)
	p7 := math.CreatePoint(1.0, 1.0, 1.0)
	p8 := math.CreatePoint(-1.0, -1.0, -1.0)
	expected1 := math.CreateVector(1.0, 0.0, 0.0)
	expected2 := math.CreateVector(-1.0, 0.0, 0.0)
	expected3 := math.CreateVector(0.0, 1.0, 0.0)
	expected4 := math.CreateVector(0.0, -1.0, 0.0)
	expected5 := math.CreateVector(0.0, 0.0, 1.0)
	expected6 := math.CreateVector(0.0, 0.0, -1.0)
	expected7 := math.CreateVector(1.0, 0.0, 0.0)
	expected8 := math.CreateVector(-1.0, 0.0, 0.0)

	assert.Assert(t, expected1.Equals(c.localCubeNormalAt(p1)))
	assert.Assert(t, expected2.Equals(c.localCubeNormalAt(p2)))
	assert.Assert(t, expected3.Equals(c.localCubeNormalAt(p3)))
	assert.Assert(t, expected4.Equals(c.localCubeNormalAt(p4)))
	assert.Assert(t, expected5.Equals(c.localCubeNormalAt(p5)))
	assert.Assert(t, expected6.Equals(c.localCubeNormalAt(p6)))
	assert.Assert(t, expected7.Equals(c.localCubeNormalAt(p7)))
	assert.Assert(t, expected8.Equals(c.localCubeNormalAt(p8)))
}

func TestCubeBoundsUntransformed(t *testing.T) {
	c := CreateCube()
	b := c.ScaledBounds()
	expected := Bounds{
		Minimum: math.CreatePoint(-1.0, -1.0, -1.0),
		Maximum: math.CreatePoint(1.0, 1.0, 1.0),
	}

	assert.Assert(t, expected.Equals(b))
}

func TestCubeBoundsScaled(t *testing.T) {
	c := CreateCube()
	c.SetTransform(math.Scaling(3.0, 3.0, 3.0))
	b := c.ScaledBounds()
	expected := Bounds{
		Minimum: math.CreatePoint(-3.0, -3.0, -3.0),
		Maximum: math.CreatePoint(3.0, 3.0, 3.0),
	}

	assert.Assert(t, expected.Equals(b))
}

func TestCubeBoundsTransformed(t *testing.T) {
	c := CreateCube()
	c.SetTransform(math.Translation(1.0, 1.0, 1.0).MulM(math.Scaling(3.0, 3.0, 3.0)))
	b := c.ScaledBounds()
	expected := Bounds{
		Minimum: math.CreatePoint(-2.0, -2.0, -2.0),
		Maximum: math.CreatePoint(4.0, 4.0, 4.0),
	}

	assert.Assert(t, expected.Equals(b))
}

func TestCubeScaledLocalIntersect(t *testing.T) {
	c := CreateCube()
	c.SetTransform(math.Scaling(2.0, 2.0, 2.0).MulM(math.Translation(5.0, 0.0, 0.0)))
	r := CreateRay(math.CreatePoint(10.0, 0.0, -10.0), math.CreateVector(0.0, 0.0, 1.0))

	xs := c.Intersect(r)

	assert.Assert(t, len(xs) == 2)
}
