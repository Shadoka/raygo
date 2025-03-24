package geometry

import (
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCylinderLocalIntersectMisses(t *testing.T) {
	c := CreateCylinder()
	r1 := CreateRay(math.CreatePoint(1.0, 0.0, 0.0), math.CreateVector(0.0, 1.0, 0.0).Normalize())
	r2 := CreateRay(math.CreatePoint(0.0, 0.0, 0.0), math.CreateVector(0.0, 1.0, 0.0).Normalize())
	r3 := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(1.0, 1.0, 1.0).Normalize())

	xs := c.localCylinderIntersect(r1)
	xs = append(xs, c.localCylinderIntersect(r2)...)
	xs = append(xs, c.localCylinderIntersect(r3)...)

	assert.Assert(t, len(xs) == 0)
}

func TestCylinderLocalIntersectHit(t *testing.T) {
	c := CreateCylinder()
	r1 := CreateRay(math.CreatePoint(1.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0).Normalize())
	r2 := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0).Normalize())
	r3 := CreateRay(math.CreatePoint(0.5, 0.0, -5.0), math.CreateVector(0.1, 1.0, 1.0).Normalize())

	expected := []Intersection{
		CreateIntersection(5.0, c),
		CreateIntersection(5.0, c),
		CreateIntersection(4.0, c),
		CreateIntersection(6.0, c),
		CreateIntersection(6.80798, c),
		CreateIntersection(7.08872, c),
	}

	xs := c.localCylinderIntersect(r1)
	xs = append(xs, c.localCylinderIntersect(r2)...)
	xs = append(xs, c.localCylinderIntersect(r3)...)

	for i := range expected {
		assert.Assert(t, floatEquals(expected[i].IntersectionAt, xs[i].IntersectionAt))
	}
}

func TestLocalNormalCylinder(t *testing.T) {
	c := CreateCylinder()
	expected1 := math.CreateVector(1.0, 0.0, 0.0)
	expected2 := math.CreateVector(0.0, 0.0, -1.0)
	expected3 := math.CreateVector(0.0, 0.0, 1.0)
	expected4 := math.CreateVector(-1.0, 0.0, 0.0)

	assert.Assert(t, expected1.Equals(c.localCylinderNormalAt(math.CreatePoint(1.0, 0.0, 0.0))))
	assert.Assert(t, expected2.Equals(c.localCylinderNormalAt(math.CreatePoint(0.0, 0.0, -1.0))))
	assert.Assert(t, expected3.Equals(c.localCylinderNormalAt(math.CreatePoint(0.0, 0.0, 1.0))))
	assert.Assert(t, expected4.Equals(c.localCylinderNormalAt(math.CreatePoint(-1.0, 0.0, 0.0))))
}

func TestTruncatedCylinderIntersect(t *testing.T) {
	c := CreateCylinder()
	c.Minimum = 1.0
	c.Maximum = 2.0
	r1 := CreateRay(math.CreatePoint(0.0, 1.5, 0.0), math.CreateVector(0.1, 1.0, 0.0).Normalize())
	r2 := CreateRay(math.CreatePoint(0.0, 3.0, -5.0), math.CreateVector(0.0, 0.0, 1.0).Normalize())
	r3 := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0).Normalize())
	r4 := CreateRay(math.CreatePoint(0.0, 2.0, -5.0), math.CreateVector(0.0, 0.0, 1.0).Normalize())
	r5 := CreateRay(math.CreatePoint(0.0, 1.0, -5.0), math.CreateVector(0.0, 0.0, 1.0).Normalize())
	r6 := CreateRay(math.CreatePoint(0.0, 1.5, -2.0), math.CreateVector(0.0, 0.0, 1.0).Normalize())

	assert.Assert(t, len(c.localCylinderIntersect(r1)) == 0)
	assert.Assert(t, len(c.localCylinderIntersect(r2)) == 0)
	assert.Assert(t, len(c.localCylinderIntersect(r3)) == 0)
	assert.Assert(t, len(c.localCylinderIntersect(r4)) == 0)
	assert.Assert(t, len(c.localCylinderIntersect(r5)) == 0)
	assert.Assert(t, len(c.localCylinderIntersect(r6)) == 2)
}

func TestClosedCylinderIntersect(t *testing.T) {
	c := CreateCylinder()
	c.Minimum = 1.0
	c.Maximum = 2.0
	c.Closed = true
	r1 := CreateRay(math.CreatePoint(0.0, 3.0, 0.0), math.CreateVector(0.0, -1.0, 0.0).Normalize())
	r2 := CreateRay(math.CreatePoint(0.0, 3.0, -2.0), math.CreateVector(0.0, -1.0, 2.0).Normalize())
	r3 := CreateRay(math.CreatePoint(0.0, 4.0, -2.0), math.CreateVector(0.0, -1.0, 1.0).Normalize())
	r4 := CreateRay(math.CreatePoint(0.0, 0.0, -2.0), math.CreateVector(0.0, 1.0, 2.0).Normalize())
	r5 := CreateRay(math.CreatePoint(0.0, -1.0, -2.0), math.CreateVector(0.0, 1.0, 1.0).Normalize())

	assert.Assert(t, len(c.localCylinderIntersect(r1)) == 2)
	assert.Assert(t, len(c.localCylinderIntersect(r2)) == 2)
	assert.Assert(t, len(c.localCylinderIntersect(r3)) == 2)
	assert.Assert(t, len(c.localCylinderIntersect(r4)) == 2)
	assert.Assert(t, len(c.localCylinderIntersect(r5)) == 2)
}

func TestLocalNormalClosedCylinder(t *testing.T) {
	c := CreateCylinder()
	c.Minimum = 1.0
	c.Maximum = 2.0
	c.Closed = true
	expected1 := math.CreateVector(0.0, -1.0, 0.0)
	expected2 := math.CreateVector(0.0, -1.0, 0.0)
	expected3 := math.CreateVector(0.0, -1.0, 0.0)
	expected4 := math.CreateVector(0.0, 1.0, 0.0)
	expected5 := math.CreateVector(0.0, 1.0, 0.0)
	expected6 := math.CreateVector(0.0, 1.0, 0.0)

	assert.Assert(t, expected1.Equals(c.localCylinderNormalAt(math.CreatePoint(0.0, 1.0, 0.0))))
	assert.Assert(t, expected2.Equals(c.localCylinderNormalAt(math.CreatePoint(0.5, 1.0, 0.0))))
	assert.Assert(t, expected3.Equals(c.localCylinderNormalAt(math.CreatePoint(0.0, 1.0, 0.5))))
	assert.Assert(t, expected4.Equals(c.localCylinderNormalAt(math.CreatePoint(0.0, 2.0, 0.0))))
	assert.Assert(t, expected5.Equals(c.localCylinderNormalAt(math.CreatePoint(0.5, 2.0, 0.0))))
	assert.Assert(t, expected6.Equals(c.localCylinderNormalAt(math.CreatePoint(0.0, 2.0, 0.5))))
}
