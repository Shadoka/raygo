package geometry

import (
	gomath "math"
	"raygo/math"
	"reflect"

	"github.com/google/uuid"
)

type Cube struct {
	Id        string
	Transform math.Matrix
	Material  Material
	Parent    *Group
}

func CreateCube() *Cube {
	return &Cube{
		Id:        uuid.NewString(),
		Transform: math.IdentityMatrix(),
		Material:  DefaultMaterial(),
		Parent:    nil,
	}
}

func (c *Cube) SetTransform(m math.Matrix) {
	c.Transform = m
}

func (c *Cube) GetId() string {
	return c.Id
}

func (c *Cube) GetTransform() math.Matrix {
	return c.Transform
}

func (c *Cube) GetMaterial() *Material {
	return &c.Material
}

func (c *Cube) SetMaterial(m Material) {
	c.Material = m
}

func (c *Cube) GetParent() *Group {
	return c.Parent
}

func (c *Cube) SetParent(g *Group) {
	c.Parent = g
}

func (c *Cube) Equals(other Shape) bool {
	return reflect.TypeOf(c) == reflect.TypeOf(other) &&
		c.Transform.Equals(other.GetTransform()) &&
		c.Material.Equals(*other.GetMaterial()) &&
		c.Parent == other.GetParent()
}

func (c *Cube) Intersect(ray Ray) []Intersection {
	transformedRay := ray.Transform(c.Transform.Inverse())
	b := c.Bounds()
	return BoundingBoxIntersect(transformedRay, c, b.Minimum, b.Maximum)
}

func (c *Cube) localCubeIntersect(localRay Ray) []Intersection {
	xtmin, xtmax := c.checkAxis(localRay.Origin.X, localRay.Direction.X)
	ytmin, ytmax := c.checkAxis(localRay.Origin.Y, localRay.Direction.Y)
	ztmin, ztmax := c.checkAxis(localRay.Origin.Z, localRay.Direction.Z)

	tmin := gomath.Max(gomath.Max(xtmin, ytmin), ztmin)
	tmax := gomath.Min(gomath.Min(xtmax, ytmax), ztmax)

	xs := make([]Intersection, 0)

	if tmin > tmax {
		return xs
	}

	xs = append(xs, CreateIntersection(tmin, c))
	xs = append(xs, CreateIntersection(tmax, c))
	return xs
}

func (c *Cube) checkAxis(origin float64, direction float64) (float64, float64) {
	tminNumerator := -1 - origin
	tmaxNumerator := 1 - origin

	tmin := 0.0
	tmax := 0.0
	if gomath.Abs(direction) >= math.EPSILON {
		tmin = tminNumerator / direction
		tmax = tmaxNumerator / direction
	} else {
		tmin = tminNumerator * gomath.MaxFloat64
		tmax = tmaxNumerator * gomath.MaxFloat64
	}

	if tmin > tmax {
		tmp := tmin
		tmin = tmax
		tmax = tmp
	}
	return tmin, tmax
}

func (c *Cube) NormalAt(point math.Point, hit Intersection) math.Vector {
	objectSpace := WorldToObject(c, point)
	objectNormal := c.localCubeNormalAt(objectSpace)
	return NormalToWorld(c, objectNormal)
}

func (c *Cube) localCubeNormalAt(point math.Point) math.Vector {
	maxCoordinate := gomath.Max(gomath.Max(gomath.Abs(point.X), gomath.Abs(point.Y)), gomath.Abs(point.Z))

	if maxCoordinate == gomath.Abs(point.X) {
		return math.CreateVector(point.X, 0.0, 0.0)
	} else if maxCoordinate == gomath.Abs(point.Y) {
		return math.CreateVector(0.0, point.Y, 0.0)
	}
	return math.CreateVector(0.0, 0.0, point.Z)
}

func (c *Cube) Bounds() *Bounds {
	return &Bounds{
		Minimum: math.CreatePoint(-1.0, -1.0, -1.0),
		Maximum: math.CreatePoint(1.0, 1.0, 1.0),
	}
}

func (c *Cube) ScaledBounds() *Bounds {
	return &Bounds{
		Minimum: c.Transform.MulT(math.CreatePoint(-1.0, -1.0, -1.0)),
		Maximum: c.Transform.MulT(math.CreatePoint(1.0, 1.0, 1.0)),
	}
}
