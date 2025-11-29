package main

import (
	"fmt"
	"raygo/canvas"
	"raygo/math"
	"raygo/obj"
	"raygo/render"
	"raygo/scene"
	"time"
)

func main() {
	// go tool pprof -http=":8000" ./cpu.pprof
	// f, err := os.Create("cpu2.pprof")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	RenderGif()
}

func RenderSingleImage() {
	begin := time.Now()
	image := render.CreateGradientPatternScene(400, 200)
	end := time.Now()

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())

	image.WritePng("gradient_shadoka.png")
}

func RenderMultipleImages() {
	begin := time.Now()
	anim := scene.CreateCameraAnimation(math.Radians(-90), 1, 3)
	images := render.CreateTeapotMultiframeScene(400, 200, anim)
	end := time.Now()

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())

	baseFilename := "teapot_multi"
	for i, image := range images {
		image.WritePPM(baseFilename + fmt.Sprintf("_%v.ppm", i))
	}
}

func RenderGif() {
	begin := time.Now()
	animDuration := 5.0
	anim := scene.CreateCameraAnimation(math.Radians(360), animDuration, 24)
	images := render.CreateTeapotMultiframeScene(400, 200, anim)
	end := time.Now()

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())

	canvas.WriteGif(images, animDuration, "teapot.gif")
}

func ReadOBJStats(path string) {
	begin := time.Now()
	teapot := obj.ParseFile(path)
	end := time.Now()
	diff := end.Sub(begin)

	fmt.Printf("parsing obj took %v seconds\n", diff.Seconds())
	teapot.PrintStats()
}
