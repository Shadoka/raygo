package geometry

import (
	gomath "math"
	"raygo/math"
)

type Material struct {
	Color      math.Color
	Pattern    Pattern
	Ambient    float64
	Diffuse    float64
	Specular   float64
	Shininess  float64
	Reflective float64
}

func CreateMaterial(c math.Color, p Pattern, am float64, diff float64, spec float64, shin float64, refl float64) Material {
	return Material{
		Color:      c,
		Pattern:    p,
		Ambient:    am,
		Diffuse:    diff,
		Specular:   spec,
		Shininess:  shin,
		Reflective: refl,
	}
}

func DefaultMaterial() Material {
	c := math.CreateColor(1.0, 1.0, 1.0)
	return CreateMaterial(c, nil, 0.1, 0.9, 0.9, 200.0, 0.0)
}

func (m *Material) SetColor(c math.Color) {
	m.Color = c
}

func (m *Material) SetPattern(p Pattern) {
	m.Pattern = p
}

func (m *Material) SetAmbient(a float64) {
	m.Ambient = a
}

func (m *Material) SetDiffuse(d float64) {
	m.Diffuse = d
}

func (m *Material) SetSpecular(s float64) {
	m.Specular = s
}

func (m *Material) SetReflective(r float64) {
	m.Reflective = r
}

func (m *Material) SetShininess(s float64) {
	m.Shininess = s
}

func (m Material) Equals(other Material) bool {
	patternEquals := true
	if (m.Pattern == nil && other.Pattern != nil) ||
		(m.Pattern != nil && other.Pattern == nil) {
		patternEquals = false
	} else if m.Pattern != nil && other.Pattern != nil {
		patternEquals = m.Pattern.Equals(other.Pattern)
	}
	return m.Color.Equals(other.Color) &&
		floatEquals(m.Ambient, other.Ambient) &&
		floatEquals(m.Diffuse, other.Diffuse) &&
		floatEquals(m.Specular, other.Specular) &&
		floatEquals(m.Shininess, other.Shininess) &&
		patternEquals
}

func floatEquals(a float64, b float64) bool {
	return gomath.Abs(a-b) < EPSILON
}
