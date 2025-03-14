package render

import (
	"raygo/canvas"
	"raygo/math"
	"raygo/ray"
)

func CreateRedSphere(dimension int) *canvas.Canvas {
	rayOrigin := math.CreatePoint(0.0, 0.0, -5.0)
	wallZ := 10.0
	wallSize := 7.0
	pixelSize := wallSize / float64(dimension)
	half := wallSize / 2.0

	canvas := canvas.CreateCanvas(dimension, dimension)
	color := math.CreateColor(1.0, 0.0, 0.0)
	shape := ray.CreateSphere()

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
				canvas.WritePixel(x, y, color)
			}
		}
	}

	return &canvas
}
