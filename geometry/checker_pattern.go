package geometry

import (
	gomath "math"
	"raygo/math"
	"reflect"
)

type CheckerPattern struct {
	ColorA           math.Color
	ColorB           math.Color
	Transform        math.Matrix
	InverseTransform math.Matrix
}

func CreateCheckerPattern(a math.Color, b math.Color) *CheckerPattern {
	return &CheckerPattern{
		ColorA:    a,
		ColorB:    b,
		Transform: math.IdentityMatrix(),
	}
}

func (c *CheckerPattern) ColorAt(point math.Point) math.Color {
	sumDimensions := gomath.Floor(point.X) + gomath.Floor(point.Y) + gomath.Floor(point.Z)
	if int(sumDimensions)%2 == 0 {
		return c.ColorA
	}
	return c.ColorB
}

func (c *CheckerPattern) ColorAtObject(point math.Point, obj Shape) math.Color {
	objectPoint := WorldToObject(obj, point)
	patternPoint := c.GetInverseTransform().MulT(objectPoint)

	return c.ColorAt(patternPoint)
}

func (c *CheckerPattern) SetTransform(tf math.Matrix) {
	c.Transform = tf
}

func (c *CheckerPattern) GetTransform() math.Matrix {
	return c.Transform
}

func (c *CheckerPattern) Equals(other Pattern) bool {
	if reflect.TypeOf(c) == reflect.TypeOf(other) {
		otherStruct := other.(*CheckerPattern)
		return c.ColorA.Equals(otherStruct.ColorA) &&
			c.ColorB.Equals(otherStruct.ColorB) &&
			c.Transform.Equals(otherStruct.GetTransform())
	}
	return false
}

func (c *CheckerPattern) GetInverseTransform() math.Matrix {
	return c.InverseTransform
}

func (c *CheckerPattern) CalculateInverseTransform() {
	c.InverseTransform = c.Transform.Inverse()
}
