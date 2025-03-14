package scene

import (
	gomath "math"
	"raygo/canvas"
	"raygo/math"
	"raygo/ray"
)

type Camera struct {
	Hsize       int
	Vsize       int
	FieldOfView float64
	Transform   math.Matrix
	HalfWidth   float64
	HalfHeight  float64
	PixelSize   float64
}

func CreateCamera(hsize int, vsize int, fov float64) *Camera {
	c := &Camera{
		Hsize:       hsize,
		Vsize:       vsize,
		FieldOfView: fov,
		Transform:   math.IdentityMatrix(),
	}
	c.calculateCameraProperties()

	return c
}

func (c *Camera) calculateCameraProperties() {
	halfView := gomath.Tan(c.FieldOfView / 2.0)
	aspect := float64(c.Hsize) / float64(c.Vsize)

	if aspect >= 1.0 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}

	c.PixelSize = (c.HalfWidth * 2.0) / float64(c.Hsize)
}

func (c *Camera) RayForPixel(x int, y int) ray.Ray {
	// the offset from the edge of the canvas to the pixels center
	xOffset := (float64(x) + 0.5) * c.PixelSize
	yOffset := (float64(y) + 0.5) * c.PixelSize

	// the untransformed coordinates of the pixel in world space
	// (remember that the camera looks toward -z, so +x is to the *left*)
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset

	// using the camera matrix, transform the canvas point and the origin,
	// and then compute the rays direction vector.
	// (remember that the canvas is at z = -1)
	pixel := c.Transform.Inverse().MulT(math.CreatePoint(worldX, worldY, -1.0))
	origin := c.Transform.Inverse().MulT(math.CreatePoint(0.0, 0.0, 0.0))
	direction := pixel.Subtract(origin).Normalize()

	return ray.CreateRay(origin, direction)
}

func (c *Camera) SetTransform(tf math.Matrix) {
	c.Transform = tf
}

func (c *Camera) Render(w *World) *canvas.Canvas {
	canv := canvas.CreateCanvas(c.Hsize, c.Vsize)

	for y := range c.Vsize {
		for x := range c.Hsize {
			r := c.RayForPixel(x, y)
			color := w.ColorAt(r)
			canv.WritePixel(x, y, color)
		}
	}

	return &canv
}
