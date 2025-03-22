package main

import (
	"fmt"
	"raygo/render"
	"time"
)

func main() {
	begin := time.Now()
	scene := render.CreateRefractionScene(400, 200)
	end := time.Now()
	scene.WriteFile("chapter11_refractions.ppm")

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())
}
