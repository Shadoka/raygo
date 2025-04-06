package geometry

import (
	gomath "math"
	"raygo/math"
	"reflect"

	"github.com/google/uuid"
)

type Plane struct {
	Id        string
	Transform math.Matrix
	Material  Material
	Parent    *Group
}

func CreatePlane() *Plane {
	return &Plane{
		Id:        uuid.NewString(),
		Transform: math.IdentityMatrix(),
		Material:  DefaultMaterial(),
		Parent:    nil,
	}
}

func (p *Plane) SetTransform(m math.Matrix) {
	p.Transform = m
}

func (p *Plane) GetId() string {
	return p.Id
}

func (p *Plane) GetTransform() math.Matrix {
	return p.Transform
}

func (p *Plane) GetMaterial() *Material {
	return &p.Material
}

func (p *Plane) SetMaterial(m Material) {
	p.Material = m
}

func (p *Plane) Equals(other Shape) bool {
	return reflect.TypeOf(p) == reflect.TypeOf(other) &&
		p.Transform.Equals(other.GetTransform()) &&
		p.Material.Equals(*other.GetMaterial()) &&
		p.Parent == other.GetParent()
}

func (p *Plane) GetParent() *Group {
	return p.Parent
}

func (p *Plane) SetParent(g *Group) {
	p.Parent = g
}

func (p *Plane) NormalAt(point math.Point) math.Vector {
	objectNormal := p.localPlaneNormalAt()
	return NormalToWorld(p, objectNormal)
}

func (p *Plane) localPlaneNormalAt() math.Vector {
	return math.CreateVector(0.0, 1.0, 0.0)
}

func (p *Plane) Intersect(ray Ray) []Intersection {
	transformedRay := ray.Transform(p.Transform.Inverse())
	return p.localPlaneIntersect(transformedRay)
}

func (p *Plane) localPlaneIntersect(localRay Ray) []Intersection {
	xs := make([]Intersection, 0)
	if gomath.Abs(localRay.Direction.Y) < math.EPSILON {
		return xs
	}

	t := -localRay.Origin.Y / localRay.Direction.Y
	xs = append(xs, CreateIntersection(t, p))

	return xs
}

func (p *Plane) Bounds() *Bounds {
	return &Bounds{
		Minimum: p.Transform.MulT(math.CreatePoint(gomath.Inf(-1), -0.1, gomath.Inf(-1))),
		Maximum: p.Transform.MulT(math.CreatePoint(gomath.Inf(1), 0.1, gomath.Inf(1))),
	}
}
