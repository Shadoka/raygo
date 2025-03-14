package scene

import (
	"raygo/lighting"
	"raygo/math"
	"raygo/ray"
)

type World struct {
	Objects []ray.Shape
	Light   *lighting.Light
}

func CreateWorld(objs []ray.Shape, l *lighting.Light) *World {
	return &World{
		Objects: objs,
		Light:   l,
	}
}

func EmptyWorld() *World {
	objects := make([]ray.Shape, 0)
	return CreateWorld(objects, nil)
}

func DefaultWorld() *World {
	light := lighting.CreateLight(math.CreatePoint(-10.0, 10.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	objects := make([]ray.Shape, 0)

	s1 := ray.CreateSphere()
	m1 := lighting.DefaultMaterial()
	m1.SetColor(math.CreateColor(0.8, 1.0, 0.6))
	(&m1).Diffuse = 0.7
	(&m1).Specular = 0.2
	s1.SetMaterial(m1)
	objects = append(objects, s1)

	s2 := ray.CreateSphere()
	transform := math.Scaling(0.5, 0.5, 0.5)
	s2.SetTransform(transform)
	objects = append(objects, s2)

	return CreateWorld(objects, &light)
}

func (w *World) Intersect(r ray.Ray) []ray.Intersection {
	xs := make([]ray.Intersection, 0)
	for _, shape := range w.Objects {
		xs = append(xs, shape.Intersect(r)...)
	}
	ray.SortIntersections(xs)
	return xs
}

func (w *World) ShadeHit(comp ray.IntersectionComputations) math.Color {
	return lighting.PhongLighting(comp.Object.GetMaterial(),
		*w.Light,
		comp.Point, comp.Eyev, comp.Normalv)
}

func (w *World) ColorAt(r ray.Ray) math.Color {
	xs := w.Intersect(r)
	hit := ray.Hit(xs)

	color := math.CreateColor(0.0, 0.0, 0.0)
	if hit != nil {
		comps := hit.PrepareComputation(r)
		color = w.ShadeHit(comps)
	}
	return color
}

func (w *World) GetObject(index int) *ray.Shape {
	return &w.Objects[index]
}
