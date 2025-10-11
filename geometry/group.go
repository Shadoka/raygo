package geometry

import (
	"raygo/math"
	"reflect"

	"github.com/google/uuid"
)

type Group struct {
	Id                string
	Transform         math.Matrix
	Material          Material
	Children          []Shape
	Parent            *Group
	CachedBoundingBox *Bounds
	InverseTransform  *math.Matrix
}

func EmptyGroup() *Group {
	return &Group{
		Id:                uuid.NewString(),
		Transform:         math.IdentityMatrix(),
		Material:          DefaultMaterial(),
		Children:          make([]Shape, 0),
		Parent:            nil,
		CachedBoundingBox: nil,
		InverseTransform:  nil,
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
	for _, child := range g.Children {
		child.SetMaterial(m)
	}
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
	transformedRay := ray.Transform(g.GetInverseTransform())
	return g.localIntersect(transformedRay)
}

func (g *Group) localIntersect(ray Ray) []Intersection {
	xs := make([]Intersection, 0)
	bb := g.Bounds()
	if len(BoundingBoxIntersect(ray, g, bb.Minimum, bb.Maximum)) == 0 {
		return xs
	}

	for _, child := range g.Children {
		xs = append(xs, child.Intersect(ray)...)
	}
	SortIntersections(xs)
	return xs
}

func (g *Group) NormalAt(p math.Point, hit Intersection) math.Vector {
	return math.CreateVector(0.0, 1.0, 0.0)
}

func (g *Group) Bounds() *Bounds {
	if g.CachedBoundingBox != nil {
		return g.CachedBoundingBox
	}

	// those BBs are in their respective object space but not axis aligned
	nonAlignedBoundingBoxes := make([]*Bounds, 0)
	for _, child := range g.Children {
		childBB := child.Bounds()
		nonAlignedBoundingBoxes = append(nonAlignedBoundingBoxes, childBB.ApplyTransform(child.GetTransform()))
	}

	alignedMinimumBB := FindMinimalContainingBoundingBox(nonAlignedBoundingBoxes)
	g.CachedBoundingBox = alignedMinimumBB
	return g.CachedBoundingBox
}

func (g *Group) GetInverseTransform() math.Matrix {
	if g.InverseTransform != nil {
		return *g.InverseTransform
	}

	inverse := g.Transform.Inverse()
	g.InverseTransform = &inverse
	return *g.InverseTransform
}
