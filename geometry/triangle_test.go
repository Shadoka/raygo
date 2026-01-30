package geometry

import (
	"fmt"
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

var p = math.CreatePoint
var v = math.CreateVector

func TestCreateTriangle(t *testing.T) {
	p1 := math.CreatePoint(0.0, 1.0, 0.0)
	p2 := math.CreatePoint(-1.0, 0.0, 0.0)
	p3 := math.CreatePoint(1.0, 0.0, 0.0)
	e1 := math.CreateVector(-1.0, -1.0, 0.0)
	e2 := math.CreateVector(1.0, -1.0, 0.0)
	normal := math.CreateVector(0.0, 0.0, 1.0)

	tri := CreateTriangle(p1, p2, p3)

	assert.Assert(t, p1.Equals(tri.P1))
	assert.Assert(t, p2.Equals(tri.P2))
	assert.Assert(t, p3.Equals(tri.P3))
	assert.Assert(t, e1.Equals(tri.E1))
	assert.Assert(t, e2.Equals(tri.E2))
	assert.Assert(t, normal.Equals(tri.Normal))
}

func TestNormalAt(t *testing.T) {
	p1 := p(0.0, 0.5, 0.0)
	p2 := p(-0.5, 0.75, 0.0)
	p3 := p(0.5, 0.25, 0.0)

	tri := CreateTriangle(p(0.0, 1.0, 0.0), p(-1.0, 0.0, 0.0), p(1.0, 0.0, 0.0))

	assert.Assert(t, tri.Normal.Equals(tri.NormalAt(p1, Intersection{})))
	assert.Assert(t, tri.Normal.Equals(tri.NormalAt(p2, Intersection{})))
	assert.Assert(t, tri.Normal.Equals(tri.NormalAt(p3, Intersection{})))
}

func TestLocalIntersectParallelRay(t *testing.T) {
	tri := CreateTriangle(p(0.0, 1.0, 0.0), p(-1.0, 0.0, 0.0), p(1.0, 0.0, 0.0))
	r := CreateRay(p(0.0, -1.0, -2.0), v(0.0, 1.0, 0.0))

	xs := tri.localIntersect(r)

	assert.Assert(t, len(xs) == 0)
}

func TestLocalIntersectP1P3(t *testing.T) {
	tri := CreateTriangle(p(0.0, 1.0, 0.0), p(-1.0, 0.0, 0.0), p(1.0, 0.0, 0.0))
	r := CreateRay(p(1.0, 1.0, -2.0), v(0.0, 0.0, 1.0))

	xs := tri.localIntersect(r)

	assert.Assert(t, len(xs) == 0)
}

func TestLocalIntersectP1P2(t *testing.T) {
	tri := CreateTriangle(p(0.0, 1.0, 0.0), p(-1.0, 0.0, 0.0), p(1.0, 0.0, 0.0))
	r := CreateRay(p(-1.0, 1.0, -2.0), v(0.0, 0.0, 1.0))

	xs := tri.localIntersect(r)

	assert.Assert(t, len(xs) == 0)
}

func TestLocalIntersectP2P3(t *testing.T) {
	tri := CreateTriangle(p(0.0, 1.0, 0.0), p(-1.0, 0.0, 0.0), p(1.0, 0.0, 0.0))
	r := CreateRay(p(0.0, -1.0, -2.0), v(0.0, 0.0, 1.0))

	xs := tri.localIntersect(r)

	assert.Assert(t, len(xs) == 0)
}

func TestLocalIntersectHit(t *testing.T) {
	tri := CreateTriangle(p(0.0, 1.0, 0.0), p(-1.0, 0.0, 0.0), p(1.0, 0.0, 0.0))
	r := CreateRay(p(0.0, 0.5, -2.0), v(0.0, 0.0, 1.0))

	xs := tri.localIntersect(r)

	assert.Assert(t, len(xs) == 1)
	assert.Assert(t, floatEquals(xs[0].IntersectionAt, 2.0))
}

func TestSmoothTriangleIntersectionUV(t *testing.T) {
	r := CreateRay(p(-0.2, 0.3, -2.0), math.CreateVector(0.0, 0.0, 1.0))
	xs := DefaultSmoothTriangle().localIntersect(r)

	assert.Assert(t, len(xs) == 1)
	assert.Assert(t, floatEquals(xs[0].U, 0.45))
	assert.Assert(t, floatEquals(xs[0].V, 0.25))
}

func TestSmoothTriangleInterpolatedNormal(t *testing.T) {
	tri := DefaultSmoothTriangle()
	i := CreateIntersectionWithUV(1.0, tri, 0.45, 0.25)
	n := tri.NormalAt(p(0.0, 0.0, 0.0), i)
	fmt.Println(n)
	assert.Assert(t, n.Equals(v(-0.5547, 0.83205, 0.0)))
}

func TestRandomEmptyInverseBug(t *testing.T) {
	p1 := math.CreatePoint(2.2383, 6.6632, 12.4922)
	p2 := math.CreatePoint(3.248, 6.2428, 12.4922)
	p3 := math.CreatePoint(3.2866, 6.317, 12.4725)
	n1 := math.CreateVector(-0.0059, -0.0178, 0.9998)
	n2 := math.CreateVector(-0.0087, -0.0167, 0.9998)
	n3 := math.CreateVector(0.1764, 0.3448, 0.922)
	tri := CreateSmoothTriangle(p1, p2, p3, n1, n2, n3)
	tri.CalculateInverseTransform()

	b := tri.Bounds()
	assert.Assert(t, b != nil)

	inv := tri.GetInverseTransform()
	assert.Assert(t, inv.Equals(math.IdentityMatrix()))

	inv2 := tri.GetInverseTransform()
	assert.Assert(t, inv2.Equals(math.IdentityMatrix()))
}
