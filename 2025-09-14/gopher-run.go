package main

import (
	"math/rand/v2"

	"github.com/eihigh/miniten"
)

type player struct {
	x     float64
	y     float64
	vy    float64
	inAir bool
}

type barrel struct {
	x float64
	y float64
}

type enemy interface {
	move()
	hit() bool
	draw()
}

var (
	frames  = 0
	ground  = 300.0
	p       player
	enemies = []enemy{}
)

func newBarrel() *barrel {
	return &barrel{x: 640, y: ground - 50}
}

func (b *barrel) move() {
	b.x -= 2
}

func (b *barrel) hit() bool {
	cx := p.x + 30/2
	cy := p.y + 38/2

	return b.x < cx && cx < b.x+50 && b.y < cy && cy < b.y+50
}

func (b *barrel) draw() {
	miniten.DrawImage("barrel.png", int(b.x), int(b.y))
}

type ghost struct {
	x float64
	y float64
}

func newGhost() *ghost {
	return &ghost{x: 640, y: ground - 50}
}

func (g *ghost) move() {
	g.x -= 4
}

func (g *ghost) hit() bool {
	cx := p.x + 30/2
	cy := p.y + 38/2

	return g.x < cx && cx < g.x+50 && g.y < cy && cy < g.y+50
}

func (g *ghost) draw() {
	miniten.DrawImage("ghost.png", int(g.x), int(g.y))
}

func main() {
	p.y = ground - 38

	miniten.Run(draw)
}

func draw() {
	frames++

	miniten.DrawRect(0, int(ground)-10, 640, 360)

	if miniten.IsLeft() {
		p.x -= 3
	}
	if miniten.IsRight() {
		p.x += 3
	}

	if p.inAir {
		p.vy += 0.4
		p.y += p.vy

		if p.y > ground-38 {
			p.y = ground - 38
			p.inAir = false
		}
	} else {
		if miniten.IsUp() {
			p.vy -= 10
			p.inAir = true
		}
	}

	miniten.DrawImage("gopher.png", int(p.x), int(p.y))

	if frames%60 == 0 {
		switch rand.N(2) {
		case 0:
			enemies = append(enemies, newBarrel())
		case 1:
			enemies = append(enemies, newGhost())
		}
	}

	for _, enemy := range enemies {
		enemy.move()
	}

	for _, enemy := range enemies {
		if enemy.hit() {
			miniten.Println("hitting!!")
		}
	}

	for _, enemy := range enemies {
		enemy.draw()
	}
}
