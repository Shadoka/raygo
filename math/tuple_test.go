package math

import (
	gomath "math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreatePoint(t *testing.T) {
	x := 4.3
	y := -4.2
	z := 3.1
	point := CreatePoint(x, y, z)

	assert.Assert(t, point.IsPoint())
	assert.Assert(t, !point.IsVector())
	assert.Assert(t, point.X == x)
	assert.Assert(t, point.Y == y)
	assert.Assert(t, point.Z == z)
}

func TestCreateVector(t *testing.T) {
	x := 4.3
	y := -4.2
	z := 3.1
	vector := CreateVector(x, y, z)

	assert.Assert(t, !vector.IsPoint())
	assert.Assert(t, vector.IsVector())
	assert.Assert(t, vector.X == x)
	assert.Assert(t, vector.Y == y)
	assert.Assert(t, vector.Z == z)
}

func TestEqualsOk(t *testing.T) {
	x := 4.3
	y := -4.2
	z := 3.1
	point1 := CreatePoint(x, y, z)

	point2 := CreatePoint(x, y, z)

	assert.Assert(t, point1.Equals(point2))
}

func TestEqualsFail(t *testing.T) {
	x := 4.3
	y := -4.2
	z := 3.1
	point := CreatePoint(x, y, z)

	vector := CreateVector(x, y, z)

	assert.Assert(t, !point.Equals(vector))
}

func TestAddTuple(t *testing.T) {
	tuple1 := Tuple{
		X: 3.0,
		Y: -2.0,
		Z: 5.0,
		W: 1.0,
	}
	tuple2 := Tuple{
		X: -2.0,
		Y: 3.0,
		Z: 1.0,
		W: 0.0,
	}
	expectedTuple := Tuple{
		X: 1.0,
		Y: 1.0,
		Z: 6.0,
		W: 1.0,
	}
	actualTuple := tuple1.Add(tuple2)

	assert.Assert(t, expectedTuple.Equals(actualTuple))
}

func TestSubtractTwoPoints(t *testing.T) {
	p1 := CreatePoint(3.0, 2.0, 1.0)
	p2 := CreatePoint(5.0, 6.0, 7.0)
	expectedVector := Tuple{
		X: -2.0,
		Y: -4.0,
		Z: -6.0,
		W: 0.0,
	}
	actualVector := p1.Subtract(p2)

	assert.Assert(t, expectedVector.Equals(actualVector))
	assert.Assert(t, actualVector.IsVector())
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p := CreatePoint(3.0, 2.0, 1.0)
	v := CreateVector(5.0, 6.0, 7.0)
	expectedPoint := Tuple{
		X: -2.0,
		Y: -4.0,
		Z: -6.0,
		W: 1.0,
	}
	actualPoint := p.Subtract(v)

	assert.Assert(t, expectedPoint.Equals(actualPoint))
	assert.Assert(t, actualPoint.IsPoint())
}

func TestSubtractTwoVectors(t *testing.T) {
	v1 := CreateVector(3.0, 2.0, 1.0)
	v2 := CreateVector(5.0, 6.0, 7.0)
	expectedVector := Tuple{
		X: -2.0,
		Y: -4.0,
		Z: -6.0,
		W: 0.0,
	}
	actualVector := v1.Subtract(v2)

	assert.Assert(t, expectedVector.Equals(actualVector))
	assert.Assert(t, actualVector.IsVector())
}

func TestNegateTuple(t *testing.T) {
	tup := Tuple{
		X: 1.0,
		Y: -2.0,
		Z: 3.0,
		W: -4.0,
	}
	expectedTuple := Tuple{
		X: -1.0,
		Y: 2.0,
		Z: -3.0,
		W: 4.0,
	}

	assert.Assert(t, expectedTuple.Equals(tup.Negate()))
}

func TestMulTuple(t *testing.T) {
	tup := Tuple{
		X: 1.0,
		Y: -2.0,
		Z: 3.0,
		W: -4.0,
	}
	expectedTuple := Tuple{
		X: 3.5,
		Y: -7.0,
		Z: 10.5,
		W: -14.0,
	}

	assert.Assert(t, expectedTuple.Equals(tup.Mul(3.5)))
}

