package render

import (
	gomath "math"
	"raygo/canvas"
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"raygo/scene"
)

func CreateRefractionScene(width int, height int) *canvas.Canvas {
	lightGray := math.CreateColor(0.9, 0.9, 0.9)
	black := math.CreateColor(0.0, 0.0, 0.0)

	grayBlackCheckerPattern := g.CreateCheckerPattern(lightGray, black)
	grayBlackCheckerPattern.SetTransform(math.Scaling(0.33, 0.33, 0.33))

	backdrop := g.CreatePlane()
	backdrop.SetTransform(math.Translation(0.0, 0.0, 5.0).
		MulM(math.Rotation_X(-gomath.Pi / 2.0)))
	backdropMat := g.DefaultMaterial()
	backdropMat.SetPattern(grayBlackCheckerPattern)
	backdropMat.SetSpecular(0.0)
	backdrop.SetMaterial(backdropMat)

	outerBall := g.CreateSphere()
	outerBall.SetTransform(math.Translation(0.0, 1.0, 0.5))
	outerMat := g.DefaultMaterial()
	outerMat.SetDiffuse(0.1)
	outerMat.SetReflective(0.3)
	outerMat.SetTransparency(1.0)
	outerMat.SetRefractiveIndex(1.5)
	outerBall.SetMaterial(outerMat)

	middleBall := g.CreateSphere()
	middleBall.SetTransform(math.Translation(0.0, 1.0, 0.5).MulM(math.Scaling(0.5, 0.5, 0.5)))
	middleMat := g.DefaultMaterial()
	middleMat.SetDiffuse(0.1)
	middleMat.SetReflective(0.3)
	middleMat.SetTransparency(1.0)
	middleMat.SetRefractiveIndex(1.00029)
	middleBall.SetMaterial(middleMat)

	objs := make([]g.Shape, 0)
	objs = append(objs, backdrop, outerBall, middleBall)
	light := lighting.CreateLight(math.CreatePoint(-10.0, 15.0, -5.0), math.CreateColor(0.5, 0.5, 0.5))
	w := scene.EmptyWorld()
	w.Light = &light
	w.Objects = objs

	cam := scene.CreateCamera(width, height, gomath.Pi/3.0)
	from := math.CreatePoint(0.0, 1.0, -3.5)
	to := math.CreatePoint(0.0, 1.0, 0.0)
	up := math.CreateVector(0.0, 1.0, 0.0)
	cam.Position = scene.CreateCameraPosition(from, to, up)

	return cam.RenderMultithreaded(w, 16)
}
