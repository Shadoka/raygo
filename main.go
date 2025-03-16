package main

import (
	"fmt"
	"raygo/render"
	"time"
)

func main() {
	begin := time.Now()
	scene := render.CreateSceneWithPlane(1000, 500)
	end := time.Now()
	scene.WriteFile("chapter9_lowambient_hq.ppm")

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())
}
