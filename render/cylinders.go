package render

import (
	gomath "math"
	"raygo/canvas"
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"raygo/scene"
)

func CreateCylinderScene(width int, height int) *canvas.Canvas {
	lightGray := math.CreateColor(0.9, 0.9, 0.9)
	red := math.CreateColor(0.9, 0.1, 0.1)
	blue := math.CreateColor(0.1, 0.1, 0.9)
	green := math.CreateColor(0.1, 0.9, 0.1)
	burntUmber := math.CreateColor(math.BToF(110), math.BToF(38), math.BToF(14))

	brownGrayStripePattern := g.CreateStripePattern(lightGray, burntUmber)

	floor := g.CreatePlane()
	floor.GetMaterial().SetPattern(brownGrayStripePattern)

	truncatedCylinder := g.CreateCylinder()
	truncatedCylinder.Minimum = -1.0
	truncatedCylinder.Maximum = 1.0
	truncatedCylinder.GetMaterial().SetColor(red)
	truncatedCylinder.SetTransform(math.Scaling(0.5, 0.5, 0.5))

	closedCylinder := g.CreateCylinder()
	closedCylinder.Minimum = -1.0
	closedCylinder.Maximum = 1.0
	closedCylinder.Closed = true
	closedCylinder.GetMaterial().SetColor(blue)
	closedCylinder.SetTransform(math.Translation(1.0, 0.0, 2.0))
	closedCylinder.GetMaterial().SetDiffuse(0.1)
	closedCylinder.GetMaterial().SetReflective(1.0)

	closedCone := g.CreateCone()
	closedCone.Minimum = -1.0
	closedCone.Maximum = 0.0
	closedCone.Closed = true
	closedCone.GetMaterial().SetColor(green)
	closedCone.Material.SetReflective(0.3)
	closedCone.SetTransform(math.Translation(-1.5, 1.0, 2.0))

	objs := make([]g.Shape, 0)
	objs = append(objs, floor, truncatedCylinder, closedCylinder, closedCone)
	light := lighting.CreateLight(math.CreatePoint(-9.0, 9.0, -5.0), math.CreateColor(1.0, 1.0, 1.0))
	w := scene.EmptyWorld()
	w.Light = &light
	w.Objects = objs

	cam := scene.CreateCamera(width, height, gomath.Pi/3.0)
	from := math.CreatePoint(0.0, 3.0, -5.0)
	to := math.CreatePoint(0.0, 1.0, 0.0)
	up := math.CreateVector(0.0, 1.0, 0.0)
	cam.Position = scene.CreateCameraPosition(from, to, up)

	return cam.RenderMultithreaded(w, height/2)
}
