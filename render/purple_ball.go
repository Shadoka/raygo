package render

import (
	"raygo/canvas"
	"raygo/lighting"
	"raygo/math"
	"raygo/ray"
)

func CreatePurpleBall(dimension int) *canvas.Canvas {
	rayOrigin := math.CreatePoint(0.0, 0.0, -5.0)
	wallZ := 10.0
	wallSize := 7.0
	pixelSize := wallSize / float64(dimension)
	half := wallSize / 2.0

	canvas := canvas.CreateCanvas(dimension, dimension)
	shape := ray.CreateSphere()

	m := lighting.DefaultMaterial()
	m.SetColor(math.CreateColor(1.0, 0.2, 1.0))
	shape.SetMaterial(m)

	lightPosition := math.CreatePoint(-10.0, 10.0, -10.0)
	lightColor := math.CreateColor(1.0, 1.0, 1.0)
	light := lighting.CreateLight(lightPosition, lightColor)

	// for each row of pixels in the canvas
	for y := range dimension {
		// compute the world y coordinate (top = +half, bottom = -half)
		worldY := half - pixelSize*float64(y)

		// for each pixel in the row
		for x := range dimension {
			// compute the world x coordinate (left = -half, right = +half)
			worldX := -half + pixelSize*float64(x)

			// describe the point on the wall that the ray will target
			position := math.CreatePoint(worldX, worldY, wallZ)

			r := ray.CreateRay(rayOrigin, position.Subtract(rayOrigin).Normalize())
			xs := shape.Intersect(r)

			if ray.Hit(xs) != nil {
				inters := ray.Hit(xs)
				point := r.Position(inters.IntersectionAt)
				normalv := inters.Object.NormalAt(point)
				eyev := r.Direction.Negate()
				color := lighting.PhongLighting(inters.Object.GetMaterial(), light, point, eyev, normalv, false)
				canvas.WritePixel(x, y, color)
			}
		}
	}

	return &canvas
}
