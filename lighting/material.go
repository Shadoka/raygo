package lighting

import (
	gomath "math"
	"raygo/math"
)

const EPSILON = 0.00001

type Material struct {
	Color     math.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func CreateMaterial(c math.Color, am float64, diff float64, spec float64, shin float64) Material {
	return Material{
		Color:     c,
		Ambient:   am,
		Diffuse:   diff,
		Specular:  spec,
		Shininess: shin,
	}
}

func DefaultMaterial() Material {
	c := math.CreateColor(1.0, 1.0, 1.0)
	return CreateMaterial(c, 0.1, 0.9, 0.9, 200.0)
}

func (m *Material) SetColor(c math.Color) {
	m.Color = c
}

func (m Material) Equals(other Material) bool {
	return m.Color.Equals(other.Color) &&
		floatEquals(m.Ambient, other.Ambient) &&
		floatEquals(m.Diffuse, other.Diffuse) &&
		floatEquals(m.Specular, other.Specular) &&
		floatEquals(m.Shininess, other.Shininess)
}

func floatEquals(a float64, b float64) bool {
	return gomath.Abs(a-b) < EPSILON
}
