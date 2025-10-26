package render

import (
	gomath "math"
	"raygo/canvas"
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"raygo/scene"
)

func convertColorFromByte(c int) float64 {
	return float64(c) / 255.0
}

func CreateReflectionScene(width int, height int) *canvas.Canvas {
	red := math.CreateColor(1.0, 0.0, 0.0)
	green := math.CreateColor(0.0, 1.0, 0.0)
	brightPurple := math.CreateColor(0.74901, 0.25098, 0.74901)
	byzantium := math.CreateColor(0.43921, 0.16078, 0.38823)
	lightGray := math.CreateColor(0.9, 0.9, 0.9)
	black := math.CreateColor(0.0, 0.0, 0.0)
	darkGreen := math.CreateColor(convertColorFromByte(2), convertColorFromByte(48), convertColorFromByte(32))
	lightGreen := math.CreateColor(convertColorFromByte(80), convertColorFromByte(200), convertColorFromByte(120))

	redGreenGradient := g.CreateGradientPattern(red, green)
	redGreenGradient.SetTransform(math.Translation(-1.0, 0.0, 0.0).MulM(math.Scaling(2.0, 2.0, 2.0)))

	whiteRedStripedPattern := g.CreateStripePattern(math.CreateColor(1.0, 1.0, 1.0), math.CreateColor(1.0, 0.0, 0.0))

	purpleRingPattern := g.CreateRingPattern(brightPurple, byzantium)
	purpleRingPattern.SetTransform(math.Rotation_Z(gomath.Pi / 4.0).
		MulM(math.Rotation_X(-gomath.Pi / 3.0)).
		MulM(math.Scaling(0.1, 0.1, 0.1)))

	grayBlackCheckerPattern := g.CreateCheckerPattern(lightGray, black)

	lightGreenDarkGreenCheckerPattern := g.CreateCheckerPattern(darkGreen, lightGreen)

	floor := g.CreatePlane()
	floorMat := g.DefaultMaterial()
	floorMat.SetPattern(grayBlackCheckerPattern)
	floorMat.SetSpecular(0.0)
	floorMat.SetReflective(0.3)
	floorMat.SetShininess(100.0)
	floor.SetMaterial(floorMat)

	backdrop := g.CreatePlane()
	backdrop.SetTransform(math.Translation(0.0, 0.0, 3.0).
		MulM(math.Rotation_X(-gomath.Pi / 2.0)))
	backdropMat := g.DefaultMaterial()
	backdropMat.SetPattern(lightGreenDarkGreenCheckerPattern)
	backdropMat.SetSpecular(0.0)
	backdropMat.SetReflective(0.3)
	backdropMat.SetShininess(100.0)
	backdrop.SetMaterial(backdropMat)

	middle := g.CreateSphere()
	middle.SetTransform(math.Translation(-0.5, 1.0, 0.5))
	middleMat := g.DefaultMaterial()
	middleMat.SetPattern(purpleRingPattern)
	middleMat.SetDiffuse(0.7)
	middleMat.SetSpecular(0.3)
	middle.SetMaterial(middleMat)

	right := g.CreateSphere()
	right.SetTransform(math.Translation(1.5, 0.5, -0.5).MulM(math.Scaling(0.5, 0.5, 0.5)))
	rightMat := g.DefaultMaterial()
	rightMat.SetPattern(redGreenGradient)
	rightMat.SetDiffuse(0.7)
	rightMat.SetSpecular(0.3)
	rightMat.SetReflective(0.1)
	right.SetMaterial(rightMat)

	left := g.CreateSphere()
	left.SetTransform(math.Translation(-1.5, 0.33, -0.75).MulM(math.Scaling(0.33, 0.33, 0.33)))
	leftMat := g.DefaultMaterial()
	leftMat.SetPattern(whiteRedStripedPattern)
	leftMat.SetDiffuse(0.7)
	leftMat.SetSpecular(0.3)
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
