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
	blue := math.CreateColor(0.1, 0.1, 0.9)
	darkGreen := math.CreateColor(convertColorFromByte(2), convertColorFromByte(48), convertColorFromByte(32))
	lightGreen := math.CreateColor(convertColorFromByte(80), convertColorFromByte(200), convertColorFromByte(120))
	burntUmber := math.CreateColor(math.BToF(110), math.BToF(38), math.BToF(14))

	grayBlackCheckerPattern := g.CreateCheckerPattern(lightGray, black)
	grayBlackCheckerPattern.SetTransform(math.Scaling(0.33, 0.33, 0.33))

	umberRedStripePattern := g.CreateStripePattern(burntUmber, red)

	lightGreenDarkGreenCheckerPattern := g.CreateCheckerPattern(lightGreen, darkGreen)

	floor := g.CreatePlane()
	floor.GetMaterial().SetPattern(grayBlackCheckerPattern)
	floor.Material.SetReflective(0.3)

	wallBack := g.CreatePlane()
	wallBack.GetMaterial().SetPattern(umberRedStripePattern)
	wallBack.SetTransform(math.Translation(0.0, 0.0, 6.0).MulM(math.Rotation_X(gomath.Pi / 2.0)))

	wallBehindCamera := g.CreatePlane()
	wallBehindCamera.GetMaterial().SetPattern(lightGreenDarkGreenCheckerPattern)
	wallBehindCamera.SetTransform(math.Translation(0.0, 0.0, -4.0).MulM(math.Rotation_X(gomath.Pi / 2.0)))

	mirror1 := g.CreateCube()
	mirror1.GetMaterial().SetReflective(1.0)
	mirror1.GetMaterial().SetDiffuse(0.0)
	mirror1.GetMaterial().SetAmbient(0.0)
	mirror1.GetMaterial().SetShininess(300)
	mirror1.GetMaterial().SetTransparency(0.0)
	mirror1.SetTransform(math.Translation(0.0, 0.5, 5.0).MulM(math.Rotation_Y(gomath.Pi / 4.0).MulM(math.Scaling(0.5, 0.5, 0.001))))

	group1 := g.EmptyGroup()
	mirror2 := g.CreateCube()
	mirror2.GetMaterial().SetReflective(1.0)
	mirror2.GetMaterial().SetDiffuse(0.0)
	mirror2.GetMaterial().SetAmbient(0.0)
	mirror2.GetMaterial().SetShininess(300)
	mirror2.GetMaterial().SetTransparency(0.0)
	mirror2.SetTransform(math.Translation(-3.0, 0.5, 5.0).MulM(math.Rotation_Y(-gomath.Pi / 4.0).MulM(math.Scaling(0.5, 0.5, 0.001))))
	group1.AddChild(mirror2)

	middleWall := g.CreateCube()
	middleWall.GetMaterial().SetColor(burntUmber)
	middleWall.SetTransform(math.Translation(-1.0, 0.0, 0.0).MulM(math.Scaling(0.001, 20.0, 3.0)))

	group2 := g.EmptyGroup()
	ball := g.CreateSphere()
	ball.GetMaterial().SetColor(blue)
	ball.GetMaterial().SetReflective(0.5)
	ball.GetMaterial().SetDiffuse(0.5)
	ball.SetTransform(math.Translation(-3.0, 0.5, 0.0).MulM(math.Scaling(0.5, 0.5, 0.5)))
	group2.AddChild(ball)

	objs := make([]g.Shape, 0)
	objs = append(objs, floor, mirror1, group1, group2, wallBack, middleWall, wallBehindCamera)
	light := lighting.CreateLight(math.CreatePoint(-9.0, 9.0, -1.0), math.CreateColor(1.0, 1.0, 1.0))
	w := scene.EmptyWorld()
	w.Light = &light
	w.Objects = objs

	cam := scene.CreateCamera(width, height, gomath.Pi/3.0)
	from := math.CreatePoint(0.0, 0.5, -3.0)
	to := math.CreatePoint(0.0, 0.5, 1.0)
	up := math.CreateVector(0.0, 1.0, 0.0)
	cam.SetTransform(math.ViewTransform(from, to, up))

	return cam.RenderMultithreaded(w, height/2)
}
