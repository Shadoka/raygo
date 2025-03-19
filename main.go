package main

import (
	"fmt"
	"raygo/render"
	"time"
)

func main() {
	begin := time.Now()
	scene := render.CreateGradientPatternScene(400, 200)
	end := time.Now()
	scene.WriteFile("chapter10_gradient_translated.ppm")

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())
}
