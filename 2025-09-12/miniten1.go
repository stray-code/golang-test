package main

import "github.com/eihigh/miniten"

func main() {
	miniten.Run(draw)
}

func draw() {
	miniten.Println("hello world")
	miniten.Println("good morning")

	miniten.DrawImage("gopher.png", 100, 100)
	miniten.DrawCircle(200, 200, 100)
	miniten.DrawRect(50, 50, 10, 10)
}
