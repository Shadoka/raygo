package main

import (
	"fmt"
	"raygo/render"
	"time"
)

func main() {
	begin := time.Now()
	scene := render.CreateSceneWithPlane(200, 100)
	end := time.Now()
	scene.WriteFile("chapter9_lowambient.ppm")

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())
}
