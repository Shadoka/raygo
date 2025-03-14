package ray

import (
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateRay(t *testing.T) {
	o := math.CreatePoint(1.0, 2.0, 3.0)
	d := math.CreateVector(4.0, 5.0, 6.0)
	ray := CreateRay(o, d)

	assert.Assert(t, o.Equals(ray.Origin))
	assert.Assert(t, d.Equals(ray.Direction))

	o = math.CreatePoint(2.0, 3.0, 4.0)
	assert.Assert(t, !o.Equals(ray.Origin))
}

func TestPosition(t *testing.T) {
	r := CreateRay(math.CreatePoint(2.0, 3.0, 4.0), math.CreateVector(1.0, 0.0, 0.0))
	expected1 := math.CreatePoint(2.0, 3.0, 4.0)
	expected2 := math.CreatePoint(3.0, 3.0, 4.0)
	expected3 := math.CreatePoint(1.0, 3.0, 4.0)
	expected4 := math.CreatePoint(4.5, 3.0, 4.0)

	assert.Assert(t, expected1.Equals(r.Position(0.0)))
	assert.Assert(t, expected2.Equals(r.Position(1.0)))
	assert.Assert(t, expected3.Equals(r.Position(-1.0)))
	assert.Assert(t, expected4.Equals(r.Position(2.5)))
}

func TestCreateIntersection(t *testing.T) {
	s := CreateSphere()
	i := CreateIntersection(3.5, s)

	assert.Assert(t, i.IntersectionAt == 3.5)
	assert.Assert(t, s.Equals(i.Object))
}

func TestAggregate(t *testing.T) {
	s := CreateSphere()
	i1 := CreateIntersection(1.0, s)
	i2 := CreateIntersection(2.0, s)
	xs := i1.Aggregate(i2)

	assert.Assert(t, len(xs) == 2)
	assert.Assert(t, xs[0].IntersectionAt == 1.0)
	assert.Assert(t, xs[1].IntersectionAt == 2.0)
}

func TestHitAllPositiveT(t *testing.T) {
	s := CreateSphere()
	i1 := CreateIntersection(1.0, s)
	i2 := CreateIntersection(2.0, s)
	xs := i1.Aggregate(i2)

	i := Hit(xs)

	assert.Assert(t, i1.Equals(*i))
}

func TestHitSomeNegativeT(t *testing.T) {
	s := CreateSphere()
	i1 := CreateIntersection(-1.0, s)
	i2 := CreateIntersection(1.0, s)
	xs := i1.Aggregate(i2)

	i := Hit(xs)

	assert.Assert(t, i2.Equals(*i))
}

func TestHitAllNegativeT(t *testing.T) {
	s := CreateSphere()
	i1 := CreateIntersection(-1.0, s)
	i2 := CreateIntersection(-2.0, s)
	xs := i1.Aggregate(i2)

	i := Hit(xs)

	assert.Assert(t, i == nil)
}

func TestHitRandomOrder(t *testing.T) {
	s := CreateSphere()
	i1 := CreateIntersection(5.0, s)
	i2 := CreateIntersection(7.0, s)
	i3 := CreateIntersection(-3.0, s)
	i4 := CreateIntersection(2.0, s)
	xs := i1.Aggregate(i2)
	xs = AddIntersection(xs, i3)
	xs = AddIntersection(xs, i4)

	i := Hit(xs)

	assert.Assert(t, i4.Equals(*i))
}

func TestTranslateRay(t *testing.T) {
	r := CreateRay(math.CreatePoint(1.0, 2.0, 3.0), math.CreateVector(0.0, 1.0, 0.0))
	m := math.Translation(3.0, 4.0, 5.0)
	expectedOrigin := math.CreatePoint(4.0, 6.0, 8.0)
	expectedDirection := math.CreateVector(0.0, 1.0, 0.0)

	r2 := r.Transform(m)

	assert.Assert(t, expectedOrigin.Equals(r2.Origin))
	assert.Assert(t, expectedDirection.Equals(r2.Direction))
}

func TestScaleRay(t *testing.T) {
	r := CreateRay(math.CreatePoint(1.0, 2.0, 3.0), math.CreateVector(0.0, 1.0, 0.0))
	m := math.Scaling(2.0, 3.0, 4.0)
	expectedOrigin := math.CreatePoint(2.0, 6.0, 12.0)
	expectedDirection := math.CreateVector(0.0, 3.0, 0.0)

	r2 := r.Transform(m)

	assert.Assert(t, expectedOrigin.Equals(r2.Origin))
	assert.Assert(t, expectedDirection.Equals(r2.Direction))
}

func TestPrepareComputation(t *testing.T) {
	r := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))
	s := CreateSphere()
	i := CreateIntersection(4.0, s)
	expected := IntersectionComputations{
		IntersectionAt: i.IntersectionAt,
		Object:         i.Object,
		Point:          math.CreatePoint(0.0, 0.0, -1.0),
		Eyev:           math.CreateVector(0.0, 0.0, -1.0),
		Normalv:        math.CreateVector(0.0, 0.0, -1.0),
		Inside:         false,
	}

	actual := i.PrepareComputation(r)

	assert.Assert(t, expected.IntersectionAt == actual.IntersectionAt)
	assert.Assert(t, expected.Object.Equals(actual.Object))
	assert.Assert(t, expected.Point.Equals(actual.Point))
	assert.Assert(t, expected.Eyev.Equals(actual.Eyev))
	assert.Assert(t, expected.Normalv.Equals(actual.Normalv))
	assert.Assert(t, expected.Inside == actual.Inside)
}

func TestPrepareComputationRayInside(t *testing.T) {
	r := CreateRay(math.CreatePoint(0.0, 0.0, 0.0), math.CreateVector(0.0, 0.0, 1.0))
	s := CreateSphere()
	i := CreateIntersection(1.0, s)
	expected := IntersectionComputations{
		IntersectionAt: i.IntersectionAt,
		Object:         i.Object,
		Point:          math.CreatePoint(0.0, 0.0, 1.0),
		Eyev:           math.CreateVector(0.0, 0.0, -1.0),
		Normalv:        math.CreateVector(0.0, 0.0, -1.0),
		Inside:         true,
	}

	actual := i.PrepareComputation(r)

	assert.Assert(t, expected.IntersectionAt == actual.IntersectionAt)
	assert.Assert(t, expected.Object.Equals(actual.Object))
	assert.Assert(t, expected.Point.Equals(actual.Point))
	assert.Assert(t, expected.Eyev.Equals(actual.Eyev))
	assert.Assert(t, expected.Normalv.Equals(actual.Normalv))
	assert.Assert(t, expected.Inside == actual.Inside)
}
