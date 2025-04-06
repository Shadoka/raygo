package geometry

import (
	gomath "math"
	"raygo/math"
)

type Shape interface {
	Equals(other Shape) bool
	GetId() string
	SetTransform(m math.Matrix)
	GetTransform() math.Matrix
	SetMaterial(m Material)
	GetMaterial() *Material
	Intersect(ray Ray) []Intersection
	NormalAt(p math.Point) math.Vector
	GetParent() *Group
	SetParent(g *Group)
	Bounds() *Bounds
}

type Bounds struct {
	Minimum math.Point
	Maximum math.Point
}

func WorldToObject(s Shape, p math.Point) math.Point {
	if s.GetParent() != nil {
		p = WorldToObject(s.GetParent(), p)
	}
	return s.GetTransform().Inverse().MulT(p)
}

func NormalToWorld(s Shape, normal math.Vector) math.Vector {
	normal = s.GetTransform().Inverse().Transpose().MulT(normal)
	normal.W = 0
	normal = normal.Normalize()

	if s.GetParent() != nil {
		normal = NormalToWorld(s.GetParent(), normal)
	}
	return normal
}

func BoundsToObjectSpace(b Bounds, tf math.Matrix) *Bounds {
	tmpBB := Bounds{
		Minimum: tf.MulT(b.Minimum),
		Maximum: tf.MulT(b.Maximum),
	}
	corners := CalculateBBCorners(tmpBB)
	return FindBoundingBox(corners)
}

func CalculateBBCorners(b Bounds) []math.Point {
	corners := make([]math.Point, 0)
	corners = append(corners, b.Minimum, b.Maximum)

	p1 := math.CreatePoint(b.Minimum.X, b.Minimum.Y, b.Maximum.Z)
	p2 := math.CreatePoint(b.Minimum.X, b.Maximum.Y, b.Minimum.Z)
	p3 := math.CreatePoint(b.Minimum.X, b.Maximum.Y, b.Maximum.Z)
	p4 := math.CreatePoint(b.Maximum.X, b.Minimum.Y, b.Minimum.Z)
	p5 := math.CreatePoint(b.Maximum.X, b.Maximum.Y, b.Minimum.Z)
	p6 := math.CreatePoint(b.Maximum.X, b.Minimum.Y, b.Maximum.Z)
	corners = append(corners, p1, p2, p3, p4, p5, p6)

	return corners
}

func FindBoundingBox(corners []math.Point) *Bounds {
	minX, minY, minZ := gomath.MaxFloat64, gomath.MaxFloat64, gomath.MaxFloat64
	maxX, maxY, maxZ := gomath.SmallestNonzeroFloat64, gomath.SmallestNonzeroFloat64, gomath.SmallestNonzeroFloat64

	for _, p := range corners {
		minX = gomath.Min(minX, p.X)
		minY = gomath.Min(minY, p.Y)
		minZ = gomath.Min(minZ, p.Z)
		maxX = gomath.Max(maxX, p.X)
		maxY = gomath.Max(maxY, p.Y)
		maxZ = gomath.Max(maxZ, p.Z)
	}

	return &Bounds{
		Minimum: math.CreatePoint(minX, minY, minZ),
		Maximum: math.CreatePoint(maxX, maxY, maxZ),
	}
}

func (b *Bounds) Equals(other *Bounds) bool {
	return b.Minimum.Equals(other.Minimum) &&
		b.Maximum.Equals(other.Maximum)
}
