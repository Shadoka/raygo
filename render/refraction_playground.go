package render

import (
	gomath "math"
	"raygo/canvas"
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"raygo/scene"
)

func CreateRefractionPlaygroundScene(width int, height int) *canvas.Canvas {
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

	leftBall := g.CreateSphere()
	leftBall.SetTransform(math.Translation(0.0, 1.0, 0.5))
	leftMat := g.DefaultMaterial()
	leftMat.SetDiffuse(0.1)
	leftMat.SetReflective(0.3)
	leftMat.SetTransparency(1.0)
	leftMat.SetRefractiveIndex(1.5)
	leftBall.SetMaterial(leftMat)

	// middleBall := g.CreateSphere()
	// middleBall.SetTransform(math.Translation(0.0, 1.0, 0.5).MulM(math.Scaling(0.5, 0.5, 0.5)))
	// middleMat := g.DefaultMaterial()
	// middleMat.SetDiffuse(0.1)
	// middleMat.SetReflective(0.3)
	// middleMat.SetTransparency(1.0)
	// middleMat.SetRefractiveIndex(1.00029)
	// middleBall.SetMaterial(middleMat)

	rightBall := g.CreateSphere()
	rightBall.SetTransform(math.Translation(0.0, 1.0, -1.5).MulM(math.Scaling(0.2, 0.2, 0.2)))
	rightMat := g.DefaultMaterial()
	rightMat.SetDiffuse(0.1)
	rightMat.SetReflective(0.3)
	rightMat.SetTransparency(1.0)
	rightMat.SetRefractiveIndex(1.33)
	rightBall.SetMaterial(rightMat)

	rightBall2 := g.CreateSphere()
	rightBall2.SetTransform(math.Translation(-1.0, 1.0, -1.5).MulM(math.Scaling(0.2, 0.2, 0.2)))
	rightMat2 := g.DefaultMaterial()
	rightMat2.SetDiffuse(0.1)
	rightMat2.SetReflective(0.3)
	rightMat2.SetTransparency(1.0)
	rightMat2.SetRefractiveIndex(1.33)
	rightBall2.SetMaterial(rightMat2)

	rightBall3 := g.CreateSphere()
	rightBall3.SetTransform(math.Translation(1.0, 1.0, -1.5).MulM(math.Scaling(0.2, 0.2, 0.2)))
	rightMat3 := g.DefaultMaterial()
	rightMat3.SetDiffuse(0.1)
	rightMat3.SetReflective(0.3)
	rightMat3.SetTransparency(1.0)
	rightMat3.SetRefractiveIndex(1.33)
	rightBall3.SetMaterial(rightMat3)

	objs := make([]g.Shape, 0)
	objs = append(objs, backdrop, leftBall, rightBall, rightBall2, rightBall3)
	light := lighting.CreateLight(math.CreatePoint(-10.0, 15.0, -5.0), math.CreateColor(0.7, 0.7, 0.7))
	w := scene.EmptyWorld()
	w.Light = &light
	w.Objects = objs

	cam := scene.CreateCamera(width, height, gomath.Pi/3.0)
	from := math.CreatePoint(0.0, 1.0, -3.5)
	to := math.CreatePoint(0.0, 1.0, 0.0)
	up := math.CreateVector(0.0, 1.0, 0.0)
	cam.SetTransform(math.ViewTransform(from, to, up))

	return cam.RenderMultithreaded(w, 8)
}
