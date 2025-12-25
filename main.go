package main

import (
	"fmt"
	"os"
	"raygo/app"
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

	app.Run(os.Args)
	// if fileFlagIndex := slices.Index(os.Args, "-f"); fileFlagIndex != -1 {
	// 	if len(os.Args) <= fileFlagIndex+1 {
	// 		panic("missing file path after -f flag")
	// 	}
	// 	data, err := os.ReadFile(os.Args[fileFlagIndex+1])
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	yml := parser.ParseYaml(string(data))
	// 	validationResult := yml.Validate()
	// 	validationResult = append(validationResult, parser.ValidateReferences(yml)...)
	// 	if len(validationResult) != 0 {
	// 		for i, vr := range validationResult {
	// 			fmt.Printf("%v. %v\n", i, vr.Error())
	// 		}
	// 	}
	// } else {
	// 	RenderGif()
	// }
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
