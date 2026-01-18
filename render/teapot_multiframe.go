package render

import (
	gomath "math"
	"raygo/canvas"
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"raygo/obj"
	"raygo/scene"
)

func CreateTeapotMultiframeScene(width int, height int, animation *scene.CameraAnimation) []*canvas.Canvas {
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

	ceiling := g.CreatePlane()
	ceiling.GetMaterial().SetPattern(grayBlackCheckerPattern)
	ceiling.SetTransform(math.Translation(0.0, 40.0, 0.0))

	wallBack := g.CreatePlane()
	wallBack.GetMaterial().SetPattern(umberRedStripePattern)
	wallBack.SetTransform(math.Translation(0.0, 0.0, 40.0).MulM(math.Rotation_X(gomath.Pi / 2.0)))

	wallLeft := g.CreatePlane()
	wallLeft.GetMaterial().SetPattern(umberRedStripePattern)
	// order of rotations: first x, then y
	wallSideRotation := math.Rotation_Y(gomath.Pi / 2.0).MulM(math.Rotation_X(gomath.Pi / 2.0))
	wallLeft.SetTransform(math.Translation(-40.0, 0.0, 0.0).MulM(wallSideRotation))

	wallRight := g.CreatePlane()
	wallRight.GetMaterial().SetPattern(umberRedStripePattern)
	wallRight.SetTransform(math.Translation(40.0, 0.0, 0.0).MulM(wallSideRotation))

	wallBehindCamera := g.CreatePlane()
	wallBehindCamera.GetMaterial().SetPattern(umberRedStripePattern)
	wallBehindCamera.SetTransform(math.Translation(0.0, 0.0, -40.0).MulM(math.Rotation_X(gomath.Pi / 2.0)))

	teapot := obj.ParseFile("resources/teapot_high.obj")
	teapotGroup := teapot.ToGroup(true)
	teapotMaterial := g.DefaultMaterial()
	teapotMaterial.SetReflective(0)
	teapotGroup.SetMaterial(teapotMaterial)
	teapotGroup.SetTransform(math.Rotation_X(-gomath.Pi / 2))

	objs := make([]g.Shape, 0)
	objs = append(objs, floor, ceiling, wallBack, wallLeft, wallRight, teapotGroup, wallBehindCamera)
	light := lighting.CreateLight(math.CreatePoint(-3.0, 30.0, -20.0), math.CreateColor(1.0, 1.0, 1.0))
	w := scene.EmptyWorld()
	w.Light = &light
	w.Objects = objs

	cam := scene.CreateCamera(width, height, gomath.Pi/3.0)
	from := math.CreatePoint(0.0, 25.0, -30.0)
	to := g.GetCenter(teapotGroup)
	up := math.CreateVector(0.0, 1.0, 0.0)
	cam.Position = scene.CreateCameraPosition(from, to, up)

	cam.Animation = animation

	return cam.Render(w, true)
}
