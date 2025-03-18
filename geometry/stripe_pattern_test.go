package geometry

import (
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateStripePattern(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)

	p := CreateStripePattern(white, black)

	assert.Assert(t, white.Equals(p.ColorA))
	assert.Assert(t, black.Equals(p.ColorB))
}

func TestColorAtConstantXVarY(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)

	p := CreateStripePattern(white, black)

	assert.Assert(t, white.Equals(p.ColorAt(math.CreatePoint(0.0, 0.0, 0.0))))
	assert.Assert(t, white.Equals(p.ColorAt(math.CreatePoint(0.0, 1.0, 0.0))))
	assert.Assert(t, white.Equals(p.ColorAt(math.CreatePoint(0.0, 2.0, 0.0))))
}

func TestColorAtConstantXVarZ(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)

	p := CreateStripePattern(white, black)

	assert.Assert(t, white.Equals(p.ColorAt(math.CreatePoint(0.0, 0.0, 0.0))))
	assert.Assert(t, white.Equals(p.ColorAt(math.CreatePoint(0.0, 0.0, 1.0))))
	assert.Assert(t, white.Equals(p.ColorAt(math.CreatePoint(0.0, 0.0, 2.0))))
}

func TestColorAtVarX(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)

	p := CreateStripePattern(white, black)

	assert.Assert(t, white.Equals(p.ColorAt(math.CreatePoint(0.0, 0.0, 0.0))))
	assert.Assert(t, white.Equals(p.ColorAt(math.CreatePoint(0.9, 0.0, 0.0))))
	assert.Assert(t, black.Equals(p.ColorAt(math.CreatePoint(1.0, 2.0, 0.0))))
	assert.Assert(t, black.Equals(p.ColorAt(math.CreatePoint(-0.1, 0.0, 0.0))))
	assert.Assert(t, black.Equals(p.ColorAt(math.CreatePoint(-1.0, 1.0, 0.0))))
	assert.Assert(t, white.Equals(p.ColorAt(math.CreatePoint(-1.1, 2.0, 0.0))))
}

func TestColorAtObjectWithObjTransform(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)

	p := CreateStripePattern(white, black)

	s := CreateSphere()
	s.SetTransform(math.Scaling(2.0, 2.0, 2.0))

	assert.Assert(t, white.Equals(p.ColorAtObject(math.CreatePoint(1.5, 0.0, 0.0), s)))
}

func TestColorAtObjectWithPatternTransform(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)

	p := CreateStripePattern(white, black)
	p.SetTransform(math.Scaling(2.0, 2.0, 2.0))

	s := CreateSphere()

	assert.Assert(t, white.Equals(p.ColorAtObject(math.CreatePoint(1.5, 0.0, 0.0), s)))
}

func TestColorAtObjectWithPatternAndObjTransform(t *testing.T) {
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)

	p := CreateStripePattern(white, black)
	p.SetTransform(math.Translation(0.5, 0.0, 0.0))

	s := CreateSphere()
	s.SetTransform(math.Scaling(2.0, 2.0, 2.0))

	assert.Assert(t, white.Equals(p.ColorAtObject(math.CreatePoint(2.5, 0.0, 0.0), s)))
}
