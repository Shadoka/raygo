package ray

import (
	gomath "math"
	"raygo/math"

	"github.com/google/uuid"
)

type Shape interface {
	Equals(other Shape) bool
	GetId() string
	SetTransform(m math.Matrix)
	GetTransform() math.Matrix
	Intersect(ray Ray) []Intersection
	NormalAt(p math.Point) math.Vector
}

type Sphere struct {
	Id        string
	Transform math.Matrix
}

func CreateSphere() *Sphere {
	return &Sphere{
		Id:        uuid.NewString(),
		Transform: math.IdentityMatrix(),
	}
}

func (s *Sphere) Equals(other Shape) bool {
	return s.Id == other.GetId()
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
