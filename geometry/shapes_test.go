package geometry

import (
	gomath "math"
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateSphere(t *testing.T) {
	s := CreateSphere()

	assert.Assert(t, s.Transform.Equals(math.IdentityMatrix()))
}

func TestSetTransform(t *testing.T) {
	s := CreateSphere()
	tf := math.Translation(2.0, 3.0, 4.0)

	s.SetTransform(tf)

	assert.Assert(t, s.Transform.Equals(tf))
}

func TestIntersectSphereTwoPoints(t *testing.T) {
	r := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))
	s := CreateSphere()

	xs := s.Intersect(r)

	assert.Assert(t, len(xs) == 2)
	assert.Assert(t, xs[0].IntersectionAt == 4.0)
	assert.Assert(t, xs[1].IntersectionAt == 6.0)
}

func TestIntersectSphereTangent(t *testing.T) {
	r := CreateRay(math.CreatePoint(0.0, 1.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))
	s := CreateSphere()

	xs := s.Intersect(r)

	assert.Assert(t, len(xs) == 2)
	assert.Assert(t, xs[0].IntersectionAt == 5.0)
	assert.Assert(t, xs[1].IntersectionAt == 5.0)
}

func TestIntersectSphereMiss(t *testing.T) {
	r := CreateRay(math.CreatePoint(0.0, 2.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))
	s := CreateSphere()

	xs := s.Intersect(r)

	assert.Assert(t, len(xs) == 0)
}

func TestIntersectInsideSphere(t *testing.T) {
	r := CreateRay(math.CreatePoint(0.0, 0.0, 0.0), math.CreateVector(0.0, 0.0, 1.0))
	s := CreateSphere()

	xs := s.Intersect(r)

	assert.Assert(t, len(xs) == 2)
	assert.Assert(t, xs[0].IntersectionAt == -1.0)
	assert.Assert(t, xs[1].IntersectionAt == 1.0)
}

func TestIntersectBehindSphere(t *testing.T) {
	r := CreateRay(math.CreatePoint(0.0, 0.0, 2.0), math.CreateVector(0.0, 0.0, 1.0))
	s := CreateSphere()

	xs := s.Intersect(r)

	assert.Assert(t, len(xs) == 2)
	assert.Assert(t, xs[0].IntersectionAt == -3.0)
	assert.Assert(t, xs[1].IntersectionAt == -1.0)
}

func TestIntersectSetsObject(t *testing.T) {
	r := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))
	s := CreateSphere()

	xs := s.Intersect(r)

	assert.Assert(t, len(xs) == 2)
	assert.Assert(t, xs[0].Object.Equals(s))
	assert.Assert(t, xs[1].Object.Equals(s))
}

func TestIntersectSphereWithScaling(t *testing.T) {
	r := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))
	s := CreateSphere()
	s.SetTransform(math.Scaling(2.0, 2.0, 2.0))

	xs := s.Intersect(r)

	assert.Assert(t, len(xs) == 2)
	assert.Assert(t, xs[0].IntersectionAt == 3.0)
	assert.Assert(t, xs[1].IntersectionAt == 7.0)
}

func TestIntersectSphereWithTranslation(t *testing.T) {
	r := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))
	s := CreateSphere()
	s.SetTransform(math.Translation(5.0, 0.0, 0.0))

	xs := s.Intersect(r)

	assert.Assert(t, len(xs) == 0)
}

func TestNormalAtSphereX(t *testing.T) {
	s := CreateSphere()
	expected := math.CreateVector(1.0, 0.0, 0.0)

	assert.Assert(t, expected.Equals(s.NormalAt(math.CreatePoint(1.0, 0.0, 0.0))))
}

func TestNormalAtSphereY(t *testing.T) {
	s := CreateSphere()
	expected := math.CreateVector(0.0, 1.0, 0.0)

	assert.Assert(t, expected.Equals(s.NormalAt(math.CreatePoint(0.0, 1.0, 0.0))))
}

func TestNormalAtSphereZ(t *testing.T) {
	s := CreateSphere()
	expected := math.CreateVector(0.0, 0.0, 1.0)

	assert.Assert(t, expected.Equals(s.NormalAt(math.CreatePoint(0.0, 0.0, 1.0))))
}

func TestNormalAtSphereNonaxialPoint(t *testing.T) {
	s := CreateSphere()
	expected := math.CreateVector(gomath.Sqrt(3.0)/3.0, gomath.Sqrt(3.0)/3.0, gomath.Sqrt(3.0)/3.0)

	assert.Assert(t, expected.Equals(s.NormalAt(math.CreatePoint(gomath.Sqrt(3.0)/3.0, gomath.Sqrt(3.0)/3.0, gomath.Sqrt(3.0)/3.0))))
}

