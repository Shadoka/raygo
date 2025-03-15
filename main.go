package main

import "raygo/render"

func main() {
	scene := render.CreateSceneFromCamera(200, 100)
	scene.WriteFile("chapter8.ppm")
}
