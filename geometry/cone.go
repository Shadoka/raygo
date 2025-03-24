package geometry

import (
	gomath "math"
	"raygo/math"
	"reflect"

	"github.com/google/uuid"
)

type Cone struct {
	Id        string
	Transform math.Matrix
	Material  Material
	Minimum   float64
	Maximum   float64
	Closed    bool
}

func CreateCone() *Cone {
	return &Cone{
		Id:        uuid.NewString(),
		Transform: math.IdentityMatrix(),
		Material:  DefaultMaterial(),
		Minimum:   gomath.Inf(-1.0),
		Maximum:   gomath.Inf(1.0),
		Closed:    false,
	}
}

func (c *Cone) SetTransform(m math.Matrix) {
	c.Transform = m
}

func (c *Cone) GetId() string {
	return c.Id
}

func (c *Cone) GetTransform() math.Matrix {
	return c.Transform
}

func (c *Cone) GetMaterial() *Material {
	return &c.Material
}

func (c *Cone) SetMaterial(m Material) {
	c.Material = m
}

func (c *Cone) Equals(other Shape) bool {
	return reflect.TypeOf(c) == reflect.TypeOf(other) &&
		c.Transform.Equals(other.GetTransform()) &&
		c.Material.Equals(*other.GetMaterial())
}

func (c *Cone) Intersect(ray Ray) []Intersection {
	transformedRay := ray.Transform(c.Transform.Inverse())
	return c.localConeIntersect(transformedRay)
}

func (co *Cone) localConeIntersect(ray Ray) []Intersection {
	xs := make([]Intersection, 0)
	a := gomath.Pow(ray.Direction.X, 2) - gomath.Pow(ray.Direction.Y, 2) + gomath.Pow(ray.Direction.Z, 2)

	b := 2*ray.Origin.X*ray.Direction.X -
		2*ray.Origin.Y*ray.Direction.Y +
		2*ray.Origin.Z*ray.Direction.Z

	if floatEquals(a, 0.0) && floatEquals(b, 0.0) {
		// return xs
		return co.intersectCaps(ray)
	}

	c := gomath.Pow(ray.Origin.X, 2) - gomath.Pow(ray.Origin.Y, 2) + gomath.Pow(ray.Origin.Z, 2)

	if floatEquals(a, 0.0) && !floatEquals(b, 0.0) {
		xs = append(xs, CreateIntersection(-c/(2*b), co))
		// return xs
	}

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
	if co.Minimum < y0 && y0 < co.Maximum {
		xs = append(xs, CreateIntersection(t0, co))
	}

	y1 := ray.Origin.Y + t1*ray.Direction.Y
	if co.Minimum < y1 && y1 < co.Maximum {
		xs = append(xs, CreateIntersection(t1, co))
	}

	xs = append(xs, co.intersectCaps(ray)...)

	return xs
}

func (c *Cone) checkCap(ray Ray, t float64, radius float64) bool {
	x := ray.Origin.X + t*ray.Direction.X
	z := ray.Origin.Z + t*ray.Direction.Z

	return (x*x + z*z) <= gomath.Abs(radius)
}

func (c *Cone) intersectCaps(ray Ray) []Intersection {
	xs := make([]Intersection, 0)
	if !c.Closed || floatEquals(ray.Direction.Y, 0.0) {
		return xs
	}

	t := (c.Minimum - ray.Origin.Y) / ray.Direction.Y
	if c.checkCap(ray, t, c.Minimum) {
		xs = append(xs, CreateIntersection(t, c))
	}

	t = (c.Maximum - ray.Origin.Y) / ray.Direction.Y
	if c.checkCap(ray, t, c.Maximum) {
		xs = append(xs, CreateIntersection(t, c))
	}

	return xs
}

func (c *Cone) NormalAt(point math.Point) math.Vector {
	objectSpace := c.Transform.Inverse().MulT(point)
	objectNormal := c.localConeNormalAt(objectSpace)
	worldNormal := c.Transform.Inverse().Transpose().MulT(objectNormal)
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

func (c *Cone) localConeNormalAt(point math.Point) math.Vector {
	dist := point.X*point.X + point.Z*point.Z

	if dist < 1 && point.Y >= c.Maximum-math.EPSILON {
		return math.CreateVector(0.0, 1.0, 0.0)
	} else if dist < 1 && point.Y <= c.Minimum+math.EPSILON {
		return math.CreateVector(0.0, -1.0, 0.0)
	}

	y := gomath.Sqrt(point.X*point.X + point.Z*point.Z)
	if point.Y > 0.0 {
		y = -y
	}

	return math.CreateVector(point.X, y, point.Z)
}