func TestNormalAtSphereIsNormalized(t *testing.T) {
	s := CreateSphere()
	expected := math.CreateVector(gomath.Sqrt(3.0)/3.0, gomath.Sqrt(3.0)/3.0, gomath.Sqrt(3.0)/3.0).Normalize()

	assert.Assert(t, expected.Equals(s.NormalAt(math.CreatePoint(gomath.Sqrt(3.0)/3.0, gomath.Sqrt(3.0)/3.0, gomath.Sqrt(3.0)/3.0))))
}

func TestNormalAtSphereTranslated(t *testing.T) {
	s := CreateSphere()
	s.SetTransform(math.Translation(0.0, 1.0, 0.0))
	expected := math.CreateVector(0.0, 0.70711, -0.70711)

	n := s.NormalAt(math.CreatePoint(0.0, 1.70711, -0.70711))
	assert.Assert(t, expected.Equals(n))
}

func TestNormalAtSphereTransformed(t *testing.T) {
	s := CreateSphere()
	s.SetTransform(math.Scaling(1.0, 0.5, 1.0).MulM(math.Rotation_Z(gomath.Pi / 5.0)))
	expected := math.CreateVector(0.0, 0.97014, -0.24254)

	n := s.NormalAt(math.CreatePoint(0.0, gomath.Sqrt(2)/2.0, -gomath.Sqrt(2)/2))
	assert.Assert(t, expected.Equals(n))
}

func TestDefaultMaterialSphere(t *testing.T) {
	s := CreateSphere()
	m := DefaultMaterial()

	assert.Assert(t, m.Equals(*s.GetMaterial()))
}

func TestSetMaterialSphere(t *testing.T) {
	s := CreateSphere()
	m := DefaultMaterial()
	m.Ambient = 1.0

	s.SetMaterial(m)

	assert.Assert(t, m.Equals(*s.GetMaterial()))
}

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

	actual := p.NormalAt(point)

	assert.Assert(t, expected.Equals(actual))
}

func TestNormalAtDefaultPlaneTranslated(t *testing.T) {
	p := CreatePlane()
	p.SetTransform(math.Translation(4.0, 0.0, 0.0))
	point := math.CreatePoint(0.0, 0.0, 5.0)
	expected := math.CreateVector(0.0, 1.0, 0.0)

	actual := p.NormalAt(point)

	assert.Assert(t, expected.Equals(actual))
}

func TestNormalAtDefaultPlaneRotatedX(t *testing.T) {
	p := CreatePlane()
	p.SetTransform(math.Rotation_X(gomath.Pi / 2.0))
	point := math.CreatePoint(0.0, 0.0, 5.0)
	expected := math.CreateVector(0.0, 0.0, 1.0)

	actual := p.NormalAt(point)

	assert.Assert(t, expected.Equals(actual))
}

func TestCreateGlassSphere(t *testing.T) {
	s := CreateGlassSphere()

	assert.Assert(t, math.IdentityMatrix().Equals(s.GetTransform()))
	assert.Assert(t, s.GetMaterial().Transparency == 1.0)
	assert.Assert(t, s.GetMaterial().RefractiveIndex == 1.5)
}

