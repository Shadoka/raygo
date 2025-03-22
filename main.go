package main

import (
	"fmt"
	"raygo/render"
	"time"
)

func main() {
	begin := time.Now()
	scene := render.CreateRefractionPlaygroundScene(1920, 1080)
	end := time.Now()
	scene.WriteFile("chapter11_playground.ppm")

	diff := end.Sub(begin)
	fmt.Printf("rendering took %v seconds\n", diff.Seconds())
}
