package geometry

import (
	gomath "math"
	"raygo/math"
	"reflect"
)

type StripePattern struct {
	ColorA           math.Color
	ColorB           math.Color
	Transform        math.Matrix
	InverseTransform *math.Matrix
}

func CreateStripePattern(a math.Color, b math.Color) *StripePattern {
	return &StripePattern{
		ColorA:           a,
		ColorB:           b,
		Transform:        math.IdentityMatrix(),
		InverseTransform: nil,
	}
}

func (sp *StripePattern) ColorAt(point math.Point) math.Color {
	if int(gomath.Floor(point.X))%2 == 0 {
		return sp.ColorA
	}
	return sp.ColorB
}

func (sp *StripePattern) ColorAtObject(point math.Point, obj Shape) math.Color {
	objectPoint := WorldToObject(obj, point)
	patternPoint := sp.GetInverseTransform().MulT(objectPoint)

	return sp.ColorAt(patternPoint)
}

func (sp *StripePattern) GetTransform() math.Matrix {
	return sp.Transform
}

func (sp *StripePattern) SetTransform(tf math.Matrix) {
	sp.Transform = tf
}

func (sp *StripePattern) Equals(other Pattern) bool {
	if reflect.TypeOf(sp) == reflect.TypeOf(other) {
		concreteType := other.(*StripePattern)
		return sp.ColorA.Equals(concreteType.ColorA) &&
			sp.ColorB.Equals(concreteType.ColorB) &&
			sp.Transform.Equals(concreteType.GetTransform())
	}
	return false
}

func (sp *StripePattern) GetInverseTransform() math.Matrix {
	if sp.InverseTransform != nil {
		return *sp.InverseTransform
	}

	inverse := sp.Transform.Inverse()
	sp.InverseTransform = &inverse
	return *sp.InverseTransform
}
