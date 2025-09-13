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
	scene := render.CreateHexagonScene(1920, 1080)
	end := time.Now()

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())

	beginWrite := time.Now()
	scene.WriteFile("hexagon_transformed.ppm")
	endWrite := time.Now()

	diffWrite := endWrite.Sub(beginWrite)
	fmt.Printf("writing file took %v seconds\n", diffWrite.Seconds())
}
