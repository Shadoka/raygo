package main

import "raygo/scenes"

func main() {
	scene := scenes.CreateRedSquare(100)
	scene.WriteFile("red_square.ppm")
}
