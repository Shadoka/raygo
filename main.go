package main

import (
	"fmt"
	"raygo/render"
	"time"
)

func main() {
	// f, err := os.Create("cpu.pprof")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	begin := time.Now()
	//
	scene := render.CreateTeapotScene(400, 200)
	end := time.Now()

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())

	beginWrite := time.Now()
	// high res, 400x200 => 5425s
	scene.WriteFile("teapot_highres.ppm")
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
