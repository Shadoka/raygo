package geometry

import (
	gomath "math"
	"raygo/math"
)

type Bounds struct {
	Minimum math.Point
	Maximum math.Point
}

func (b *Bounds) ApplyTransform(tf math.Matrix) *Bounds {
	currentCorners := CalculateBBCorners(*b)
	transformedCorners := make([]math.Point, 8)
	for i, c := range currentCorners {
		transformedCorners[i] = tf.MulT(c)
	}
	return FindBoundingBox(transformedCorners)
}

func CalculateBBCorners(b Bounds) []math.Point {
	corners := make([]math.Point, 0)
	corners = append(corners, b.Minimum, b.Maximum)

	p1 := math.CreatePoint(b.Minimum.X, b.Minimum.Y, b.Maximum.Z)
	p2 := math.CreatePoint(b.Minimum.X, b.Maximum.Y, b.Minimum.Z)
	p3 := math.CreatePoint(b.Minimum.X, b.Maximum.Y, b.Maximum.Z)
	p4 := math.CreatePoint(b.Maximum.X, b.Minimum.Y, b.Minimum.Z)
	p5 := math.CreatePoint(b.Maximum.X, b.Maximum.Y, b.Minimum.Z)
	p6 := math.CreatePoint(b.Maximum.X, b.Minimum.Y, b.Maximum.Z)
	corners = append(corners, p1, p2, p3, p4, p5, p6)

	return corners
}

func FindBoundingBox(corners []math.Point) *Bounds {
	minX, minY, minZ := gomath.MaxFloat64, gomath.MaxFloat64, gomath.MaxFloat64
	maxX, maxY, maxZ := gomath.SmallestNonzeroFloat64, gomath.SmallestNonzeroFloat64, gomath.SmallestNonzeroFloat64

	for _, p := range corners {
		minX = gomath.Min(minX, p.X)
		minY = gomath.Min(minY, p.Y)
		minZ = gomath.Min(minZ, p.Z)
		maxX = gomath.Max(maxX, p.X)
		maxY = gomath.Max(maxY, p.Y)
		maxZ = gomath.Max(maxZ, p.Z)
	}

	return &Bounds{
		Minimum: math.CreatePoint(minX, minY, minZ),
		Maximum: math.CreatePoint(maxX, maxY, maxZ),
	}
}

func FindMinimalContainingBoundingBox(bounds []*Bounds) *Bounds {
	points := make([]math.Point, 0)

	for _, b := range bounds {
		points = append(points, b.Minimum, b.Maximum)
	}

	return FindBoundingBox(points)
}

func BoundingBoxIntersect(localRay Ray, s Shape, min math.Point, max math.Point) []Intersection {
	// see https://tavianator.com/2011/ray_box.html
	xs := make([]Intersection, 0)
	tmin, tmax := branchlessCheck(localRay, min, max)
	if tmax >= gomath.Max(0.0, tmin) && tmin < gomath.Inf(0) {
		xs = append(xs, CreateIntersection(tmin, s))
		xs = append(xs, CreateIntersection(tmax, s))
	}
	return xs
}

func branchlessCheck(ray Ray, min math.Point, max math.Point) (float64, float64) {
	minX := (min.X - ray.Origin.X) / ray.Direction.X
	maxX := (max.X - ray.Origin.X) / ray.Direction.X

	tmin := gomath.Min(minX, maxX)
	tmax := gomath.Max(minX, maxX)

	minY := (min.Y - ray.Origin.Y) / ray.Direction.Y
	maxY := (max.Y - ray.Origin.Y) / ray.Direction.Y

	tmin = gomath.Max(tmin, gomath.Min(minY, maxY))
	tmax = gomath.Min(tmax, gomath.Max(minY, maxY))

	minZ := (min.Z - ray.Origin.Z) / ray.Direction.Z
	maxZ := (max.Z - ray.Origin.Z) / ray.Direction.Z

	tmin = gomath.Max(tmin, gomath.Min(minZ, maxZ))
	tmax = gomath.Min(tmax, gomath.Max(minZ, maxZ))

	return tmin, tmax
}

func (b *Bounds) Equals(other *Bounds) bool {
	return b.Minimum.Equals(other.Minimum) &&
		b.Maximum.Equals(other.Maximum)
}
