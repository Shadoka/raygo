package math

import (
	"fmt"
	"math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestTranslation(t *testing.T) {
	transform := Translation(5.0, -3.0, 2.0)
	p := CreatePoint(-3.0, 4.0, 5.0)
	expected := CreatePoint(2.0, 1.0, 7.0)

	assert.Assert(t, expected.Equals(transform.MulT(p)))
}

func TestTranslationWithInverse(t *testing.T) {
	transform := Translation(5.0, -3.0, 2.0)
	inverse := transform.Inverse()
	p := CreatePoint(-3.0, 4.0, 5.0)
	expected := CreatePoint(-8.0, 7.0, 3.0)

	assert.Assert(t, expected.Equals(inverse.MulT(p)))
}

func TestTranslationWithVector(t *testing.T) {
	transform := Translation(5.0, -3.0, 2.0)
	v := CreateVector(-3.0, 4.0, 5.0)

	assert.Assert(t, v.Equals(transform.MulT(v)))
}

func TestScaling(t *testing.T) {
	scaling := Scaling(2.0, 3.0, 4.0)
	p := CreatePoint(-4.0, 6.0, 8.0)
	expected := CreatePoint(-8.0, 18.0, 32.0)

	assert.Assert(t, expected.Equals(scaling.MulT(p)))
}

func TestScalingVector(t *testing.T) {
	scaling := Scaling(2.0, 3.0, 4.0)
	v := CreateVector(-4.0, 6.0, 8.0)
	expected := CreateVector(-8.0, 18.0, 32.0)

	assert.Assert(t, expected.Equals(scaling.MulT(v)))
}

func TestScalingInverse(t *testing.T) {
	scaling := Scaling(2.0, 3.0, 4.0)
	inverse := scaling.Inverse()
	v := CreateVector(-4.0, 6.0, 8.0)
	expected := CreateVector(-2.0, 2.0, 2.0)

	assert.Assert(t, expected.Equals(inverse.MulT(v)))
}

func TestReflectWithScaling(t *testing.T) {
	scaling := Scaling(-1.0, 1.0, 1.0)
	p := CreatePoint(2.0, 3.0, 4.0)
	expected := CreatePoint(-2.0, 3.0, 4.0)

	assert.Assert(t, expected.Equals(scaling.MulT(p)))
}

func TestRotateX(t *testing.T) {
	rotationHalfQuarter := Rotation_X(math.Pi / 4.0) // 45°
	rotationFullQuarter := Rotation_X(math.Pi / 2.0) // 90°
	p := CreatePoint(0.0, 1.0, 0.0)
	expectedHalfQuarterRotation := CreatePoint(0.0, math.Sqrt(2)/2, math.Sqrt(2)/2)
	expectedFullQuarterRotation := CreatePoint(0.0, 0.0, 1.0)

	assert.Assert(t, expectedHalfQuarterRotation.Equals(rotationHalfQuarter.MulT(p)))
	assert.Assert(t, expectedFullQuarterRotation.Equals(rotationFullQuarter.MulT(p)))
}

func TestRotateXInverse(t *testing.T) {
	rotationHalfQuarter := Rotation_X(math.Pi / 4.0) // 45°
	rotationInverse := rotationHalfQuarter.Inverse()
	p := CreatePoint(0.0, 1.0, 0.0)
	expectedHalfQuarterRotation := CreatePoint(0.0, math.Sqrt(2)/2, math.Sqrt(2)/2*-1.0)

	assert.Assert(t, expectedHalfQuarterRotation.Equals(rotationInverse.MulT(p)))
}

func TestRotateY(t *testing.T) {
	rotationHalfQuarter := Rotation_Y(math.Pi / 4.0) // 45°
	rotationFullQuarter := Rotation_Y(math.Pi / 2.0) // 90°
	p := CreatePoint(0.0, 0.0, 1.0)
	expectedHalfQuarterRotation := CreatePoint(math.Sqrt(2)/2, 0.0, math.Sqrt(2)/2)
	expectedFullQuarterRotation := CreatePoint(1.0, 0.0, 0.0)

	assert.Assert(t, expectedHalfQuarterRotation.Equals(rotationHalfQuarter.MulT(p)))
	assert.Assert(t, expectedFullQuarterRotation.Equals(rotationFullQuarter.MulT(p)))
}

func TestRotateZ(t *testing.T) {
	rotationHalfQuarter := Rotation_Z(math.Pi / 4.0) // 45°
	rotationFullQuarter := Rotation_Z(math.Pi / 2.0) // 90°
	p := CreatePoint(0.0, 1.0, 0.0)
	expectedHalfQuarterRotation := CreatePoint(math.Sqrt(2)/2*-1.0, math.Sqrt(2)/2, 0.0)
	expectedFullQuarterRotation := CreatePoint(-1.0, 0.0, 0.0)

	assert.Assert(t, expectedHalfQuarterRotation.Equals(rotationHalfQuarter.MulT(p)))
	assert.Assert(t, expectedFullQuarterRotation.Equals(rotationFullQuarter.MulT(p)))
}

