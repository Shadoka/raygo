package canvas

import (
	"raygo/math"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateCanvas(t *testing.T) {
	c := CreateCanvas(10, 20)
	blackColor := math.CreateColor(0.0, 0.0, 0.0)

	assert.Assert(t, c.Width == 10)
	assert.Assert(t, c.Height == 20)
	assert.Assert(t, len(c.Pixels) == 10*20)

	for x := range 10 {
		for y := range 20 {
			currentPixel := c.GetPixelAt(x, y)
			assert.Assert(t, currentPixel.Equals(blackColor))
		}
	}
}

func TestWritePixel(t *testing.T) {
	c := CreateCanvas(10, 20)
	color := math.CreateColor(1.0, 0.0, 0.0)
	red := math.CreateColor(1.0, 0.0, 0.0)

	c.WritePixel(5, 5, color)

	assert.Assert(t, c.GetPixelAt(5, 5).Equals(color))

	color = math.CreateColor(0.0, 1.0, 0.0)
	c.WritePixel(7, 7, color)

	assert.Assert(t, c.GetPixelAt(7, 7).Equals(color))
	assert.Assert(t, c.GetPixelAt(5, 5).Equals(red))
}

func TestCreatePPMHeader(t *testing.T) {
	c := CreateCanvas(5, 3)
	expectedMagicString := "P3"
	expectedDimensions := "5 3"
	expectedMaximumColorValue := "255"

	ppmHeader := c.CreatePPMHeader()
	headerParts := strings.Split(ppmHeader, "\n")

	assert.Assert(t, len(headerParts) == 4)
	assert.Assert(t, headerParts[0] == expectedMagicString)
	assert.Assert(t, headerParts[1] == expectedDimensions)
	assert.Assert(t, headerParts[2] == expectedMaximumColorValue)
}

func TestCreatePPMBody(t *testing.T) {
	c := CreateCanvas(10, 2)
	color := math.CreateColor(1.0, 0.8, 0.6)
	expectedBody := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`

	for y := range 2 {
		for x := range 10 {
			c.WritePixel(x, y, color)
		}
	}

	ppmBody := c.CreatePPMBody()

	assert.Assert(t, ppmBody == expectedBody)
}
