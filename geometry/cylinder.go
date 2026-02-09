package geometry

import (
	"log"
	gomath "math"
	"raygo/math"
	"reflect"

	"github.com/google/uuid"
)

type Cylinder struct {
	Id               string
	Transform        math.Matrix
	Material         Material
	Minimum          float64
	Maximum          float64
	Closed           bool
	Parent           *Group
	InverseTransform math.Matrix
}

func CreateCylinder() *Cylinder {
	return &Cylinder{
		Id:        uuid.NewString(),
		Transform: math.IdentityMatrix(),
		Material:  DefaultMaterial(),
		Minimum:   gomath.Inf(-1.0),
		Maximum:   gomath.Inf(1.0),
		Closed:    false,
		Parent:    nil,
	}
}

func (c *Cylinder) SetTransform(m math.Matrix) {
	c.Transform = m
}

func (c *Cylinder) GetId() string {
	return c.Id
}

func (c *Cylinder) GetTransform() math.Matrix {
	return c.Transform
}

func (c *Cylinder) GetMaterial() *Material {
	return &c.Material
}

func (c *Cylinder) SetMaterial(m Material) {
	c.Material = m
}

func (c *Cylinder) GetParent() *Group {
	return c.Parent
}

func (c *Cylinder) SetParent(g *Group) {
	c.Parent = g
}

func (c *Cylinder) Equals(other Shape) bool {
	return reflect.TypeOf(c) == reflect.TypeOf(other) &&
		c.Transform.Equals(other.GetTransform()) &&
		c.Material.Equals(*other.GetMaterial()) &&
		c.Parent == other.GetParent()
}

func (c *Cylinder) Intersect(ray Ray) []Intersection {
	transformedRay := ray.Transform(c.GetInverseTransform())
	return c.localCylinderIntersect(transformedRay)
}

func (cy *Cylinder) localCylinderIntersect(ray Ray) []Intersection {
	xs := make([]Intersection, 0)
	a := gomath.Pow(ray.Direction.X, 2) + gomath.Pow(ray.Direction.Z, 2)

	if floatEquals(a, 0.0) {
		return cy.intersectCaps(ray)
	}

	b := 2*ray.Origin.X*ray.Direction.X +
		2*ray.Origin.Z*ray.Direction.Z

	c := gomath.Pow(ray.Origin.X, 2) + gomath.Pow(ray.Origin.Z, 2) - 1.0

	disc := b*b - 4*a*c

	if disc < 0 {
		return xs
	}

	t0 := (-b - gomath.Sqrt(disc)) / (2 * a)
	t1 := (-b + gomath.Sqrt(disc)) / (2 * a)

	if t0 > t1 {
		tmp := t1
		t1 = t0
		t0 = tmp
	}

	y0 := ray.Origin.Y + t0*ray.Direction.Y
	if cy.Minimum < y0 && y0 < cy.Maximum {
		xs = append(xs, CreateIntersection(t0, cy))
	}

	y1 := ray.Origin.Y + t1*ray.Direction.Y
	if cy.Minimum < y1 && y1 < cy.Maximum {
		xs = append(xs, CreateIntersection(t1, cy))
	}

	xs = append(xs, cy.intersectCaps(ray)...)

	return xs
}

func (c *Cylinder) checkCap(ray Ray, t float64) bool {
	x := ray.Origin.X + t*ray.Direction.X
	z := ray.Origin.Z + t*ray.Direction.Z

	return (x*x + z*z) <= 1.0
}

func (c *Cylinder) intersectCaps(ray Ray) []Intersection {
	xs := make([]Intersection, 0)
	if !c.Closed || floatEquals(ray.Direction.Y, 0.0) {
		return xs
	}

	t := (c.Minimum - ray.Origin.Y) / ray.Direction.Y
	if c.checkCap(ray, t) {
		xs = append(xs, CreateIntersection(t, c))
	}

	t = (c.Maximum - ray.Origin.Y) / ray.Direction.Y
	if c.checkCap(ray, t) {
		xs = append(xs, CreateIntersection(t, c))
	}

	return xs
}

func (c *Cylinder) NormalAt(point math.Point, hit Intersection) math.Vector {
	objectSpace := WorldToObject(c, point)
	objectNormal := c.localCylinderNormalAt(objectSpace)
	return NormalToWorld(c, objectNormal)
}

func (c *Cylinder) localCylinderNormalAt(point math.Point) math.Vector {
	dist := point.X*point.X + point.Z*point.Z

	if dist < 1 && point.Y >= c.Maximum-math.EPSILON {
		return math.CreateVector(0.0, 1.0, 0.0)
	} else if dist < 1 && point.Y <= c.Minimum+math.EPSILON {
		return math.CreateVector(0.0, -1.0, 0.0)
	}

	return math.CreateVector(point.X, 0.0, point.Z)
}

func (c *Cylinder) Bounds() *Bounds {
	return &Bounds{
		Minimum: math.CreatePoint(-1.0, c.Minimum, -1.0),
		Maximum: math.CreatePoint(1.0, c.Maximum, 1.0),
	}
}

func (c *Cylinder) GetInverseTransform() math.Matrix {
	return c.InverseTransform
}

func (c *Cylinder) CalculateInverseTransform() {
	c.InverseTransform = c.Transform.Inverse()
}

func (c *Cylinder) GetUvCoordinate(direction math.Vector) (float64, float64) {
	log.Fatal("GetUvCoordinate NOP")
	return 0, 0
}
