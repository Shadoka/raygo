package render

import (
	"raygo/canvas"
	"raygo/math"
)

func CreateRedSquare(dimension int) *canvas.Canvas {
	cv := canvas.CreateCanvas(dimension, dimension)
	color := math.CreateColor(1.0, 0.0, 0.0)

	for y := range dimension {
		for x := range dimension {
			cv.WritePixel(x, y, color)
		}
	}

	return &cv
}
