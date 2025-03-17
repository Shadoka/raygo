package lighting

import (
	gomath "math"
	"raygo/math"
	"raygo/ray"
	"testing"

	"gotest.tools/v3/assert"
)

func TestLightingEyeBetweenLightAndSurface(t *testing.T) {
	s := ray.CreateSphere()
	eyev := math.CreateVector(0.0, 0.0, -1.0)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := ray.DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 0.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(1.9, 1.9, 1.9)

	actual := PhongLighting(m, s, light, p, eyev, normalv, false)

	assert.Assert(t, expected.Equals(actual))
}

func TestLightingEyeOffsetBetweenLightAndSurface(t *testing.T) {
	s := ray.CreateSphere()
	eyev := math.CreateVector(0.0, gomath.Sqrt(2)/2, -gomath.Sqrt(2)/2)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := ray.DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 0.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(1.0, 1.0, 1.0)

	assert.Assert(t, expected.Equals(PhongLighting(m, s, light, p, eyev, normalv, false)))
}

func TestLightingEyeBetweenLightOffsetAndSurface(t *testing.T) {
	s := ray.CreateSphere()
	eyev := math.CreateVector(0.0, 0.0, -1.0)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := ray.DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 10.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(0.7364, 0.7364, 0.7364)

	assert.Assert(t, expected.Equals(PhongLighting(m, s, light, p, eyev, normalv, false)))
}

func TestLightingEyeInReflectionVector(t *testing.T) {
	s := ray.CreateSphere()
	eyev := math.CreateVector(0.0, -gomath.Sqrt(2)/2, -gomath.Sqrt(2)/2)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := ray.DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 10.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(1.6364, 1.6364, 1.6364)

	assert.Assert(t, expected.Equals(PhongLighting(m, s, light, p, eyev, normalv, false)))
}

func TestLightingLightBehindSurface(t *testing.T) {
	s := ray.CreateSphere()
	eyev := math.CreateVector(0.0, 0.0, -1.0)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := ray.DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 10.0, 10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(0.1, 0.1, 0.1)

	assert.Assert(t, expected.Equals(PhongLighting(m, s, light, p, eyev, normalv, false)))
}

func TestLightingEyeBetweenLightAndSurfaceShadow(t *testing.T) {
	s := ray.CreateSphere()
	eyev := math.CreateVector(0.0, 0.0, -1.0)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	m := ray.DefaultMaterial()
	p := math.CreatePoint(0.0, 0.0, 0.0)
	light := CreateLight(math.CreatePoint(0.0, 0.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	expected := math.CreateColor(0.1, 0.1, 0.1)

	actual := PhongLighting(m, s, light, p, eyev, normalv, true)

	assert.Assert(t, expected.Equals(actual))
}

func TestMaterialWithPattern(t *testing.T) {
	s := ray.CreateSphere()
	white := math.CreateColor(1.0, 1.0, 1.0)
	black := math.CreateColor(0.0, 0.0, 0.0)
	p := ray.CreateStripePattern(white, black)
	m := ray.DefaultMaterial()
	m.SetPattern(p)
	m.SetAmbient(1.0)
	m.SetDiffuse(0.0)
	m.SetSpecular(0.0)

	eyev := math.CreateVector(0.0, 0.0, -1.0)
	normalv := math.CreateVector(0.0, 0.0, -1.0)
	light := CreateLight(math.CreatePoint(0.0, 0.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))

	c1 := PhongLighting(m, s, light, math.CreatePoint(0.9, 0.0, 0.0), eyev, normalv, false)
	c2 := PhongLighting(m, s, light, math.CreatePoint(1.1, 0.0, 0.0), eyev, normalv, false)

	assert.Assert(t, white.Equals(c1))
	assert.Assert(t, black.Equals(c2))
}
