package render

import (
	gomath "math"
	"raygo/canvas"
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"raygo/scene"
)

func CreateGradientPatternScene(width int, height int) *canvas.Canvas {
	red := math.CreateColor(1.0, 0.0, 0.0)
	green := math.CreateColor(0.0, 1.0, 0.0)
	blue := math.CreateColor(0.0, 0.0, 1.0)
	redGreenGradient := g.CreateGradientPattern(red, green)
	redGreenGradient.SetTransform(math.Translation(-1.0, 0.0, 0.0).MulM(math.Scaling(2.0, 2.0, 2.0)))

	redGreenHorizontalGradient := g.CreateGradientPattern(red, green)
	redGreenHorizontalGradient.SetTransform(math.Rotation_Z(gomath.Pi / 2.0).MulM(math.Translation(-1.0, 0.0, 0.0).MulM(math.Scaling(2.0, 2.0, 2.0))))

	redBlueTiltedGradient := g.CreateGradientPattern(red, blue)
	redBlueTiltedGradient.SetTransform(math.Rotation_Z(gomath.Pi / 4.0).MulM(math.Translation(-1.0, 0.0, 0.0).MulM(math.Scaling(2.0, 2.0, 2.0))))

	floor := g.CreatePlane()
	floorMat := g.DefaultMaterial()
	floorColor := math.CreateColor(1.0, 0.9, 0.9)
	floorMat.SetColor(floorColor)
	(&floorMat).Specular = 0.0
	floor.SetMaterial(floorMat)

	backdrop := g.CreatePlane()
	backdrop.SetTransform(math.Translation(0.0, 0.0, 3.0).
		MulM(math.Rotation_X(-gomath.Pi / 2.0)))
	backdropMat := g.DefaultMaterial()
	backdropColor := math.CreateColor(0.1, 0.1, 0.7)
	backdropMat.SetColor(backdropColor)
	(&backdropMat).Specular = 0.0
	backdrop.SetMaterial(backdropMat)

	middle := g.CreateSphere()
	middle.SetTransform(math.Translation(-0.5, 1.0, 0.5))
	middleMat := g.DefaultMaterial()
	middleMat.SetPattern(redGreenGradient)
	(&middleMat).Diffuse = 0.7
	(&middleMat).Specular = 0.3
	middle.SetMaterial(middleMat)

	right := g.CreateSphere()
	right.SetTransform(math.Translation(1.5, 0.5, -0.5).MulM(math.Scaling(0.5, 0.5, 0.5)))
	rightMat := g.DefaultMaterial()
	rightMat.SetPattern(redGreenHorizontalGradient)
	(&rightMat).Diffuse = 0.7
	(&rightMat).Specular = 0.3
	right.SetMaterial(rightMat)

	left := g.CreateSphere()
	left.SetTransform(math.Translation(-1.5, 0.33, -0.75).MulM(math.Scaling(0.33, 0.33, 0.33)))
	leftMat := g.DefaultMaterial()
	leftMat.SetPattern(redBlueTiltedGradient)
	(&leftMat).Diffuse = 0.7
	(&leftMat).Specular = 0.3
	left.SetMaterial(leftMat)

	objs := make([]g.Shape, 0)
	objs = append(objs, floor, backdrop, middle, right, left)
	light := lighting.CreateLight(math.CreatePoint(-10.0, 10.0, -10.0), math.CreateColor(1.0, 1.0, 1.0))
	w := scene.EmptyWorld()
	w.Light = &light
	w.Objects = objs

	cam := scene.CreateCamera(width, height, gomath.Pi/3.0)
	from := math.CreatePoint(0.0, 1.0, -5.0)
	to := math.CreatePoint(0.0, 1.0, 0.0)
	up := math.CreateVector(0.0, 1.0, 0.0)
	cam.Position = scene.CreateCameraPosition(from, to, up)

	return cam.RenderSinglethreaded(w)
}
