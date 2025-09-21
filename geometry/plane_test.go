package geometry

import (
	gomath "math"
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestLocalPlaneIntersectParallel(t *testing.T) {
	p := CreatePlane()
	r := CreateRay(math.CreatePoint(0.0, 10.0, 0.0), math.CreateVector(0.0, 0.0, 1.0))

	xs := p.localPlaneIntersect(r)

	assert.Assert(t, len(xs) == 0)
}

func TestLocalPlaneIntersectCoplanar(t *testing.T) {
	p := CreatePlane()
	r := CreateRay(math.CreatePoint(0.0, 0.0, 0.0), math.CreateVector(0.0, 0.0, 1.0))

	xs := p.localPlaneIntersect(r)

	assert.Assert(t, len(xs) == 0)
}

func TestLocalPlaneIntersectFromAbove(t *testing.T) {
	p := CreatePlane()
	r := CreateRay(math.CreatePoint(0.0, 1.0, 0.0), math.CreateVector(0.0, -1.0, 0.0))
	expected := CreateIntersection(1.0, p)

	xs := p.localPlaneIntersect(r)

	assert.Assert(t, len(xs) == 1)
	assert.Assert(t, expected.IntersectionAt == xs[0].IntersectionAt)
	assert.Assert(t, p.Equals(xs[0].Object))
}

func TestLocalPlaneIntersectFromBelow(t *testing.T) {
	p := CreatePlane()
	r := CreateRay(math.CreatePoint(0.0, -2.0, 0.0), math.CreateVector(0.0, 1.0, 0.0))
	expected := CreateIntersection(2.0, p)

	xs := p.localPlaneIntersect(r)

	assert.Assert(t, len(xs) == 1)
	assert.Assert(t, expected.IntersectionAt == xs[0].IntersectionAt)
	assert.Assert(t, p.Equals(xs[0].Object))
}

func TestNormalAtDefaultPlane(t *testing.T) {
	p := CreatePlane()
	point := math.CreatePoint(0.0, 0.0, 5.0)
	expected := math.CreateVector(0.0, 1.0, 0.0)

	actual := p.NormalAt(point, Intersection{})

	assert.Assert(t, expected.Equals(actual))
}

func TestNormalAtDefaultPlaneTranslated(t *testing.T) {
	p := CreatePlane()
	p.SetTransform(math.Translation(4.0, 0.0, 0.0))
	point := math.CreatePoint(0.0, 0.0, 5.0)
	expected := math.CreateVector(0.0, 1.0, 0.0)

	actual := p.NormalAt(point, Intersection{})

	assert.Assert(t, expected.Equals(actual))
}

func TestNormalAtDefaultPlaneRotatedX(t *testing.T) {
	p := CreatePlane()
	p.SetTransform(math.Rotation_X(gomath.Pi / 2.0))
	point := math.CreatePoint(0.0, 0.0, 5.0)
	expected := math.CreateVector(0.0, 0.0, 1.0)

	actual := p.NormalAt(point, Intersection{})

	assert.Assert(t, expected.Equals(actual))
}

func TestPlaneBoundsUntransformed(t *testing.T) {
	p := CreatePlane()
	b := p.Bounds()
	expected := Bounds{
		Minimum: math.CreatePoint(gomath.Inf(-1), -0.1, gomath.Inf(-1)),
		Maximum: math.CreatePoint(gomath.Inf(1), 0.1, gomath.Inf(1)),
	}

	assert.Assert(t, expected.Equals(b))
}

func TestPlaneBoundsTransformed(t *testing.T) {
	p := CreatePlane()
	p.SetTransform(math.Translation(1.0, 1.0, 1.0).MulM(math.Scaling(3.0, 3.0, 3.0)))
	b := p.Bounds()
	expected := Bounds{
		Minimum: math.CreatePoint(gomath.Inf(-1), -0.1, gomath.Inf(-1)),
		Maximum: math.CreatePoint(gomath.Inf(1), 0.1, gomath.Inf(1)),
	}

	assert.Assert(t, expected.Equals(b))
}
