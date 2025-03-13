package ray

import (
	"raygo/math"
)

type Ray struct {
	Origin    math.Point
	Direction math.Vector
}

func CreateRay(origin math.Point, direction math.Vector) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (r Ray) Position(t float64) math.Vector {
	return r.Origin.Add(r.Direction.Mul(t))
}
