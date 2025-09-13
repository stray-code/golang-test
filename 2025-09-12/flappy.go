package main

import (
	"embed"
	"math"
	"math/rand/v2"

	"github.com/eihigh/miniten"
)

var y float64
var vy float64

var frames int
var wallXs = []int{}
var wallYs = []int{}

var scene = "title"

var isPrevClicked = false
var isJustClicked = false

var highScore = 0

//go:embed *.png
var fsys embed.FS

func main() {
	miniten.Run(draw)
}

func draw() {
	isJustClicked = !isPrevClicked && miniten.IsClicked()
	isPrevClicked = miniten.IsClicked()

	miniten.DrawImageFS(fsys, "sky.png", 0, 0)

	switch scene {
	case "title":
		drawTitle()
	case "game":
		drawGame()
	case "gameOver":
		drawGameOver()
	}
}

func drawTitle() {
	miniten.DrawImageFS(fsys, "gopher.png", 100, int(y))
	miniten.Println("はねるgopherくんゲーム")
	miniten.Println("クリックでスタート")

	if isJustClicked {
		scene = "game"
	}
}

func drawGame() {
	if miniten.IsClicked() {
		vy = -6
	}

	vy += 0.5
	y += vy

	y = math.Max(y, 0)
	y = math.Min(y, 300)

	miniten.DrawImageFS(fsys, "gopher.png", 100, int(y))

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
		miniten.DrawImageFS(fsys, "wall.png", wallX, wallYs[i])
	}

	cx := 100 + 30/2
	cy := int(y) + 31/2

	for i := range wallXs {
		left := wallXs[i]
		right := left + 100
		top := wallYs[i]
		bottom := top + 360

		if left < cx && cx < right && top < cy && cy < bottom {
			scene = "gameOver"

			if frames > highScore {
				highScore = frames
			}
		}
	}

	miniten.Println("Score:", frames)
}

func drawGameOver() {
	miniten.DrawImageFS(fsys, "gopher.png", 100, int(y))

	for i := range wallXs {
		miniten.DrawImageFS(fsys, "wall.png", wallXs[i], wallYs[i])
	}

	miniten.Println("Game Over")
	miniten.Println("Score:", frames)
	miniten.Println("Hight Score:", highScore)

	if isJustClicked {
		scene = "title"

		y = 0.0
		vy = 0.0
		frames = 0
		wallXs = []int{}
		wallYs = []int{}
	}
}
