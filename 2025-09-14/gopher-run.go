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

type eagle struct {
	x      float64
	y      float64
	frames int
}

func newEagle() *eagle {
	return &eagle{x: 640, y: ground - 150}
}

func (e *eagle) move() {
	e.x -= 5
	e.frames++

	if e.frames < 60 {
		e.y += 1.5
	} else if e.frames < 120 {
		e.y -= 1.5
	} else {
		e.frames = 0
	}
}

func (e *eagle) hit() bool {
	cx := p.x + 30/2
	cy := p.y + 38/2

	return e.x < cx && cx < e.x+50 && e.y < cy && cy < e.y+50
}

func (e *eagle) draw() {
	miniten.DrawImage("eagle.png", int(e.x), int(e.y))
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
		switch rand.N(3) {
		case 0:
			enemies = append(enemies, newBarrel())
		case 1:
			enemies = append(enemies, newGhost())
		case 2:
			enemies = append(enemies, newEagle())
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
