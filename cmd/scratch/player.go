package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Img *ebiten.Image
	Pos
	Speed  int
	IsIdle bool
}

func NewPlayer(img *ebiten.Image, x, y int) *Player {
	return &Player{
		Img:   img,
		Pos:   Pos{X: x, Y: y},
		Speed: 3,
	}
}

func (p *Player) MovePlayer(newPos Pos) {
	p.Pos = newPos
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Pos.X), float64(p.Pos.Y))
	screen.DrawImage(p.Img, op)
}
