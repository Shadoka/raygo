package geometry

import (
	gomath "math"
	"raygo/math"
	"reflect"
)

type GradientPattern struct {
	ColorA           math.Color
	ColorB           math.Color
	Transform        math.Matrix
	InverseTransform *math.Matrix
}

func CreateGradientPattern(a math.Color, b math.Color) *GradientPattern {
	return &GradientPattern{
		ColorA:           a,
		ColorB:           b,
		Transform:        math.IdentityMatrix(),
		InverseTransform: nil,
	}
}

func (gp *GradientPattern) ColorAt(point math.Point) math.Color {
	distance := gp.ColorB.Subtract(gp.ColorA)
	fraction := point.X - gomath.Floor(point.X)

	return gp.ColorA.Add(distance.Mul(fraction))
}

func (gp *GradientPattern) ColorAtObject(point math.Point, obj Shape) math.Color {
	objectPoint := WorldToObject(obj, point)
	patternPoint := gp.GetInverseTransform().MulT(objectPoint)

	return gp.ColorAt(patternPoint)
}

func (gp *GradientPattern) SetTransform(tf math.Matrix) {
	gp.Transform = tf
}

func (gp *GradientPattern) GetTransform() math.Matrix {
	return gp.Transform
}

func (gp *GradientPattern) Equals(other Pattern) bool {
	if reflect.TypeOf(gp) == reflect.TypeOf(other) {
		otherStruct := other.(*StripePattern)
		return gp.ColorA.Equals(otherStruct.ColorA) &&
			gp.ColorB.Equals(otherStruct.ColorB) &&
			gp.Transform.Equals(otherStruct.GetTransform())
	}
	return false
}

func (gp *GradientPattern) GetInverseTransform() math.Matrix {
	if gp.InverseTransform != nil {
		return *gp.InverseTransform
	}

	inverse := gp.Transform.Inverse()
	gp.InverseTransform = &inverse
	return *gp.InverseTransform
}
