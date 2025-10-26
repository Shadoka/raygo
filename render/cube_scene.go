package render

import (
	gomath "math"
	"raygo/canvas"
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"raygo/scene"
)

func CreateCubeScene(width int, height int) *canvas.Canvas {
	lightGray := math.CreateColor(0.9, 0.9, 0.9)
	black := math.CreateColor(0.0, 0.0, 0.0)
	burntUmber := math.CreateColor(math.BToF(110), math.BToF(38), math.BToF(14))

	grayBlackCheckerPattern := g.CreateCheckerPattern(lightGray, black)
	grayBlackCheckerPattern.SetTransform(math.Scaling(0.33, 0.33, 0.33))

	room := g.CreateCube()
	room.SetTransform(math.Scaling(10.0, 10.0, 10.0))
	room.GetMaterial().SetPattern(grayBlackCheckerPattern)

	tableDownwardsTranslation := math.Translation(0.0, -6.0, 0.0)
	tableLegScaling := math.Scaling(0.3, 2.0, 0.3)

	tableSurface := g.CreateCube()
	tableSurface.GetMaterial().SetColor(burntUmber)
	tableSurface.SetTransform(tableDownwardsTranslation.MulM(math.Scaling(3.0, 0.3, 2.0)))

	leg1 := g.CreateCube()
	leg1.GetMaterial().SetColor(burntUmber)
	leg1.SetTransform(tableDownwardsTranslation.MulM(math.Translation(-2.5, -1.8, -1.5)).MulM(tableLegScaling))

	leg2 := g.CreateCube()
	leg2.GetMaterial().SetColor(burntUmber)
	leg2.SetTransform(tableDownwardsTranslation.MulM(math.Translation(2.5, -1.8, -1.5)).MulM(tableLegScaling))

	leg3 := g.CreateCube()
	leg3.GetMaterial().SetColor(burntUmber)
	leg3.SetTransform(tableDownwardsTranslation.MulM(math.Translation(2.5, -1.8, 1.5)).MulM(tableLegScaling))

	leg4 := g.CreateCube()
	leg4.GetMaterial().SetColor(burntUmber)
	leg4.SetTransform(tableDownwardsTranslation.MulM(math.Translation(-2.5, -1.8, 1.5)).MulM(tableLegScaling))

	glassCube := g.CreateCube()
	glassCube.GetMaterial().SetTransparency(1.0)
	glassCube.GetMaterial().SetRefractiveIndex(1.5)
	glassCube.GetMaterial().SetDiffuse(0.1)
	glassCube.GetMaterial().SetReflective(0.3)
	glassCube.SetTransform(tableDownwardsTranslation.MulM(math.Translation(0.0, 0.83, 0.0)).MulM(math.Scaling(0.66, 0.66, 0.66)))

	mirror := g.CreateCube()
	mirror.GetMaterial().SetReflective(1.0)
	mirror.GetMaterial().SetDiffuse(0.1)
	mirror.GetMaterial().SetShininess(300)
	mirror.GetMaterial().SetTransparency(0.0)
	mirror.SetTransform(tableDownwardsTranslation.
		MulM(math.Translation(9.94, 0.0, 0.0)).
		MulM(math.Rotation_Y(gomath.Pi / 2.0)).
		MulM(math.Scaling(5.0, 2.0, 0.05)))

	objs := make([]g.Shape, 0)
	objs = append(objs, room, tableSurface, leg1, leg2, leg3, leg4, glassCube, mirror)
	light := lighting.CreateLight(math.CreatePoint(-9.0, 9.0, -5.0), math.CreateColor(0.7, 0.7, 0.7))
	w := scene.EmptyWorld()
	w.Light = &light
	w.Objects = objs

	cam := scene.CreateCamera(width, height, gomath.Pi/3.0)
	// from := math.CreatePoint(-7.0, 8.0, -6.0)
	// to := math.CreatePoint(0.0, -5.0, 0.0)
	// from := math.CreatePoint(-9.0, -4.0, 0.0)
	// to := math.CreatePoint(0.0, -8.0, 0.0)

	// line view up with top of table
	// from := math.CreatePoint(-9.0, -5.67, -6.0)
	// to := math.CreatePoint(0.0, -5.67, 0.0)

	// slightly angled from top
	from := math.CreatePoint(-9.0, -4.5, -6.0)
	to := math.CreatePoint(0.0, -5.67, 0.0)
	up := math.CreateVector(0.0, 1.0, 0.0)
	cam.Position = scene.CreateCameraPosition(from, to, up)

	return cam.RenderMultithreaded(w, height/2)
}
