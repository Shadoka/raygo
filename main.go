package main

import (
	"fmt"
	"raygo/render"
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

	begin := time.Now()
	images := render.CreateTeapotMultiframeScene(400, 200)
	end := time.Now()

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())

	baseFilename := "teapot_multi"
	for i, image := range images {
		image.WriteFile(baseFilename + fmt.Sprintf("_%v.ppm", i))
	}

	// begin := time.Now()
	// teapot := obj.ParseFile("resources/teapot_high.obj")
	// end := time.Now()
	// diff := end.Sub(begin)

	// fmt.Printf("parsing obj took %v seconds\n", diff.Seconds())
	// teapot.PrintStats()
}
