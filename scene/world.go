package scene

import (
	gomath "math"
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
)

const MAX_REFLECTION_LIMIT = 4

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

func (w *World) ShadeHit(comp g.IntersectionComputations, remainingReflections int) math.Color {
	shadowed := w.IsShadowed(comp.OverPoint)

	surfaceColor := lighting.PhongLighting(*comp.Object.GetMaterial(),
		comp.Object,
		*w.Light,
		comp.OverPoint, comp.Eyev, comp.Normalv,
		shadowed)

	reflectedColor := w.ReflectedColor(comp, remainingReflections)
	refractedColor := w.RefractedColor(comp, remainingReflections)

	m := comp.Object.GetMaterial()
	if m.Reflective > 0.0 && m.Transparency > 0.0 {
		reflectance := comp.Schlick()
		return surfaceColor.
			Add(reflectedColor.Mul(reflectance)).
			Add(refractedColor.Mul(1.0 - reflectance))
	} else {
		return surfaceColor.Add(reflectedColor).Add(refractedColor)
	}
}

func (w *World) ColorAt(r g.Ray, remainingReflections int) math.Color {
	xs := w.Intersect(r)
	hit := g.Hit(xs)

	color := math.CreateColor(0.0, 0.0, 0.0)
	if hit != nil {
		comps := hit.PrepareComputation(r, xs)
		color = w.ShadeHit(comps, remainingReflections)
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

func (w *World) ReflectedColor(precomps g.IntersectionComputations, remainingReflections int) math.Color {
	if precomps.Object.GetMaterial().Reflective == 0.0 || remainingReflections <= 0 {
		return math.CreateColor(0.0, 0.0, 0.0)
	}

	reflectedRay := g.CreateRay(precomps.OverPoint, precomps.Reflectv)
	colorAtReflectionTarget := w.ColorAt(reflectedRay, remainingReflections-1)

	return colorAtReflectionTarget.Mul(precomps.Object.GetMaterial().Reflective)
}

func (w *World) RefractedColor(precomps g.IntersectionComputations, remainingRefractions int) math.Color {
	if precomps.Object.GetMaterial().Transparency == 0.0 || remainingRefractions <= 0 {
		return math.CreateColor(0.0, 0.0, 0.0)
	}

	// find the ratio of first index of refraction to the second
	refractionRatio := precomps.N1 / precomps.N2

	// cos(theta_i) is the same as the dot product of the two vectors
	cosI := precomps.Eyev.Dot(precomps.Normalv)

	// find sin(theta_t)^2 via trigonometric identity
	sin2T := (refractionRatio * refractionRatio) * (1.0 - cosI*cosI)

	if sin2T > 1.0 {
		// total internal reflection
		return math.CreateColor(0.0, 0.0, 0.0)
	}

	// find cos(theta_t) via trigonometric identity
	cosT := gomath.Sqrt(1.0 - sin2T)

	dirComp1 := precomps.Normalv.Mul(refractionRatio*cosI - cosT)
	dirComp2 := precomps.Eyev.Mul(refractionRatio)
	// compute the direction of the refracted ray
	direction := dirComp1.Subtract(dirComp2)

	// create the refracted ray
	refractRay := g.CreateRay(precomps.UnderPoint, direction)

	// find the color of the refracted ray, making sure to multiply
	// by the transparency value to account for any opacity
	color := w.ColorAt(refractRay, remainingRefractions-1).Mul(precomps.Object.GetMaterial().Transparency)

	return color
}
