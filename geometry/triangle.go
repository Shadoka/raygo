package geometry

import (
	gomath "math"
	"raygo/math"
	"reflect"

	"github.com/google/uuid"
)

type Triangle struct {
	P1           math.Point
	P2           math.Point
	P3           math.Point
	E1           math.Vector // Direction P1 -> P2
	E2           math.Vector // Direction P1 -> P3
	Normal       math.Vector
	Id           string
	Transform    math.Matrix
	Material     Material
	Parent       *Group
	CachedBounds *Bounds
}

func CreateTriangle(p1 math.Point, p2 math.Point, p3 math.Point) *Triangle {
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)
	return &Triangle{
		P1:           p1,
		P2:           p2,
		P3:           p3,
		E1:           e1,
		E2:           e2,
		Normal:       e1.Cross(e2).Normalize(),
		Id:           uuid.NewString(),
		Transform:    math.IdentityMatrix(),
		Material:     DefaultMaterial(),
		Parent:       nil,
		CachedBounds: nil,
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

func (t *Triangle) NormalAt(p math.Point) math.Vector {
	return t.Normal
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
	localRay := ray.Transform(t.Transform.Inverse())
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
	return append(xs, CreateIntersection(t2, t))
}
