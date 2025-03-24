package render

import (
	gomath "math"
	"raygo/canvas"
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"raygo/scene"
)

func CreateMirrorScene(width int, height int) *canvas.Canvas {
	lightGray := math.CreateColor(0.9, 0.9, 0.9)
	black := math.CreateColor(0.0, 0.0, 0.0)
	red := math.CreateColor(0.9, 0.1, 0.1)
	burntUmber := math.CreateColor(math.BToF(110), math.BToF(38), math.BToF(14))

	grayBlackCheckerPattern := g.CreateCheckerPattern(lightGray, black)
	grayBlackCheckerPattern.SetTransform(math.Scaling(0.33, 0.33, 0.33))

	floor := g.CreatePlane()
	floor.GetMaterial().SetPattern(grayBlackCheckerPattern)
	floor.Material.SetReflective(0.3)

	mirror1 := g.CreateCube()
	mirror1.GetMaterial().SetReflective(1.0)
	mirror1.GetMaterial().SetDiffuse(0.1)
	mirror1.GetMaterial().SetShininess(300)
	mirror1.GetMaterial().SetTransparency(0.0)
	mirror1.SetTransform(math.Translation(0.0, 0.4, 5.0).MulM(math.Rotation_Y(gomath.Pi / 4.0)).MulM(math.Scaling(0.2, 0.2, 0.001)))

	mirror2 := g.CreateCube()
	mirror2.GetMaterial().SetReflective(1.0)
	mirror2.GetMaterial().SetDiffuse(0.1)
	mirror2.GetMaterial().SetShininess(300)
	mirror2.GetMaterial().SetTransparency(0.0)

	wall := g.CreateCube()
	wall.GetMaterial().SetColor(burntUmber)

	ball := g.CreateSphere()
	ball.GetMaterial().SetColor(red)
	ball.GetMaterial().SetReflective(0.5)
	ball.GetMaterial().SetDiffuse(0.5)
	ball.SetTransform(math.Translation(-2.0, 0.5, 5.0).MulM(math.Scaling(0.5, 0.5, 0.5)))

	objs := make([]g.Shape, 0)
	objs = append(objs, floor, mirror1, ball)
	// objs = append(objs, floor, wall, mirror1, mirror2, ball)
	light := lighting.CreateLight(math.CreatePoint(-9.0, 9.0, -1.0), math.CreateColor(1.0, 1.0, 1.0))
	w := scene.EmptyWorld()
	w.Light = &light
	w.Objects = objs

	cam := scene.CreateCamera(width, height, gomath.Pi/3.0)
	from := math.CreatePoint(0.0, 0.5, 0.0)
	to := math.CreatePoint(0.0, 0.5, 1.0)
	up := math.CreateVector(0.0, 1.0, 0.0)
	cam.SetTransform(math.ViewTransform(from, to, up))

	return cam.RenderMultithreaded(w, height/2)
}
