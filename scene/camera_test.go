package scene

import (
	gomath "math"
	g "raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

const EPSILON = 0.00001

func floatEquals(a float64, b float64) bool {
	diff := gomath.Abs(a - b)
	// if a and b are Inf diff becomes NaN
	if gomath.IsNaN(diff) {
		return true
	}
	return diff < EPSILON
}

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

	assert.Assert(t, floatEquals(c.PixelSize, 0.01))
}

func TestCreateCameraVerticalCanvas(t *testing.T) {
	hsize := 125
	vsize := 200
	fov := gomath.Pi / 2.0

	c := CreateCamera(hsize, vsize, fov)

	assert.Assert(t, floatEquals(c.PixelSize, 0.01))
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

func TestInvestigateAcneBug(t *testing.T) {
	// this test was added after acne on the side of a cube appeared
	// cause was a bug in calculating the normal vector of a cube
	// the conversion from world to object space was missing
	lightGray := math.CreateColor(0.9, 0.9, 0.9)
	black := math.CreateColor(0.0, 0.0, 0.0)
	burntUmber := math.CreateColor(math.BToF(110), math.BToF(38), math.BToF(14))

	grayBlackCheckerPattern := g.CreateCheckerPattern(lightGray, black)
	grayBlackCheckerPattern.SetTransform(math.Scaling(0.33, 0.33, 0.33))

	room := g.CreateCube()
	room.SetTransform(math.Scaling(10.0, 10.0, 10.0))
	room.GetMaterial().SetPattern(grayBlackCheckerPattern)

	tableDownwardsTranslation := math.Translation(0.0, -6.0, 0.0)

	tableSurface := g.CreateCube()
	tableSurface.GetMaterial().SetColor(burntUmber)
	tableSurface.SetTransform(tableDownwardsTranslation.MulM(math.Scaling(3.0, 0.3, 2.0)))

	objs := make([]g.Shape, 0)
	objs = append(objs, room, tableSurface)
	light := lighting.CreateLight(math.CreatePoint(-9.0, 9.0, -5.0), math.CreateColor(0.7, 0.7, 0.7))
	w := EmptyWorld()
	w.Light = &light
	w.Objects = objs

	cam := CreateCamera(400, 200, gomath.Pi/3.0)
	from := math.CreatePoint(-9.0, -4.0, 0.0)
	to := math.CreatePoint(0.0, -8.0, 0.0)
	up := math.CreateVector(0.0, 1.0, 0.0)
	cam.SetTransform(math.ViewTransform(from, to, up))

	expected := math.CreateColor(0.12450, 0.04301, 0.01584)

	r := cam.RayForPixel(139, 68)
	actual := w.ColorAt(r, MAX_REFLECTION_LIMIT)

	assert.Assert(t, expected.Equals(actual))
}
