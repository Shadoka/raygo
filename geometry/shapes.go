package geometry

import (
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
	NormalAt(p math.Point, hit Intersection) math.Vector
	GetParent() *Group
	SetParent(g *Group)
	Bounds() *Bounds
	GetInverseTransform() math.Matrix
	CalculateInverseTransform()
	GetUvCoordinate(direction math.Vector) (float64, float64)
}

func GetCenter(s Shape) math.Point {
	tfBounds := s.Bounds().ApplyTransform(s.GetTransform())
	return tfBounds.Maximum.Add(tfBounds.Minimum).Mul(0.5)
}

func WorldToObject(s Shape, p math.Point) math.Point {
	if s.GetParent() != nil {
		p = WorldToObject(s.GetParent(), p)
	}
	return s.GetInverseTransform().MulT(p)
}

func NormalToWorld(s Shape, normal math.Vector) math.Vector {
	normal = s.GetInverseTransform().Transpose().MulT(normal)
	normal.W = 0
	normal = normal.Normalize()

	if s.GetParent() != nil {
		normal = NormalToWorld(s.GetParent(), normal)
	}
	return normal
}
