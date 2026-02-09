package geometry

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"raygo/math"
)

type Texture struct {
	File string
	Data *image.Image
}

func (t *Texture) LoadTexture(directory string) {
	actualFile := fmt.Sprintf("%v%v", directory, t.File)
	fileReader, err := os.Open(actualFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fileReader.Close()

	i, _, err := image.Decode(fileReader)
	if err != nil {
		log.Fatal(err)
	}

	t.Data = &i
}

func (t *Texture) ColorAt(u float64, v float64) math.Color {
	bb := (*t.Data).Bounds()

	x := u * float64(bb.Size().X)
	y := v * float64(bb.Size().Y)

	// rgb values are premultiplied by the alpha value
	// no clue how to change that behaviour
	r, g, b, a := (*t.Data).At(int(x), int(y)).RGBA()
	r2 := float64(r) / float64(a)
	g2 := float64(g) / float64(a)
	b2 := float64(b) / float64(a)
	return math.CreateColor(r2, g2, b2)
}

func (t *Texture) Exists() bool {
	return t.File != ""
}
