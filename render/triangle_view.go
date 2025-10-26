package render

import (
	gomath "math"
	"raygo/canvas"
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"raygo/scene"
)

func CreateTriangleScene(width int, height int) *canvas.Canvas {
	lightGray := math.CreateColor(0.9, 0.9, 0.9)
	black := math.CreateColor(0.0, 0.0, 0.0)
	red := math.CreateColor(0.9, 0.1, 0.1)
	burntUmber := math.CreateColor(math.BToF(110), math.BToF(38), math.BToF(14))

	grayBlackCheckerPattern := g.CreateCheckerPattern(lightGray, black)
	grayBlackCheckerPattern.SetTransform(math.Scaling(0.33, 0.33, 0.33))

	umberRedStripePattern := g.CreateStripePattern(burntUmber, red)

	floor := g.CreatePlane()
	floor.GetMaterial().SetPattern(grayBlackCheckerPattern)
	floor.Material.SetReflective(0.3)

	wallBack := g.CreatePlane()
	wallBack.GetMaterial().SetPattern(umberRedStripePattern)
	wallBack.SetTransform(math.Translation(0.0, 0.0, 6.0).MulM(math.Rotation_X(gomath.Pi / 2.0)))

	p1 := math.CreatePoint(-1.0, 0.1, 0.0)
	p2 := math.CreatePoint(1.0, 0.1, 0.0)
	p3 := math.CreatePoint(0.0, 1.0, 1.0)
	triangle := g.CreateTriangle(p1, p2, p3)

	objs := make([]g.Shape, 0)
	objs = append(objs, floor, wallBack, triangle)
	light := lighting.CreateLight(math.CreatePoint(-9.0, 9.0, -1.0), math.CreateColor(1.0, 1.0, 1.0))
	w := scene.EmptyWorld()
	w.Light = &light
	w.Objects = objs

	cam := scene.CreateCamera(width, height, gomath.Pi/3.0)
	from := math.CreatePoint(0.0, 2.5, -3.0)
	to := math.CreatePoint(0.0, 0.5, 0.0)
	up := math.CreateVector(0.0, 1.0, 0.0)
	cam.Position = scene.CreateCameraPosition(from, to, up)

	return cam.RenderMultithreaded(w, height/2)
}
