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
}