func TestShearingXY(t *testing.T) {
	p := CreatePoint(2.0, 3.0, 4.0)
	shearing := Shearing(1.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	expected := CreatePoint(5.0, 3.0, 4.0)

	assert.Assert(t, expected.Equals(shearing.MulT(p)))
}

func TestShearingXZ(t *testing.T) {
	p := CreatePoint(2.0, 3.0, 4.0)
	shearing := Shearing(0.0, 1.0, 0.0, 0.0, 0.0, 0.0)
	expected := CreatePoint(6.0, 3.0, 4.0)

	assert.Assert(t, expected.Equals(shearing.MulT(p)))
}

func TestShearingYX(t *testing.T) {
	p := CreatePoint(2.0, 3.0, 4.0)
	shearing := Shearing(0.0, 0.0, 1.0, 0.0, 0.0, 0.0)
	expected := CreatePoint(2.0, 5.0, 4.0)

	assert.Assert(t, expected.Equals(shearing.MulT(p)))
}

func TestShearingYZ(t *testing.T) {
	p := CreatePoint(2.0, 3.0, 4.0)
	shearing := Shearing(0.0, 0.0, 0.0, 1.0, 0.0, 0.0)
	expected := CreatePoint(2.0, 7.0, 4.0)

	assert.Assert(t, expected.Equals(shearing.MulT(p)))
}

func TestShearingZX(t *testing.T) {
	p := CreatePoint(2.0, 3.0, 4.0)
	shearing := Shearing(0.0, 0.0, 0.0, 0.0, 1.0, 0.0)
	expected := CreatePoint(2.0, 3.0, 6.0)

	assert.Assert(t, expected.Equals(shearing.MulT(p)))
}

func TestShearingZY(t *testing.T) {
	p := CreatePoint(2.0, 3.0, 4.0)
	shearing := Shearing(0.0, 0.0, 0.0, 0.0, 0.0, 1.0)
	expected := CreatePoint(2.0, 3.0, 7.0)

	assert.Assert(t, expected.Equals(shearing.MulT(p)))
}

func TestTransformsSequence(t *testing.T) {
	p := CreatePoint(1.0, 0.0, 1.0)
	rot := Rotation_X(math.Pi / 2)
	scale := Scaling(5.0, 5.0, 5.0)
	transl := Translation(10.0, 5.0, 7.0)

	expectedPostRot := CreatePoint(1.0, -1.0, 0.0)
	expectedPostScale := CreatePoint(5.0, -5.0, 0)
	expectedPostTransl := CreatePoint(15.0, 0.0, 7.0)

	assert.Assert(t, expectedPostRot.Equals(rot.MulT(p)))
	assert.Assert(t, expectedPostScale.Equals(scale.MulT(expectedPostRot)))
	assert.Assert(t, expectedPostTransl.Equals(transl.MulT(expectedPostScale)))
}

func TestTransformsChained(t *testing.T) {
	p := CreatePoint(1.0, 0.0, 1.0)
	rot := Rotation_X(math.Pi / 2)
	scale := Scaling(5.0, 5.0, 5.0)
	transl := Translation(10.0, 5.0, 7.0)

	expected := CreatePoint(15.0, 0.0, 7.0)

	assert.Assert(t, expected.Equals(transl.MulM(scale).MulM(rot).MulT(p)))
}

func TestViewTransformDefault(t *testing.T) {
	from := CreatePoint(0.0, 0.0, 0.0)
	to := CreatePoint(0.0, 0.0, -1.0)
	up := CreateVector(0.0, 1.0, 0.0)

	tr := ViewTransform(from, to, up)

	assert.Assert(t, IdentityMatrix().Equals(tr))
}

func TestViewTransformPositiveZ(t *testing.T) {
	from := CreatePoint(0.0, 0.0, 0.0)
	to := CreatePoint(0.0, 0.0, 1.0)
	up := CreateVector(0.0, 1.0, 0.0)
	expected := Scaling(-1.0, 1.0, -1.0)

	tr := ViewTransform(from, to, up)

	assert.Assert(t, expected.Equals(tr))
}

func TestViewTransformMoved(t *testing.T) {
	from := CreatePoint(0.0, 0.0, 8.0)
	to := CreatePoint(0.0, 0.0, 0.0)
	up := CreateVector(0.0, 1.0, 0.0)
	expected := Translation(0.0, 0.0, -8.0)

	tr := ViewTransform(from, to, up)
	fmt.Println(tr)
	assert.Assert(t, expected.Equals(tr))
}

func TestViewTransformArbitrary(t *testing.T) {
	from := CreatePoint(1.0, 3.0, 2.0)
	to := CreatePoint(4.0, -2.0, 8.0)
	up := CreateVector(1.0, 1.0, 0.0)
	data := []float64{
		-0.50709, 0.50709, 0.67612, -2.36643,
		0.76772, 0.60609, 0.12122, -2.82843,
		-0.35857, 0.59761, -0.71714, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
	expected := CreateMatrixFlat(data)

	tr := ViewTransform(from, to, up)

	assert.Assert(t, expected.Equals(tr))
}