func TestMulTupleSmallScalar(t *testing.T) {
	tup := Tuple{
		X: 1.0,
		Y: -2.0,
		Z: 3.0,
		W: -4.0,
	}
	expectedTuple := Tuple{
		X: 0.5,
		Y: -1.0,
		Z: 1.5,
		W: -2.0,
	}

	assert.Assert(t, expectedTuple.Equals(tup.Mul(0.5)))
}

func TestDivTuple(t *testing.T) {
	tup := Tuple{
		X: 1.0,
		Y: -2.0,
		Z: 3.0,
		W: -4.0,
	}
	expectedTuple := Tuple{
		X: 0.5,
		Y: -1.0,
		Z: 1.5,
		W: -2.0,
	}

	assert.Assert(t, expectedTuple.Equals(tup.Div(2)))
}

func TestMagnitudeSpecial1(t *testing.T) {
	v := CreateVector(1.0, 0.0, 0.0)
	expectedMagnitude := 1.0

	assert.Assert(t, floatEquals(expectedMagnitude, v.Magnitude()))
}

func TestMagnitudeSpecial2(t *testing.T) {
	v := CreateVector(0.0, 1.0, 0.0)
	expectedMagnitude := 1.0

	assert.Assert(t, floatEquals(expectedMagnitude, v.Magnitude()))
}

func TestMagnitudeSpecial3(t *testing.T) {
	v := CreateVector(0.0, 0.0, 1.0)
	expectedMagnitude := 1.0

	assert.Assert(t, floatEquals(expectedMagnitude, v.Magnitude()))
}

func TestMagnitudeRandomPositiveValues(t *testing.T) {
	v := CreateVector(1.0, 2.0, 3.0)
	expectedMagnitude := gomath.Sqrt(14)

	assert.Assert(t, floatEquals(expectedMagnitude, v.Magnitude()))
}

func TestMagnitudeRandomNegativeValues(t *testing.T) {
	v := CreateVector(-1.0, -2.0, -3.0)
	expectedMagnitude := gomath.Sqrt(14)

	assert.Assert(t, floatEquals(expectedMagnitude, v.Magnitude()))
}

func TestNormalizeVector1(t *testing.T) {
	v := CreateVector(4.0, 0.0, 0.0)
	expectedVector := CreateVector(1.0, 0.0, 0.0)

	assert.Assert(t, expectedVector.Equals(v.Normalize()))
}

func TestNormalizeVector2(t *testing.T) {
	v := CreateVector(1.0, 2.0, 3.0)
	expectedVector := CreateVector(0.26726, 0.53452, 0.80178)

	assert.Assert(t, expectedVector.Equals(v.Normalize()))
}

func TestMagnitudeOfNormalizedVector(t *testing.T) {
	v := CreateVector(1.0, 2.0, 3.0)

	assert.Assert(t, floatEquals(1.0, v.Normalize().Magnitude()))
}

func TestDotProduct(t *testing.T) {
	v1 := CreateVector(1.0, 2.0, 3.0)
	v2 := CreateVector(2.0, 3.0, 4.0)
	expectedValue := 20.0

	assert.Assert(t, floatEquals(expectedValue, v1.Dot(v2)))
}

func TestCrossProduct(t *testing.T) {
	v1 := CreateVector(1.0, 2.0, 3.0)
	v2 := CreateVector(2.0, 3.0, 4.0)
	expectedCrossProduct := CreateVector(-1.0, 2.0, -1.0)
	expectedReverseCrossProduct := CreateVector(1.0, -2.0, 1.0)

	assert.Assert(t, expectedCrossProduct.Equals(v1.Cross(v2)))
	assert.Assert(t, expectedReverseCrossProduct.Equals(v2.Cross(v1)))
}

func TestCreateColor(t *testing.T) {
	c := CreateColor(0.9, 0.6, 0.75)

	assert.Assert(t, floatEquals(0.9, c.X))
	assert.Assert(t, floatEquals(0.6, c.Y))
	assert.Assert(t, floatEquals(0.75, c.Z))
}

func TestBlendColor(t *testing.T) {
	c1 := CreateColor(1.0, 0.2, 0.4)
	c2 := CreateColor(0.9, 1.0, 0.1)
	expectedColor := CreateColor(0.9, 0.2, 0.04)

	assert.Assert(t, expectedColor.Equals(c1.Blend(c2)))
}
