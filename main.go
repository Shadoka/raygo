package main

import (
	"fmt"
	"raygo/render"
	"time"
)

func main() {
	begin := time.Now()
	scene := render.CreateMultiplePatternScene(400, 200)
	end := time.Now()
	scene.WriteFile("chapter10_patterns.ppm")

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())
}
