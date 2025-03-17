package ray

import (
	gomath "math"
	"raygo/math"
	"reflect"

	"github.com/google/uuid"
)

type Shape interface {
	Equals(other Shape) bool
	GetId() string
	SetTransform(m math.Matrix)
	GetTransform() math.Matrix
	SetMaterial(m Material)
	GetMaterial() Material
	Intersect(ray Ray) []Intersection
	NormalAt(p math.Point) math.Vector
}

type Sphere struct {
	Id        string
	Transform math.Matrix
	Material  Material
}

type Plane struct {
	Id        string
	Transform math.Matrix
	Material  Material
}

func CreateSphere() *Sphere {
	return &Sphere{
		Id:        uuid.NewString(),
		Transform: math.IdentityMatrix(),
		Material:  DefaultMaterial(),
	}
}

func CreatePlane() *Plane {
	return &Plane{
		Id:        uuid.NewString(),
		Transform: math.IdentityMatrix(),
		Material:  DefaultMaterial(),
	}
}

// BEGIN SPHERE FUNCTIONS

func (s *Sphere) Equals(other Shape) bool {
	return reflect.TypeOf(s) == reflect.TypeOf(other) &&
		s.Transform.Equals(other.GetTransform()) &&
		s.Material.Equals(other.GetMaterial())
}

func (s *Sphere) SetTransform(m math.Matrix) {
	s.Transform = m
}

func (s *Sphere) GetId() string {
	return s.Id
}

func (s *Sphere) GetTransform() math.Matrix {
	return s.Transform
}

func (s *Sphere) GetMaterial() Material {
	return s.Material
}

func (s *Sphere) SetMaterial(m Material) {
	s.Material = m
}

func (sphere *Sphere) Intersect(ray Ray) []Intersection {
	// the vector from the sphere's center, to the ray origin
	// remember: the sphere is centered at the world origin
	transformedRay := ray.Transform(sphere.Transform.Inverse())
	sphereToRay := transformedRay.Origin.Subtract(math.CreatePoint(0.0, 0.0, 0.0))

	a := transformedRay.Direction.Dot(transformedRay.Direction)
	b := 2.0 * transformedRay.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1.0

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return make([]Intersection, 0)
	}

	t1 := (-b - gomath.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + gomath.Sqrt(discriminant)) / (2 * a)

	return CreateIntersection(t1, sphere).Aggregate(CreateIntersection(t2, sphere))
}

func (s *Sphere) NormalAt(p math.Point) math.Vector {
	objectSpace := s.Transform.Inverse().MulT(p)
	objectNormal := objectSpace.Subtract(math.CreatePoint(0.0, 0.0, 0.0))
	worldNormal := s.Transform.Inverse().Transpose().MulT(objectNormal)
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

// END SPHERE FUNCTIONS

// BEGIN PLANE FUNCTIONS

func (p *Plane) SetTransform(m math.Matrix) {
	p.Transform = m
}

func (p *Plane) GetId() string {
	return p.Id
}

func (p *Plane) GetTransform() math.Matrix {
	return p.Transform
}

func (p *Plane) GetMaterial() Material {
	return p.Material
}

func (p *Plane) SetMaterial(m Material) {
	p.Material = m
}

func (p *Plane) Equals(other Shape) bool {
	return reflect.TypeOf(p) == reflect.TypeOf(other) &&
		p.Transform.Equals(other.GetTransform()) &&
		p.Material.Equals(other.GetMaterial())
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

// END PLANE FUNCTIONS
