package scene

import (
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
)

type World struct {
	Objects []g.Shape
	Light   *lighting.Light
}

func CreateWorld(objs []g.Shape, l *lighting.Light) *World {
	return &World{
		Objects: objs,
		Light:   l,
	}
}

func EmptyWorld() *World {
	objects := make([]g.Shape, 0)
	return CreateWorld(objects, nil)
}

func DefaultWorld() *World {
	light := lighting.CreateLight(math.CreatePoint(-10.0, 10.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	objects := make([]g.Shape, 0)

	s1 := g.CreateSphere()
	m1 := g.DefaultMaterial()
	m1.SetColor(math.CreateColor(0.8, 1.0, 0.6))
	(&m1).Diffuse = 0.7
	(&m1).Specular = 0.2
	s1.SetMaterial(m1)
	objects = append(objects, s1)

	s2 := g.CreateSphere()
	transform := math.Scaling(0.5, 0.5, 0.5)
	s2.SetTransform(transform)
	objects = append(objects, s2)

	return CreateWorld(objects, &light)
}

func (w *World) Intersect(r g.Ray) []g.Intersection {
	xs := make([]g.Intersection, 0)
	for _, shape := range w.Objects {
		xs = append(xs, shape.Intersect(r)...)
	}
	g.SortIntersections(xs)
	return xs
}

func (w *World) ShadeHit(comp g.IntersectionComputations) math.Color {
	shadowed := w.IsShadowed(comp.OverPoint)

	return lighting.PhongLighting(comp.Object.GetMaterial(),
		comp.Object,
		*w.Light,
		comp.Point, comp.Eyev, comp.Normalv,
		shadowed)
}

func (w *World) ColorAt(r g.Ray) math.Color {
	xs := w.Intersect(r)
	hit := g.Hit(xs)

	color := math.CreateColor(0.0, 0.0, 0.0)
	if hit != nil {
		comps := hit.PrepareComputation(r)
		color = w.ShadeHit(comps)
	}
	return color
}

func (w *World) GetObject(index int) *g.Shape {
	return &w.Objects[index]
}

func (w *World) IsShadowed(p math.Point) bool {
	v := w.Light.Position.Subtract(p)
	distance := v.Magnitude()
	direction := v.Normalize()

	r := g.CreateRay(p, direction)
	xs := w.Intersect(r)

	h := g.Hit(xs)
	return h != nil && h.IntersectionAt < distance
}
