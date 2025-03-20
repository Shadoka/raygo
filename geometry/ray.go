package geometry

import (
	gomath "math"
	"raygo/math"
	"sort"
)

const EPSILON = 0.00001

type Ray struct {
	Origin    math.Point
	Direction math.Vector
}

type Intersection struct {
	IntersectionAt float64 // t value
	Object         Shape
}

type IntersectionComputations struct {
	IntersectionAt float64 // t value
	Object         Shape
	Point          math.Point
	Eyev           math.Vector
	Normalv        math.Vector
	Reflectv       math.Vector
	Inside         bool
	OverPoint      math.Point
}

func CreateRay(origin math.Point, direction math.Vector) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func CreateIntersection(t float64, object Shape) Intersection {
	return Intersection{
		IntersectionAt: t,
		Object:         object,
	}
}

func (i Intersection) Equals(other Intersection) bool {
	return i.IntersectionAt == other.IntersectionAt &&
		i.Object.Equals(other.Object)
}

func (r Ray) Position(t float64) math.Vector {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (i Intersection) Aggregate(other Intersection) []Intersection {
	result := make([]Intersection, 0)
	result = append(result, i)
	result = append(result, other)
	return result
}

func AddIntersection(xs []Intersection, toAdd Intersection) []Intersection {
	return append(xs, toAdd)
}

func Hit(xs []Intersection) *Intersection {
	var result *Intersection
	minT := gomath.MaxFloat64
	for _, v := range xs {
		if v.IntersectionAt > 0.0 && v.IntersectionAt <= minT {
			result = &v
			minT = v.IntersectionAt
		}
	}
	return result
}

func (r Ray) Transform(transformMatrix math.Matrix) Ray {
	newOrigin := transformMatrix.MulT(r.Origin)
	newDirection := transformMatrix.MulT(r.Direction)
	return CreateRay(newOrigin, newDirection)
}

func SortIntersections(xs []Intersection) {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].IntersectionAt < xs[j].IntersectionAt
	})
}

func (i Intersection) PrepareComputation(r Ray) IntersectionComputations {
	p := r.Position(i.IntersectionAt)
	normal := i.Object.NormalAt(p)
	eye := r.Direction.Negate()

	inside := false
	if normal.Dot(eye) < 0.0 {
		inside = true
		normal = normal.Negate()
	}

	overP := p.Add(normal.Mul(EPSILON))
	reflect := r.Direction.Reflect(normal)

	return IntersectionComputations{
		IntersectionAt: i.IntersectionAt,
		Object:         i.Object,
		Point:          p,
		Eyev:           eye,
		Normalv:        normal,
		Inside:         inside,
		OverPoint:      overP,
		Reflectv:       reflect,
	}
}
