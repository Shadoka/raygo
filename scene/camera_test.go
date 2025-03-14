package scene

import (
	gomath "math"
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateCamera(t *testing.T) {
	hsize := 160
	vsize := 120
	fov := gomath.Pi / 2.0

	c := CreateCamera(hsize, vsize, fov)

	assert.Assert(t, c.Hsize == hsize)
	assert.Assert(t, c.Vsize == vsize)
	assert.Assert(t, c.FieldOfView == fov)
	assert.Assert(t, c.Transform.Equals(math.IdentityMatrix()))
}

func TestCreateCameraHorizontalCanvas(t *testing.T) {
	hsize := 200
	vsize := 125
	fov := gomath.Pi / 2.0

	c := CreateCamera(hsize, vsize, fov)

	assert.Assert(t, c.PixelSize == 0.01)
}

func TestCreateCameraVerticalCanvas(t *testing.T) {
	hsize := 125
	vsize := 200
	fov := gomath.Pi / 2.0

	c := CreateCamera(hsize, vsize, fov)

	assert.Assert(t, c.PixelSize == 0.01)
}

func TestRayForPixelCenter(t *testing.T) {
	c := CreateCamera(201, 101, gomath.Pi/2.0)
	expectedOrigin := math.CreatePoint(0.0, 0.0, 0.0)
	expectedDirection := math.CreateVector(0.0, 0.0, -1.0)

	r := c.RayForPixel(100, 50)

	assert.Assert(t, expectedOrigin.Equals(r.Origin))
	assert.Assert(t, expectedDirection.Equals(r.Direction))
}

func TestRayForPixelCorner(t *testing.T) {
	c := CreateCamera(201, 101, gomath.Pi/2.0)
	expectedOrigin := math.CreatePoint(0.0, 0.0, 0.0)
	expectedDirection := math.CreateVector(0.66519, 0.33259, -0.66851)

	r := c.RayForPixel(0, 0)

	assert.Assert(t, expectedOrigin.Equals(r.Origin))
	assert.Assert(t, expectedDirection.Equals(r.Direction))
}

func TestRayForPixelWithTransform(t *testing.T) {
	c := CreateCamera(201, 101, gomath.Pi/2.0)
	tf := math.Rotation_Y(gomath.Pi / 4.0).MulM(math.Translation(0.0, -2.0, 5.0))
	c.SetTransform(tf)

	expectedOrigin := math.CreatePoint(0.0, 2.0, -5.0)
	expectedDirection := math.CreateVector(gomath.Sqrt(2)/2.0, 0.0, -gomath.Sqrt(2)/2.0)

	r := c.RayForPixel(100, 50)

	assert.Assert(t, expectedOrigin.Equals(r.Origin))
	assert.Assert(t, expectedDirection.Equals(r.Direction))
}

func TestRender(t *testing.T) {
	w := DefaultWorld()
	c := CreateCamera(11.0, 11.0, gomath.Pi/2.0)
	from := math.CreatePoint(0.0, 0.0, -5.0)
	to := math.CreatePoint(0.0, 0.0, 0.0)
	up := math.CreateVector(0.0, 1.0, 0.0)
	c.SetTransform(math.ViewTransform(from, to, up))
	expectedColor := math.CreateColor(0.38066, 0.47583, 0.2855)

	canv := c.Render(w)

	assert.Assert(t, expectedColor.Equals(canv.GetPixelAt(5, 5)))
}
