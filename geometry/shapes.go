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
	NormalAt(p math.Point) math.Vector
	GetParent() *Group
	SetParent(g *Group)
	Bounds() *Bounds
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
