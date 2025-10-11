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
	scene := render.CreateTeapotScene(1920, 1080)
	end := time.Now()

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())

	beginWrite := time.Now()
	scene.WriteFile("teapot_highres4.ppm")
	endWrite := time.Now()

	diffWrite := endWrite.Sub(beginWrite)
	fmt.Printf("writing file took %v seconds\n", diffWrite.Seconds())

	// begin := time.Now()
	// teapot := obj.ParseFile("resources/teapot_high.obj")
	// end := time.Now()
	// diff := end.Sub(begin)

	// fmt.Printf("parsing obj took %v seconds\n", diff.Seconds())
	// teapot.PrintStats()
}
