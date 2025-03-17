package main

import (
	"fmt"
	"raygo/render"
	"time"
)

func main() {
	begin := time.Now()
	scene := render.CreateStripedPatternScene(200, 100)
	end := time.Now()
	scene.WriteFile("chapter10_stripedpattern.ppm")

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())
}
