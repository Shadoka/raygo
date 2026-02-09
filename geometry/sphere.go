package geometry

import (
	gomath "math"
	"raygo/math"
	"reflect"

	"github.com/google/uuid"
)

type Sphere struct {
	Id               string
	Transform        math.Matrix
	Material         Material
	Parent           *Group
	InverseTransform math.Matrix
}

func CreateSphere() *Sphere {
	return &Sphere{
		Id:        uuid.NewString(),
		Transform: math.IdentityMatrix(),
		Material:  DefaultMaterial(),
		Parent:    nil,
	}
}

func CreateGlassSphere() *Sphere {
	m := DefaultMaterial()
	m.SetTransparency(1.0)
	m.SetRefractiveIndex(1.5)
	return &Sphere{
		Id:        uuid.NewString(),
		Transform: math.IdentityMatrix(),
		Material:  m,
	}
}

func (s *Sphere) Equals(other Shape) bool {
	return reflect.TypeOf(s) == reflect.TypeOf(other) &&
		s.Transform.Equals(other.GetTransform()) &&
		s.Material.Equals(*other.GetMaterial()) &&
		s.Parent == other.GetParent()
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

func (s *Sphere) GetMaterial() *Material {
	return &s.Material
}

func (s *Sphere) SetMaterial(m Material) {
	s.Material = m
}

func (s *Sphere) GetParent() *Group {
	return s.Parent
}

func (s *Sphere) SetParent(g *Group) {
	s.Parent = g
}

func (sphere *Sphere) Intersect(ray Ray) []Intersection {
	// the vector from the sphere's center, to the ray origin
	// remember: the sphere is centered at the world origin
	transformedRay := ray.Transform(sphere.GetInverseTransform())
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

func (s *Sphere) NormalAt(p math.Point, hit Intersection) math.Vector {
	objectSpace := WorldToObject(s, p)
	objectNormal := objectSpace.Subtract(math.CreatePoint(0.0, 0.0, 0.0))
	return NormalToWorld(s, objectNormal)
}

func (s *Sphere) Bounds() *Bounds {
	return &Bounds{
		Minimum: math.CreatePoint(-1.0, -1.0, -1.0),
		Maximum: math.CreatePoint(1.0, 1.0, 1.0),
	}
}

func (s *Sphere) GetInverseTransform() math.Matrix {
	return s.InverseTransform
}

func (s *Sphere) CalculateInverseTransform() {
	s.InverseTransform = s.Transform.Inverse()
}

func (s *Sphere) GetUvCoordinate(direction math.Vector) (float64, float64) {
	u := 0.5 - gomath.Atan2(direction.Z, direction.X)/(2.*gomath.Pi)
	v := 0.5 + gomath.Asin(direction.Y)/gomath.Pi
	return u, v
}
