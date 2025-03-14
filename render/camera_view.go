package render

import (
	gomath "math"
	"raygo/canvas"
	"raygo/lighting"
	"raygo/math"
	"raygo/ray"
	"raygo/scene"
)

func CreateSceneFromCamera(width int, height int) *canvas.Canvas {
	floor := ray.CreateSphere()
	floor.SetTransform(math.Scaling(10.0, 0.01, 10.0))
	floorMat := lighting.DefaultMaterial()
	floorColor := math.CreateColor(1.0, 0.9, 0.9)
	floorMat.SetColor(floorColor)
	(&floorMat).Specular = 0.0
	floor.SetMaterial(floorMat)

	leftWall := ray.CreateSphere()
	leftWall.SetMaterial(floorMat)
	leftWallTf := math.Translation(0.0, 0.0, 5.0).
		MulM(math.Rotation_Y(-gomath.Pi / 4.0)).
		MulM(math.Rotation_X(gomath.Pi / 2.0)).
		MulM(math.Scaling(10.0, 0.01, 10.0))
	leftWall.SetTransform(leftWallTf)

	rightWall := ray.CreateSphere()
	rightWall.SetMaterial(floorMat)
	rightWallTf := math.Translation(0.0, 0.0, 5.0).
		MulM(math.Rotation_Y(gomath.Pi / 4.0)).
		MulM(math.Rotation_X(gomath.Pi / 2.0)).
		MulM(math.Scaling(10.0, 0.01, 10.0))
	rightWall.SetTransform(rightWallTf)

	middle := ray.CreateSphere()
	middle.SetTransform(math.Translation(-0.5, 1.0, 0.5))
	middleMat := lighting.DefaultMaterial()
	middleColor := math.CreateColor(0.1, 1.0, 0.5)
	middleMat.SetColor(middleColor)
	(&middleMat).Diffuse = 0.7
	(&middleMat).Specular = 0.3
	middle.SetMaterial(middleMat)

	right := ray.CreateSphere()
	right.SetTransform(math.Translation(1.5, 0.5, -0.5).MulM(math.Scaling(0.5, 0.5, 0.5)))
	rightMat := lighting.DefaultMaterial()
	rightColor := math.CreateColor(0.5, 1.0, 0.1)
	rightMat.SetColor(rightColor)
	(&rightMat).Diffuse = 0.7
	(&rightMat).Specular = 0.3
	right.SetMaterial(rightMat)

	left := ray.CreateSphere()
	left.SetTransform(math.Translation(-1.5, 0.33, -0.75).MulM(math.Scaling(0.33, 0.33, 0.33)))
	leftMat := lighting.DefaultMaterial()
	leftColor := math.CreateColor(1.0, 0.8, 0.1)
	leftMat.SetColor(leftColor)
	(&leftMat).Diffuse = 0.7
	(&leftMat).Specular = 0.3
	left.SetMaterial(leftMat)

	objs := make([]ray.Shape, 0)
	objs = append(objs, floor, rightWall, leftWall, middle, right, left)
	light := lighting.CreateLight(math.CreatePoint(-10.0, 10.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	w := scene.EmptyWorld()
	w.Light = &light
	w.Objects = objs

	cam := scene.CreateCamera(width, height, gomath.Pi/3.0)
	from := math.CreatePoint(0.0, 1.5, -5.0)
	to := math.CreatePoint(0.0, 1.0, 0.0)
	up := math.CreateVector(0.0, 1.0, 0.0)
	cam.SetTransform(math.ViewTransform(from, to, up))

	return cam.Render(w)
}
