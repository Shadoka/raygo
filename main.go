package main

import (
	"fmt"
	"raygo/render"
	"time"
)

func main() {
	begin := time.Now()
	scene := render.CreateMirrorScene(400, 200)
	end := time.Now()

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())

	beginWrite := time.Now()
	scene.WriteFile("mirrors.ppm")
	endWrite := time.Now()

	diffWrite := endWrite.Sub(beginWrite)
	fmt.Printf("writing file took %v seconds\n", diffWrite.Seconds())
}
