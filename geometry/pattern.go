package geometry

import (
	"raygo/math"
)

type Pattern interface {
	ColorAt(point math.Point) math.Color
	ColorAtObject(point math.Point, obj Shape) math.Color
	SetTransform(tf math.Matrix)
	GetTransform() math.Matrix
	Equals(other Pattern) bool
	CalculateInverseTransform()
}
