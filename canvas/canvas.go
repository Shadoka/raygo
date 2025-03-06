package canvas

import (
	"fmt"
	"os"
	"raygo/math"
)

type Canvas struct {
	Width, Height int
	Pixels        []math.Color
}

type ppmColor struct {
	r, g, b uint8
}

func CreateCanvas(width int, height int) Canvas {
	px := make([]math.Color, width*height)

	for i := range px {
		blackColor := math.CreateColor(0.0, 0.0, 0.0)
		px[i] = blackColor
	}

	return Canvas{
		Width:  width,
		Height: height,
		Pixels: px,
	}
}

func (c *Canvas) GetPixelAt(x int, y int) math.Color {
	return c.Pixels[y*c.Width+x]
}

func (c *Canvas) WritePixel(x int, y int, color math.Color) {
	c.Pixels[y*c.Width+x] = color
}

func (c *Canvas) CreatePPMHeader() string {
	return fmt.Sprintf("P3\n%v %v\n255\n", c.Width, c.Height)
}

func (c *Canvas) CreatePPMBody() string {
	ppmBody := ""

	for y := range c.Height {
		currentRow := ""
		for x := range c.Width {
			tColor := mapToTrueColor(c.GetPixelAt(x, y))
			separator := ""
			if x != 0 {
				separator = " "
			}

			// TODO: THIS LOOKS ABYSMAL WTF
			if len(fmt.Sprintf("%v%v%d", currentRow, separator, tColor.r)) > 70 {
				ppmBody = fmt.Sprintf("%v%v\n", ppmBody, currentRow)
				currentRow = fmt.Sprintf("%v %v %v", tColor.r, tColor.g, tColor.b)
			} else if len(fmt.Sprintf("%v%v%v %v", currentRow, separator, tColor.r, tColor.g)) > 70 {
				ppmBody = fmt.Sprintf("%v%v %v\n", ppmBody, currentRow, tColor.r)
				currentRow = fmt.Sprintf("%v %v", tColor.g, tColor.b)
			} else if len(fmt.Sprintf("%v%v%v %v %v", currentRow, separator, tColor.r, tColor.g, tColor.b)) > 70 {
				ppmBody = fmt.Sprintf("%v%v %v %v\n", ppmBody, currentRow, tColor.r, tColor.g)
				currentRow = fmt.Sprintf("%v", tColor.b)
			} else {
				currentRow = fmt.Sprintf("%v%v%v %v %v", currentRow, separator, tColor.r, tColor.g, tColor.b)
			}
		}
		ppmBody = fmt.Sprintf("%v%v\n", ppmBody, currentRow)
	}

	return ppmBody
}

func (c *Canvas) WriteFile(location string) {
	fileContent := fmt.Sprintf("%v%v", c.CreatePPMHeader(), c.CreatePPMBody())
	err := os.WriteFile(location, []byte(fileContent), 0644)
	if err != nil {
		panic(err)
	}
}

func mapToTrueColor(mColor math.Color) ppmColor {
	return ppmColor{
		r: math.ClampToByte(mColor.X * 255),
		g: math.ClampToByte(mColor.Y * 255),
		b: math.ClampToByte(mColor.Z * 255),
	}
}