func TestCubeIntersectXFaces(t *testing.T) {
	c := CreateCube()
	rFromRight := CreateRay(math.CreatePoint(5.0, 0.5, 0.0), math.CreateVector(-1.0, 0.0, 0.0))
	rFromLeft := CreateRay(math.CreatePoint(-5.0, 0.5, 0.0), math.CreateVector(1.0, 0.0, 0.0))

	xsFromRight := c.localCubeIntersect(rFromRight)
	xsFromLeft := c.localCubeIntersect(rFromLeft)

	assert.Assert(t, xsFromRight[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromRight[1].IntersectionAt == 6.0)
	assert.Assert(t, xsFromLeft[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromLeft[1].IntersectionAt == 6.0)
}

func TestCubeIntersectYFaces(t *testing.T) {
	c := CreateCube()
	rFromTop := CreateRay(math.CreatePoint(0.5, 5.0, 0.0), math.CreateVector(0.0, -1.0, 0.0))
	rFromBottom := CreateRay(math.CreatePoint(0.5, -5.0, 0.0), math.CreateVector(0.0, 1.0, 0.0))

	xsFromTop := c.localCubeIntersect(rFromTop)
	xsFromBottom := c.localCubeIntersect(rFromBottom)

	assert.Assert(t, xsFromTop[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromTop[1].IntersectionAt == 6.0)
	assert.Assert(t, xsFromBottom[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromBottom[1].IntersectionAt == 6.0)
}

func TestCubeIntersectZFaces(t *testing.T) {
	c := CreateCube()
	rFromBehind := CreateRay(math.CreatePoint(0.5, 0.0, 5.0), math.CreateVector(0.0, 0.0, -1.0))
	rFromFront := CreateRay(math.CreatePoint(0.5, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))

	xsFromBehind := c.localCubeIntersect(rFromBehind)
	xsFromFront := c.localCubeIntersect(rFromFront)

	assert.Assert(t, xsFromBehind[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromBehind[1].IntersectionAt == 6.0)
	assert.Assert(t, xsFromFront[0].IntersectionAt == 4.0)
	assert.Assert(t, xsFromFront[1].IntersectionAt == 6.0)
}

func TestCubeIntersectFromInside(t *testing.T) {
	c := CreateCube()
	r := CreateRay(math.CreatePoint(0.0, 0.5, 0.0), math.CreateVector(0.0, 0.0, 1.0))

	xsFromBehind := c.localCubeIntersect(r)

	assert.Assert(t, xsFromBehind[0].IntersectionAt == -1.0)
	assert.Assert(t, xsFromBehind[1].IntersectionAt == 1.0)
}

func TestCubeIntersectRayMiss(t *testing.T) {
	c := CreateCube()
	r1 := CreateRay(math.CreatePoint(-2.0, 0.0, 0.0), math.CreateVector(0.2673, 0.5345, 0.8018))
	r2 := CreateRay(math.CreatePoint(0.0, -2.0, 0.0), math.CreateVector(0.8018, 0.2673, 0.5345))
	r3 := CreateRay(math.CreatePoint(0.0, 0.0, -2.0), math.CreateVector(0.5345, 0.8018, 0.2673))
	r4 := CreateRay(math.CreatePoint(2.0, 0.0, 2.0), math.CreateVector(0.0, 0.0, -1.0))
	r5 := CreateRay(math.CreatePoint(0.0, 2.0, 2.0), math.CreateVector(0.0, -1.0, 0.0))
	r6 := CreateRay(math.CreatePoint(2.0, 2.0, 0.0), math.CreateVector(-1.0, 0.0, 0.0))

	xs := c.localCubeIntersect(r1)
	xs = append(xs, c.localCubeIntersect(r2)...)
	xs = append(xs, c.localCubeIntersect(r3)...)
	xs = append(xs, c.localCubeIntersect(r4)...)
	xs = append(xs, c.localCubeIntersect(r5)...)
	xs = append(xs, c.localCubeIntersect(r6)...)

	assert.Assert(t, len(xs) == 0)
}

func TestCubeLocalNormalAt(t *testing.T) {
	c := CreateCube()
	p1 := math.CreatePoint(1.0, 0.5, -0.8)
	p2 := math.CreatePoint(-1.0, -0.2, 0.9)
	p3 := math.CreatePoint(-0.4, 1.0, -0.1)
	p4 := math.CreatePoint(0.3, -1.0, -0.7)
	p5 := math.CreatePoint(-0.6, 0.3, 1.0)
	p6 := math.CreatePoint(0.4, 0.4, -1.0)
	p7 := math.CreatePoint(1.0, 1.0, 1.0)
	p8 := math.CreatePoint(-1.0, -1.0, -1.0)
	expected1 := math.CreateVector(1.0, 0.0, 0.0)
	expected2 := math.CreateVector(-1.0, 0.0, 0.0)
	expected3 := math.CreateVector(0.0, 1.0, 0.0)
	expected4 := math.CreateVector(0.0, -1.0, 0.0)
	expected5 := math.CreateVector(0.0, 0.0, 1.0)
	expected6 := math.CreateVector(0.0, 0.0, -1.0)
	expected7 := math.CreateVector(1.0, 0.0, 0.0)
	expected8 := math.CreateVector(-1.0, 0.0, 0.0)

	assert.Assert(t, expected1.Equals(c.localCubeNormalAt(p1)))
	assert.Assert(t, expected2.Equals(c.localCubeNormalAt(p2)))
	assert.Assert(t, expected3.Equals(c.localCubeNormalAt(p3)))
	assert.Assert(t, expected4.Equals(c.localCubeNormalAt(p4)))
	assert.Assert(t, expected5.Equals(c.localCubeNormalAt(p5)))
	assert.Assert(t, expected6.Equals(c.localCubeNormalAt(p6)))
	assert.Assert(t, expected7.Equals(c.localCubeNormalAt(p7)))
	assert.Assert(t, expected8.Equals(c.localCubeNormalAt(p8)))
}
