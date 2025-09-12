package main

import (
	"math/rand/v2"

	"github.com/eihigh/miniten"
)

var xs = []int{100, 200, 300}

func main() {
	miniten.Run(draw)
}

func draw() {
	if miniten.IsClicked() {
		xs = append(xs, rand.N(500))
	}

	for i := range xs {
		xs[i] += 6
	}

	for _, x := range xs {
		miniten.DrawImage("gopher.png", x, 0)
	}
}
