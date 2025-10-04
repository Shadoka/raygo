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
	scene := render.CreateTeapotScene(400, 200)
	end := time.Now()

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())

	beginWrite := time.Now()
	// high res, 400x200 => 5425s
	// high res, 400x200, amr => 1859s
	// low res, 400x200, after matrix refactor => 75s
	// low res, 400x200, amr, single thread => 236s
	// low res, 400x200, amr, single thread, inverse cache => 300s, IIRC
	scene.WriteFile("teapot_highres2.ppm")
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
