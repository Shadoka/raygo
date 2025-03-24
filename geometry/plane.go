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
}

func CreatePlane() *Plane {
	return &Plane{
		Id:        uuid.NewString(),
		Transform: math.IdentityMatrix(),
		Material:  DefaultMaterial(),
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
		p.Material.Equals(*other.GetMaterial())
}

func (p *Plane) NormalAt(point math.Point) math.Vector {
	objectNormal := p.localPlaneNormalAt()
	worldNormal := p.Transform.Inverse().Transpose().MulT(objectNormal)
	worldNormal.W = 0.0
	return worldNormal.Normalize()
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
