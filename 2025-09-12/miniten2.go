package main

import "github.com/eihigh/miniten"

var vy float64
var y float64

func main() {
	miniten.Run(draw)
}

func draw() {
	if miniten.IsClicked() {
		vy = -6
	}

	vy += 0.5
	y += vy

	if y > 360 {
		y = 360
	}

	miniten.DrawImage("gopher.png", 0, int(y))
}
