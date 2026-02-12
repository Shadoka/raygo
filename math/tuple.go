package math

import (
	"fmt"
	gomath "math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

type Point = Tuple
type Vector = Tuple
type Color = Tuple

func CreateVector(x float64, y float64, z float64) Vector {
	return Vector{X: x,
		Y: y,
		Z: z,
		W: 0.0}
}

func CreatePoint(x float64, y float64, z float64) Point {
	return Point{X: x,
		Y: y,
		Z: z,
		W: 1.0}
}

func CreateColor(x float64, y float64, z float64) Color {
	return Color{
		X: x,
		Y: y,
		Z: z,
		W: 0.0,
	}
}

func CreateTuple(x float64, y float64, z float64, w float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t Vector) IsVector() bool {
	return t.W == 0.0
}

func (t Tuple) Equals(other Tuple) bool {
	return floatEquals(t.X, other.X) &&
		floatEquals(t.Y, other.Y) &&
		floatEquals(t.Z, other.Z) &&
		floatEquals(t.W, other.W)
}

func (t Tuple) Add(other Tuple) Tuple {
	return Tuple{
		X: t.X + other.X,
		Y: t.Y + other.Y,
		Z: t.Z + other.Z,
		W: t.W + other.W,
	}
}

func (t Tuple) Subtract(other Tuple) Tuple {
	if t.W-other.W < 0 {
		panic("tuple subtraction leads to w component < 0")
	}

	return Tuple{
		X: t.X - other.X,
		Y: t.Y - other.Y,
		Z: t.Z - other.Z,
		W: t.W - other.W,
	}
}

func (t Tuple) Negate() Tuple {
	return Tuple{
		X: 0.0 - t.X,
		Y: 0.0 - t.Y,
		Z: 0.0 - t.Z,
		W: 0.0 - t.W,
	}
}

func (t Tuple) Mul(s float64) Tuple {
	return Tuple{
		X: t.X * s,
		Y: t.Y * s,
		Z: t.Z * s,
		W: t.W * s,
	}
}

func (t Tuple) Div(s float64) Tuple {
	return Tuple{
		X: t.X / s,
		Y: t.Y / s,
		Z: t.Z / s,
		W: t.W / s,
	}
}

func (v Vector) Magnitude() float64 {
	return gomath.Sqrt(v.X*v.X +
		v.Y*v.Y +
		v.Z*v.Z +
		v.W*v.W)
}

func (v Vector) Normalize() Vector {
	mag := v.Magnitude()
	return Vector{
		X: v.X / mag,
		Y: v.Y / mag,
		Z: v.Z / mag,
		W: v.W / mag,
	}
}

func (v Vector) Dot(other Vector) float64 {
	return v.X*other.X +
		v.Y*other.Y +
		v.Z*other.Z +
		v.W*other.W
}

func (v Vector) Cross(other Vector) Vector {
	x := v.Y*other.Z - v.Z*other.Y
	y := v.Z*other.X - v.X*other.Z
	z := v.X*other.Y - v.Y*other.X
	return CreateVector(x, y, z)
}

func (c Color) Blend(other Color) Color {
	return Color{
		X: c.X * other.X,
		Y: c.Y * other.Y,
		Z: c.Z * other.Z,
		W: c.W * other.W,
	}
}

func (v Vector) Reflect(n Vector) Vector {
	return v.Subtract(n.Mul(2.0).Mul(v.Dot(n)))
}

func (t Tuple) Abs() Tuple {
	return Tuple{
		X: gomath.Abs(t.X),
		Y: gomath.Abs(t.Y),
		Z: gomath.Abs(t.Z),
		W: gomath.Abs(t.W),
	}
}

func (p Point) ToString() string {
	return fmt.Sprintf("(%v, %v, %v)", p.X, p.Y, p.Z)
}
