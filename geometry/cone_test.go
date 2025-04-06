package geometry

import (
	gomath "math"
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestConeLocalIntersect(t *testing.T) {
	c := CreateCone()
	r1 := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0).Normalize())
	r2 := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(1.0, 1.0, 1.0).Normalize())
	r3 := CreateRay(math.CreatePoint(1.0, 1.0, -5.0), math.CreateVector(-0.5, -1.0, 1.0).Normalize())

	expected := []Intersection{
		CreateIntersection(5.0, c),
		CreateIntersection(5.0, c),
		CreateIntersection(8.66025, c),
		CreateIntersection(8.66025, c),
		CreateIntersection(4.55006, c),
		CreateIntersection(49.44994, c),
	}

	xs := c.localConeIntersect(r1)
	xs = append(xs, c.localConeIntersect(r2)...)
	xs = append(xs, c.localConeIntersect(r3)...)

	for i := range expected {
		assert.Assert(t, floatEquals(expected[i].IntersectionAt, xs[i].IntersectionAt))
	}
}

func TestConeLocalIntersectParallelRay(t *testing.T) {
	c := CreateCone()
	r1 := CreateRay(math.CreatePoint(0.0, 0.0, -1.0), math.CreateVector(0.0, 1.0, 1.0).Normalize())

	xs := c.localConeIntersect(r1)

	assert.Assert(t, len(xs) == 1)
	assert.Assert(t, floatEquals(xs[0].IntersectionAt, 0.35355))
}

func TestConeLocalIntersectWithCaps(t *testing.T) {
	c := CreateCone()
	c.Minimum = -0.5
	c.Maximum = 0.5
	c.Closed = true
	r1 := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 1.0, 0.0).Normalize())
	r2 := CreateRay(math.CreatePoint(0.0, 0.0, -0.25), math.CreateVector(0.0, 1.0, 1.0).Normalize())
	r3 := CreateRay(math.CreatePoint(0.0, 0.0, -0.25), math.CreateVector(0.0, 1.0, 0.0).Normalize())

	assert.Assert(t, len(c.localConeIntersect(r1)) == 0)
	assert.Assert(t, len(c.localConeIntersect(r2)) == 2)
	assert.Assert(t, len(c.localConeIntersect(r3)) == 4)
}

func TestConeLocalNormalAt(t *testing.T) {
	c := CreateCone()
	expected1 := math.CreateVector(0.0, 0.0, 0.0)
	expected2 := math.CreateVector(1.0, -gomath.Sqrt(2.0), 1.0)
	expected3 := math.CreateVector(-1.0, 1.0, 0.0)

	assert.Assert(t, expected1.Equals(c.localConeNormalAt(math.CreatePoint(0.0, 0.0, 0.0))))
	assert.Assert(t, expected2.Equals(c.localConeNormalAt(math.CreatePoint(1.0, 1.0, 1.0))))
	assert.Assert(t, expected3.Equals(c.localConeNormalAt(math.CreatePoint(-1.0, -1.0, 0.0))))
}

func TestConeBoundsUntransformed(t *testing.T) {
	c := CreateCone()
	b := c.Bounds()
	expected := Bounds{
		Minimum: math.CreatePoint(-1.0, gomath.Inf(-1), -1.0),
		Maximum: math.CreatePoint(1.0, gomath.Inf(1), 1.0),
	}

	assert.Assert(t, expected.Equals(b))
}

func TestConeBoundsScaled(t *testing.T) {
	c := CreateCone()
	c.SetTransform(math.Scaling(3.0, 3.0, 3.0))
	b := c.Bounds()
	expected := Bounds{
		Minimum: math.CreatePoint(-3.0, gomath.Inf(-1), -3.0),
		Maximum: math.CreatePoint(3.0, gomath.Inf(1), 3.0),
	}

	assert.Assert(t, expected.Equals(b))
}

func TestConeBoundsTransformed(t *testing.T) {
	c := CreateCone()
	c.SetTransform(math.Translation(1.0, 1.0, 1.0).MulM(math.Scaling(3.0, 3.0, 3.0)))
	b := c.Bounds()
	expected := Bounds{
		Minimum: math.CreatePoint(-2.0, gomath.Inf(-1), -2.0),
		Maximum: math.CreatePoint(4.0, gomath.Inf(1), 4.0),
	}

	assert.Assert(t, expected.Equals(b))
}

func TestConeCappedBoundsTransformed(t *testing.T) {
	c := CreateCone()
	c.Minimum = -1.0
	c.Maximum = 1.0
	c.Closed = true
	c.SetTransform(math.Translation(1.0, 1.0, 1.0).MulM(math.Scaling(3.0, 3.0, 3.0)))
	b := c.Bounds()
	expected := Bounds{
		Minimum: math.CreatePoint(-2.0, -2.0, -2.0),
		Maximum: math.CreatePoint(4.0, 4.0, 4.0),
	}

	assert.Assert(t, expected.Equals(b))
}
