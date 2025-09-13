package main

import (
	"math/rand/v2"

	"github.com/eihigh/miniten"
)

var y float64
var vy float64

var frames int
var wallXs = []int{}
var wallYs = []int{}

func main() {
	miniten.Run(draw)
}

func draw() {
	miniten.DrawImage("sky.png", 0, 0)

	if miniten.IsClicked() {
		vy = -6
	}

	vy += 0.5
	y += vy

	if y > 300 {
		y = 300
	}

	miniten.DrawImage("gopher.png", 100, int(y))

	frames++

	if frames%100 == 0 {
		holeY := rand.N(200)

		wallXs = append(wallXs, 640)
		wallYs = append(wallYs, holeY-360)

		wallXs = append(wallXs, 640)
		wallYs = append(wallYs, holeY+140)
	}

	for i := range wallXs {
		wallXs[i] -= 3
	}

	for i, wallX := range wallXs {
		miniten.DrawImage("wall.png", wallX, wallYs[i])
	}
}
