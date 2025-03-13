package main

import "raygo/scenes"

func main() {
	scene := scenes.CreateRedSphere(100)
	scene.WriteFile("red_sphere.ppm")
}
