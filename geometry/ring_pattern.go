package geometry

import (
	gomath "math"
	"raygo/math"
	"reflect"
)

type RingPattern struct {
	ColorA           math.Color
	ColorB           math.Color
	Transform        math.Matrix
	InverseTransform math.Matrix
}

func CreateRingPattern(a math.Color, b math.Color) *RingPattern {
	return &RingPattern{
		ColorA:    a,
		ColorB:    b,
		Transform: math.IdentityMatrix(),
	}
}

func (rp *RingPattern) ColorAt(point math.Point) math.Color {
	val := gomath.Sqrt(point.X*point.X + point.Z*point.Z)
	if int(gomath.Floor(val))%2 == 0 {
		return rp.ColorA
	}
	return rp.ColorB
}

func (rp *RingPattern) ColorAtObject(point math.Point, obj Shape) math.Color {
	objectPoint := WorldToObject(obj, point)
	patternPoint := rp.GetInverseTransform().MulT(objectPoint)

	return rp.ColorAt(patternPoint)
}

func (rp *RingPattern) SetTransform(tf math.Matrix) {
	rp.Transform = tf
}

func (rp *RingPattern) GetTransform() math.Matrix {
	return rp.Transform
}

func (rp *RingPattern) Equals(other Pattern) bool {
	if reflect.TypeOf(rp) == reflect.TypeOf(other) {
		otherStruct := other.(*RingPattern)
		return rp.ColorA.Equals(otherStruct.ColorA) &&
			rp.ColorB.Equals(otherStruct.ColorB) &&
			rp.Transform.Equals(otherStruct.GetTransform())
	}
	return false
}

func (rp *RingPattern) GetInverseTransform() math.Matrix {
	return rp.InverseTransform
}

func (rp *RingPattern) CalculateInverseTransform() {
	rp.InverseTransform = rp.Transform.Inverse()
}
