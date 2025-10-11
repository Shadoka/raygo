package geometry

import (
	gomath "math"
	"raygo/math"
	"reflect"

	"github.com/google/uuid"
)

type Triangle struct {
	P1               math.Point
	P2               math.Point
	P3               math.Point
	E1               math.Vector // Direction P1 -> P2
	E2               math.Vector // Direction P1 -> P3
	Normal           math.Vector
	Id               string
	Transform        math.Matrix
	Material         Material
	Parent           *Group
	CachedBounds     *Bounds
	InverseTransform *math.Matrix
	// smooth triangle fields
	N1     math.Vector
	N2     math.Vector
	N3     math.Vector
	Smooth bool
}

func CreateTriangle(p1 math.Point, p2 math.Point, p3 math.Point) *Triangle {
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)
	return &Triangle{
		P1:               p1,
		P2:               p2,
		P3:               p3,
		E1:               e1,
		E2:               e2,
		Normal:           e1.Cross(e2).Normalize(),
		Id:               uuid.NewString(),
		Transform:        math.IdentityMatrix(),
		Material:         DefaultMaterial(),
		Parent:           nil,
		CachedBounds:     nil,
		Smooth:           false,
		InverseTransform: nil,
	}
}

func CreateSmoothTriangle(p1 math.Point, p2 math.Point, p3 math.Point,
	n1 math.Vector, n2 math.Vector, n3 math.Vector) *Triangle {
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)
	return &Triangle{
		P1:               p1,
		P2:               p2,
		P3:               p3,
		E1:               e1,
		E2:               e2,
		Normal:           e1.Cross(e2).Normalize(),
		Id:               uuid.NewString(),
		Transform:        math.IdentityMatrix(),
		Material:         DefaultMaterial(),
		Parent:           nil,
		CachedBounds:     nil,
		Smooth:           true,
		N1:               n1,
		N2:               n2,
		N3:               n3,
		InverseTransform: nil,
	}
}

func (t *Triangle) Equals(other Shape) bool {
	result := reflect.TypeOf(t) == reflect.TypeOf(other) &&
		t.Transform.Equals(other.GetTransform()) &&
		t.Material.Equals(*other.GetMaterial()) &&
		t.Parent == other.GetParent()

	if result {
		otherTriangle := other.(*Triangle)
		result = result &&
			t.P1.Equals(otherTriangle.P1) &&
			t.P2.Equals(otherTriangle.P2) &&
			t.P3.Equals(otherTriangle.P3)
	}

	return result
}

func DefaultSmoothTriangle() *Triangle {
	return CreateSmoothTriangle(math.CreatePoint(0.0, 1.0, 0.0),
		math.CreatePoint(-1.0, 0.0, 0.0),
		math.CreatePoint(1.0, 0.0, 0.0),
		math.CreateVector(0.0, 1.0, 0.0),
		math.CreateVector(-1.0, 0.0, 0.0),
		math.CreateVector(1.0, 0.0, 0.0),
	)
}

func (t *Triangle) GetId() string {
	return t.Id
}

func (t *Triangle) SetTransform(m math.Matrix) {
	t.Transform = m
}

func (t *Triangle) GetTransform() math.Matrix {
	return t.Transform
}

func (t *Triangle) SetMaterial(m Material) {
	t.Material = m
}

func (t *Triangle) GetMaterial() *Material {
	return &t.Material
}

func (t *Triangle) NormalAt(p math.Point, hit Intersection) math.Vector {
	if t.Smooth {
		return t.localTriangleNormalAt(hit).Normalize()
	} else {
		return t.Normal
	}
}

func (t *Triangle) localTriangleNormalAt(hit Intersection) math.Vector {
	c1 := t.N2.Mul(hit.U)
	c2 := t.N3.Mul(hit.V)
	c3 := t.N1.Mul(1.0 - hit.U - hit.V)
	return c1.Add(c2).Add(c3)
}

func (t *Triangle) GetParent() *Group {
	return t.Parent
}

func (t *Triangle) SetParent(g *Group) {
	t.Parent = g
}

func (t *Triangle) Bounds() *Bounds {
	if t.CachedBounds != nil {
		return t.CachedBounds
	}

	minX := gomath.Min(t.P1.X, gomath.Min(t.P2.X, t.P3.X))
	minY := gomath.Min(t.P1.Y, gomath.Min(t.P2.Y, t.P3.Y))
	minZ := gomath.Min(t.P1.Z, gomath.Min(t.P2.Z, t.P3.Z))
	maxX := gomath.Max(t.P1.X, gomath.Max(t.P2.X, t.P3.X))
	maxY := gomath.Max(t.P1.Y, gomath.Max(t.P2.Y, t.P3.Y))
	maxZ := gomath.Max(t.P1.Z, gomath.Max(t.P2.Z, t.P3.Z))

	t.CachedBounds = &Bounds{
		Minimum: math.CreatePoint(minX, minY, minZ),
		Maximum: math.CreatePoint(maxX, maxY, maxZ),
	}

	return t.CachedBounds
}

func (t *Triangle) Intersect(ray Ray) []Intersection {
	localRay := ray.Transform(t.GetInverseTransform())
	return t.localIntersect(localRay)
}

func (t *Triangle) localIntersect(localRay Ray) []Intersection {
	dirCrossE2 := localRay.Direction.Cross(t.E2)
	determinant := t.E1.Dot(dirCrossE2)

	xs := make([]Intersection, 0)
	if gomath.Abs(determinant) < math.EPSILON {
		return xs
	}

	f := 1.0 / determinant
	p1ToOrigin := localRay.Origin.Subtract(t.P1)
	u := f * p1ToOrigin.Dot(dirCrossE2)

	if u < 0 || u > 1 {
		return xs
	}

	originCrossE1 := p1ToOrigin.Cross(t.E1)
	v := f * localRay.Direction.Dot(originCrossE1)

	if v < 0 || (u+v) > 1 {
		return xs
	}

	t2 := f * t.E2.Dot(originCrossE1)
	if t.Smooth {
		xs = append(xs, CreateIntersectionWithUV(t2, t, u, v))
	} else {
		xs = append(xs, CreateIntersection(t2, t))
	}
	return xs
}

func (t *Triangle) GetInverseTransform() math.Matrix {
	if t.InverseTransform != nil {
		return *t.InverseTransform
	}

	inverse := t.Transform.Inverse()
	t.InverseTransform = &inverse
	return *t.InverseTransform
}
