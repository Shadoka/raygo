package lighting

import (
	gomath "math"
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestLightingEyeBetweenLightAndSurface(t *testing.T) {
	eyev := math.CreateVector(0.0, 0.0, -1.0)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 0.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(1.9, 1.9, 1.9)

	actual := PhongLighting(m, light, p, eyev, normalv, false)

	assert.Assert(t, expected.Equals(actual))
}

func TestLightingEyeOffsetBetweenLightAndSurface(t *testing.T) {
	eyev := math.CreateVector(0.0, gomath.Sqrt(2)/2, -gomath.Sqrt(2)/2)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 0.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(1.0, 1.0, 1.0)

	assert.Assert(t, expected.Equals(PhongLighting(m, light, p, eyev, normalv, false)))
}

func TestLightingEyeBetweenLightOffsetAndSurface(t *testing.T) {
	eyev := math.CreateVector(0.0, 0.0, -1.0)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 10.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(0.7364, 0.7364, 0.7364)

	assert.Assert(t, expected.Equals(PhongLighting(m, light, p, eyev, normalv, false)))
}

func TestLightingEyeInReflectionVector(t *testing.T) {
	eyev := math.CreateVector(0.0, -gomath.Sqrt(2)/2, -gomath.Sqrt(2)/2)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 10.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(1.6364, 1.6364, 1.6364)

	assert.Assert(t, expected.Equals(PhongLighting(m, light, p, eyev, normalv, false)))
}

func TestLightingLightBehindSurface(t *testing.T) {
	eyev := math.CreateVector(0.0, 0.0, -1.0)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 10.0, 10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(0.1, 0.1, 0.1)

	assert.Assert(t, expected.Equals(PhongLighting(m, light, p, eyev, normalv, false)))
}

func TestLightingEyeBetweenLightAndSurfaceShadow(t *testing.T) {
	eyev := math.CreateVector(0.0, 0.0, -1.0)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 0.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(0.1, 0.1, 0.1)

	actual := PhongLighting(m, light, p, eyev, normalv, true)

	assert.Assert(t, expected.Equals(actual))
}
