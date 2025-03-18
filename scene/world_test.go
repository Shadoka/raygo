package scene

import (
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEmptyWorld(t *testing.T) {
	w := EmptyWorld()

	assert.Assert(t, w.Light == nil)
	assert.Assert(t, len(w.Objects) == 0)
}

func TestDefaultWorld(t *testing.T) {
	expectedLight := lighting.CreateLight(math.CreatePoint(-10.0, 10.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	expectedShape1 := g.CreateSphere()
	m1 := g.DefaultMaterial()
	m1.SetColor(math.CreateColor(0.8, 1.0, 0.6))
	(&m1).Diffuse = 0.7
	(&m1).Specular = 0.2
	expectedShape1.SetMaterial(m1)
	expectedShape2 := g.CreateSphere()
	transform := math.Scaling(0.5, 0.5, 0.5)
	expectedShape2.SetTransform(transform)

	w := DefaultWorld()

	assert.Assert(t, expectedLight.Equals(*w.Light))
	assert.Assert(t, len(w.Objects) == 2)
	assert.Assert(t, expectedShape1.Equals(w.Objects[0]))
	assert.Assert(t, expectedShape2.Equals(w.Objects[1]))
}

func TestIntersectWorld(t *testing.T) {
	w := DefaultWorld()
	r := g.CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))

	xs := w.Intersect(r)

	assert.Assert(t, len(xs) == 4)
	assert.Assert(t, xs[0].IntersectionAt == 4.0)
	assert.Assert(t, xs[1].IntersectionAt == 4.5)
	assert.Assert(t, xs[2].IntersectionAt == 5.5)
	assert.Assert(t, xs[3].IntersectionAt == 6.0)
}

func TestShadeHitDefault(t *testing.T) {
	w := DefaultWorld()
	r := g.CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))
	i := g.CreateIntersection(4.0, w.Objects[0])
	expected := math.CreateColor(0.38066, 0.47583, 0.2855)

	comps := i.PrepareComputation(r)
	actual := w.ShadeHit(comps)

	assert.Assert(t, expected.Equals(actual))
}

func TestShadeHitInside(t *testing.T) {
	w := DefaultWorld()
	l := lighting.CreateLight(math.CreatePoint(0.0, 0.25, 0.0), math.CreateColor(1.0, 1.0, 1.0))
	w.Light = &l

	r := g.CreateRay(math.CreatePoint(0.0, 0.0, 0.0), math.CreateVector(0.0, 0.0, 1.0))
	i := g.CreateIntersection(0.5, w.Objects[1])
	expected := math.CreateColor(0.1, 0.1, 0.1)

	comps := i.PrepareComputation(r)
	actual := w.ShadeHit(comps)

	assert.Assert(t, expected.Equals(actual))
}

func TestColorAtRayMiss(t *testing.T) {
	w := DefaultWorld()
	r := g.CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 1.0, 0.0))
	expected := math.CreateColor(0.0, 0.0, 0.0)

	actual := w.ColorAt(r)

	assert.Assert(t, expected.Equals(actual))
}

func TestColorAtRayHit(t *testing.T) {
	w := DefaultWorld()
	r := g.CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))
	expected := math.CreateColor(0.38066, 0.47583, 0.2855)

	actual := w.ColorAt(r)

	assert.Assert(t, expected.Equals(actual))
}

func TestColorAtHitBehindRay(t *testing.T) {
	w := EmptyWorld()
	objects := make([]g.Shape, 0)

	s1 := g.CreateSphere()
	m1 := g.DefaultMaterial()
	m1.SetColor(math.CreateColor(0.8, 1.0, 0.6))
	(&m1).Ambient = 1.0
	(&m1).Diffuse = 0.7
	(&m1).Specular = 0.2
	s1.SetMaterial(m1)
	objects = append(objects, s1)

	s2 := g.CreateSphere()
	transform := math.Scaling(0.5, 0.5, 0.5)
	s2.SetTransform(transform)
	m2 := g.DefaultMaterial()
	(&m2).Ambient = 1.0
	s2.SetMaterial(m2)
	objects = append(objects, s2)

	w.Objects = objects

	light := lighting.CreateLight(math.CreatePoint(-10.0, 10.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	w.Light = &light

	r := g.CreateRay(math.CreatePoint(0.0, 0.0, 0.75), math.CreateVector(0.0, 0.0, -1.0))

	actual := w.ColorAt(r)

	assert.Assert(t, m2.Color.Equals(actual))
}

func TestIsShadowedColinearPoint(t *testing.T) {
	w := DefaultWorld()
	p := math.CreatePoint(0.0, 10.0, 0.0)

	assert.Assert(t, w.IsShadowed(p) == false)
}

func TestIsShadowedBehindSphere(t *testing.T) {
	w := DefaultWorld()
	p := math.CreatePoint(10.0, -10.0, 10.0)

	assert.Assert(t, w.IsShadowed(p) == true)
}

func TestIsShadowedBehindLight(t *testing.T) {
	w := DefaultWorld()
	p := math.CreatePoint(-20.0, 20.0, -20.0)

	assert.Assert(t, w.IsShadowed(p) == false)
}

func TestIsShadowedBetweenLightAndShape(t *testing.T) {
	w := DefaultWorld()
	p := math.CreatePoint(-2.0, 2.0, -2.0)

	assert.Assert(t, w.IsShadowed(p) == false)
}

func TestShadeHitInShadow(t *testing.T) {
	w := EmptyWorld()
	expected := math.CreateColor(0.1, 0.1, 0.1)

	objs := make([]g.Shape, 0)
	s1 := g.CreateSphere()
	s2 := g.CreateSphere()
	s2.SetTransform(math.Translation(0.0, 0.0, 10.0))
	objs = append(objs, s1, s2)
	w.Objects = objs

	light := lighting.CreateLight(math.CreatePoint(0.0, 0.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	w.Light = &light

	r := g.CreateRay(math.CreatePoint(0.0, 0.0, 5.0), math.CreateVector(0.0, 0.0, 1.0))
	i := g.CreateIntersection(4.0, s2)

	comps := i.PrepareComputation(r)
	c := w.ShadeHit(comps)

	assert.Assert(t, expected.Equals(c))
}
