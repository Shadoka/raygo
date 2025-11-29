package canvas

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"raygo/math"
	"strconv"
	"strings"
)

type Canvas struct {
	Width, Height int
	Pixels        []math.Color
}

type ppmColor struct {
	r, g, b uint64
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
	var b strings.Builder
	b.Grow(12 * c.Width * c.Height) // 3*3 bytes for string encoding of color + 3*1 byte for blanks
	for y := range c.Height {
		currentRow := ""
		for x := range c.Width {
			tColor := mapToTrueColor(c.GetPixelAt(x, y))
			separator := ""
			if x != 0 {
				separator = " "
			}

			redValueString := strconv.FormatUint(tColor.r, 10)
			greenValueString := strconv.FormatUint(tColor.g, 10)
			blueValueString := strconv.FormatUint(tColor.b, 10)
			currentRowPlusRedLength := len(currentRow) + len(separator) + len(redValueString)
			if currentRowPlusRedLength > 70 {
				b.WriteString(currentRow + "\n")
				currentRow = redValueString + " " + greenValueString + " " + blueValueString
			} else if currentRowPlusRedLength+len(greenValueString)+1 > 70 { // +1 because of blank
				b.WriteString(currentRow + " " + redValueString + "\n")
				currentRow = greenValueString + " " + blueValueString
			} else if currentRowPlusRedLength+len(greenValueString)+len(blueValueString)+2 > 70 { // +2 because of blanks
				b.WriteString(currentRow + " " + redValueString + " " + greenValueString + "\n")
				currentRow = blueValueString
			} else {
				currentRow = currentRow + separator + redValueString + " " + greenValueString + " " + blueValueString
			}
		}
		b.WriteString(currentRow + "\n")
	}

	return b.String()
}

func (c *Canvas) WritePPM(location string) {
	fileContent := fmt.Sprintf("%v%v", c.CreatePPMHeader(), c.CreatePPMBody())

	err := os.WriteFile(location, []byte(fileContent), 0644)

	if err != nil {
		panic(err)
	}
}

func (c *Canvas) WritePng(location string) {
	img := c.CreateImage()

	f, err := os.Create(location)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (c *Canvas) CreateImage() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, c.Width, c.Height))

	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			rgbPixel := mapToTrueColor(c.GetPixelAt(x, y))
			img.Set(x, y, color.NRGBA{
				R: uint8(rgbPixel.r),
				G: uint8(rgbPixel.g),
				B: uint8(rgbPixel.b),
				A: 255,
			})
		}
	}

	return img
}

func mapToTrueColor(mColor math.Color) ppmColor {
	return ppmColor{
		r: math.ClampToByte(mColor.X * 255),
		g: math.ClampToByte(mColor.Y * 255),
		b: math.ClampToByte(mColor.Z * 255),
	}
}
