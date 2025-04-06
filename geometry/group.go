package geometry

import (
	gomath "math"
	"raygo/math"
	"reflect"

	"github.com/google/uuid"
)

type Group struct {
	Id        string
	Transform math.Matrix
	Material  Material
	Children  []Shape
	Parent    *Group
}

func EmptyGroup() *Group {
	return &Group{
		Id:        uuid.NewString(),
		Transform: math.IdentityMatrix(),
		Material:  DefaultMaterial(),
		Children:  make([]Shape, 0),
		Parent:    nil,
	}
}

func (g *Group) Equals(other Shape) bool {
	if reflect.TypeOf(g) == reflect.TypeOf(other) {
		otherGroup := other.(*Group) // does this crash when other is Group and not *Group?
		return g.Id == otherGroup.Id &&
			g.Material.Equals(otherGroup.Material) &&
			g.Transform.Equals(otherGroup.Transform) &&
			len(g.Children) == len(otherGroup.Children) && // TODO: do equals here, too lazy now
			g.Parent == otherGroup.Parent
	}
	return false
}

func (g *Group) GetId() string {
	return g.Id
}

func (g *Group) SetTransform(m math.Matrix) {
	g.Transform = m
}

func (g *Group) GetTransform() math.Matrix {
	return g.Transform
}

func (g *Group) SetMaterial(m Material) {
	g.Material = m
}

func (g *Group) GetMaterial() *Material {
	return &g.Material
}

func (g *Group) GetParent() *Group {
	return g.Parent
}

func (gr *Group) SetParent(g *Group) {
	gr.Parent = g
}

func (g *Group) AddChild(child Shape) *Group {
	g.Children = append(g.Children, child)
	child.SetParent(g)
	return g
}

func (g *Group) Size() int {
	size := 0
	for _, v := range g.Children {
		if reflect.TypeOf(v) == reflect.TypeOf(g) {
			size += v.(*Group).Size()
		} else {
			size++
		}
	}
	return size
}

func (g *Group) Intersect(ray Ray) []Intersection {
	transformedRay := ray.Transform(g.Transform.Inverse())
	return g.localIntersect(transformedRay)
}

func (g *Group) localIntersect(ray Ray) []Intersection {
	xs := make([]Intersection, 0)
	if len(g.boundsIntersect(ray)) == 0 {
		return xs
	}

	for _, child := range g.Children {
		xs = append(xs, child.Intersect(ray)...)
	}
	SortIntersections(xs)
	return xs
}

func (g *Group) NormalAt(p math.Point) math.Vector {
	return math.CreateVector(0.0, 1.0, 0.0)
}

func (g *Group) Bounds() *Bounds {
	// those BBs are in their respective object space but not axis aligned
	nonAlignedBoundingBoxes := make([]*Bounds, 0)
	for _, child := range g.Children {
		nonAlignedBoundingBoxes = append(nonAlignedBoundingBoxes, child.Bounds())
	}

	alignedMinimumBB := FindMinimalContainingBoundingBox(nonAlignedBoundingBoxes)
	// TODO: wouldnt it make sense to cache that group BB?
	return BoundsToObjectSpace(alignedMinimumBB, g.Transform)
}

func FindMinimalContainingBoundingBox(bounds []*Bounds) Bounds {
	minX, minY, minZ := gomath.MaxFloat64, gomath.MaxFloat64, gomath.MaxFloat64
	maxX, maxY, maxZ := gomath.SmallestNonzeroFloat64, gomath.SmallestNonzeroFloat64, gomath.SmallestNonzeroFloat64

	for _, b := range bounds {
		minX = gomath.Min(minX, b.Minimum.X)
		minY = gomath.Min(minY, b.Minimum.Y)
		minZ = gomath.Min(minZ, b.Minimum.Z)
		maxX = gomath.Max(maxX, b.Maximum.X)
		maxY = gomath.Max(maxY, b.Maximum.Y)
		maxZ = gomath.Max(maxZ, b.Maximum.Z)
	}

	return Bounds{
		Minimum: math.CreatePoint(minX, minY, minZ),
		Maximum: math.CreatePoint(maxX, maxY, maxZ),
	}
}

// TODO: the following two methods are basically the cube intersection from cube.go

func (g *Group) boundsIntersect(localRay Ray) []Intersection {
	bounds := g.Bounds()
	xtmin, xtmax := checkAxisBoundsIntersect(localRay.Origin.X, localRay.Direction.X, bounds.Minimum.X, bounds.Maximum.X)
	ytmin, ytmax := checkAxisBoundsIntersect(localRay.Origin.Y, localRay.Direction.Y, bounds.Minimum.Y, bounds.Maximum.Y)
	ztmin, ztmax := checkAxisBoundsIntersect(localRay.Origin.Z, localRay.Direction.Z, bounds.Minimum.Z, bounds.Maximum.Z)

	tmin := gomath.Max(gomath.Max(xtmin, ytmin), ztmin)
	tmax := gomath.Min(gomath.Min(xtmax, ytmax), ztmax)

	xs := make([]Intersection, 0)

	if tmin > tmax {
		return xs
	}

	xs = append(xs, CreateIntersection(tmin, g))
	xs = append(xs, CreateIntersection(tmax, g))
	return xs
}

func checkAxisBoundsIntersect(origin float64, direction float64, minBound float64, maxBound float64) (float64, float64) {
	tminNumerator := minBound - origin
	tmaxNumerator := maxBound - origin

	tmin := 0.0
	tmax := 0.0
	if gomath.Abs(direction) >= math.EPSILON {
		tmin = tminNumerator / direction
		tmax = tmaxNumerator / direction
	} else {
		tmin = tminNumerator * gomath.SmallestNonzeroFloat64
		tmax = tmaxNumerator * gomath.MaxFloat64
	}

	if tmin > tmax {
		return tmax, tmin
	}
	return tmin, tmax
}
