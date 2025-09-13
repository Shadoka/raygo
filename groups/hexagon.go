package groups

import (
	gomath "math"
	g "raygo/geometry"
	"raygo/math"
)

func hexagonCorner() *g.Sphere {
	corner := g.CreateSphere()
	tf := math.Translation(0.0, 0.0, -1.0)
	tf = tf.MulM(math.Scaling(0.25, 0.25, 0.25))
	corner.SetTransform(tf)
	return corner
}

func hexagonEdge() *g.Cylinder {
	edge := g.CreateCylinder()
	edge.Minimum = 0.0
	edge.Maximum = 1.0
	tf := math.Translation(0.0, 0.0, -1.0)
	tf = tf.MulM(math.Rotation_Y(-gomath.Pi / 6.0))
	tf = tf.MulM(math.Rotation_Z(-gomath.Pi / 2.0))
	tf = tf.MulM(math.Scaling(0.25, 1.0, 0.25))
	edge.SetTransform(tf)
	return edge
}

func HexagonSide() *g.Group {
	side := g.EmptyGroup()

	side.AddChild(hexagonCorner())
	side.AddChild(hexagonEdge())

	return side
}

func Hexagon() *g.Group {
	hex := g.EmptyGroup()

	for n := range 6 {
		side := HexagonSide()
		side.SetTransform(math.Rotation_Y(float64(n) * gomath.Pi / 3.0))
		hex.AddChild(side)
	}

	return hex
}
