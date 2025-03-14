package lighting

import (
	"raygo/math"
)

type Light struct {
	Position  math.Point
	Intensity math.Color
}

func CreateLight(p math.Point, intensity math.Color) Light {
	return Light{
		Position:  p,
		Intensity: intensity,
	}
}

func (l Light) Equals(other Light) bool {
	return l.Position.Equals(other.Position) &&
		l.Intensity.Equals(other.Intensity)
}
