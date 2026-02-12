package geometry

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"raygo/math"
)

type Texture struct {
	File     string
	Cubemap  bool
	Data     *image.Image
	CubeInfo *CubeMapInfo
}

type Texel struct {
	U float64
	V float64
	F Face
}

type CubeMapInfo struct {
	FaceWidth   float64
	FaceHeight  float64
	TopStart    Texel
	BottomStart Texel
	LeftStart   Texel
	RightStart  Texel
	FrontStart  Texel
	BackStart   Texel
}

type Face int

const (
	TOP Face = iota
	BOTTOM
	LEFT
	RIGHT
	FRONT
	BACK
	UNDEFINED
)

func (t *Texture) InitTexture(directory string) {
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

	if t.Cubemap {
		t.initCubeMapInfo()
	}
}

func (t *Texture) initCubeMapInfo() {
	bb := (*t.Data).Bounds()
	faceWidth := float64(bb.Size().X / 4)
	faceHeight := float64(bb.Size().Y / 3)
	// top := Texel{
	// 	U: faceWidth,
	// 	V: 0.0,
	// 	F: TOP,
	// }
	// back := Texel{
	// 	U: 0.0,
	// 	V: faceHeight,
	// 	F: BACK,
	// }
	// left := Texel{
	// 	U: faceWidth,
	// 	V: faceHeight,
	// 	F: LEFT,
	// }
	// front := Texel{
	// 	U: faceWidth * 2,
	// 	V: faceHeight,
	// 	F: FRONT,
	// }
	// right := Texel{
	// 	U: faceWidth * 3,
	// 	V: faceHeight,
	// 	F: RIGHT,
	// }
	// bottom := Texel{
	// 	U: faceWidth,
	// 	V: faceHeight * 2,
	// 	F: BOTTOM,
	// }
	top := Texel{
		U: faceWidth,
		V: faceHeight,
		F: TOP,
	}
	back := Texel{
		U: 0.0,
		V: faceHeight * 2,
		F: BACK,
	}
	left := Texel{
		U: faceWidth,
		V: faceHeight * 2,
		F: LEFT,
	}
	front := Texel{
		U: faceWidth * 2,
		V: faceHeight * 2,
		F: FRONT,
	}
	right := Texel{
		U: faceWidth * 3,
		V: faceHeight * 2,
		F: RIGHT,
	}
	bottom := Texel{
		U: faceWidth,
		V: faceHeight * 3,
		F: BOTTOM,
	}
	info := CubeMapInfo{
		FaceWidth:   faceWidth,
		FaceHeight:  faceHeight,
		TopStart:    top,
		BottomStart: bottom,
		LeftStart:   left,
		RightStart:  right,
		FrontStart:  front,
		BackStart:   back,
	}
	t.CubeInfo = &info
}

func (t *Texture) ColorAt(texel Texel) math.Color {
	bb := (*t.Data).Bounds()

	var x, y float64
	if t.Cubemap {
		x, y = t.getCubeMapCoordinate(texel)
	} else {
		x = texel.U * float64(bb.Size().X)
		y = texel.V * float64(bb.Size().Y)
	}

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

func (t *Texture) getCubeMapCoordinate(texel Texel) (float64, float64) {
	var x, y float64

	var offsetTexel Texel
	switch texel.F {
	case TOP:
		offsetTexel = t.CubeInfo.TopStart
	case BOTTOM:
		offsetTexel = t.CubeInfo.BottomStart
	case LEFT:
		offsetTexel = t.CubeInfo.LeftStart
	case RIGHT:
		offsetTexel = t.CubeInfo.RightStart
	case FRONT:
		offsetTexel = t.CubeInfo.FrontStart
	case BACK:
		offsetTexel = t.CubeInfo.BackStart
	case UNDEFINED:
		log.Fatal("cannot get an offset texel for undefined face")
	}

	x = texel.U*float64(t.CubeInfo.FaceWidth) + offsetTexel.U
	y = texel.V*float64(t.CubeInfo.FaceHeight) + offsetTexel.V

	return x, y
}
