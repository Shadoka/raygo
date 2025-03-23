package canvas

import (
	"fmt"
	"os"
	"raygo/math"
	"strconv"
	"time"
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
	ppmBody := ""

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
				ppmBody = ppmBody + currentRow + "\n"
				currentRow = redValueString + " " + greenValueString + " " + blueValueString
			} else if currentRowPlusRedLength+len(greenValueString)+1 > 70 { // +1 because of blank
				ppmBody = ppmBody + currentRow + " " + redValueString + "\n"
				currentRow = greenValueString + " " + blueValueString
			} else if currentRowPlusRedLength+len(greenValueString)+len(blueValueString)+2 > 70 { // +2 because of blanks
				ppmBody = ppmBody + currentRow + " " + redValueString + " " + greenValueString + "\n"
				currentRow = blueValueString
			} else {
				currentRow = currentRow + separator + redValueString + " " + greenValueString + " " + blueValueString
			}
		}
		ppmBody = ppmBody + currentRow + "\n"
	}

	return ppmBody
}

func (c *Canvas) WriteFile(location string) {
	beginStringProcessing := time.Now()
	fileContent := fmt.Sprintf("%v%v", c.CreatePPMHeader(), c.CreatePPMBody())
	endStringProcessing := time.Now()
	diffStringProcessing := endStringProcessing.Sub(beginStringProcessing)
	fmt.Printf("string processing took %v seconds\n", diffStringProcessing.Seconds())
	beginDiskWrite := time.Now()
	err := os.WriteFile(location, []byte(fileContent), 0644)
	endDiskWrite := time.Now()
	diffDiskWrite := endDiskWrite.Sub(beginDiskWrite)
	fmt.Printf("writing to disk took %v seconds\n", diffDiskWrite.Seconds())
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
