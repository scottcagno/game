package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Img *ebiten.Image
	Pos
	Speed  float64
	IsIdle bool
}

func NewPlayer(img *ebiten.Image, x, y float64) *Player {
	return &Player{
		Img:   img,
		Pos:   Pos{X: x, Y: y},
		Speed: 3,
	}
}

func (p *Player) Update() {
	// up
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Pos.Y -= p.Speed
	}
	// down
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Pos.Y += p.Speed
	}
	// left
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Pos.X -= p.Speed
	}
	// right
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Pos.X += p.Speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Pos.X, p.Pos.Y)
	screen.DrawImage(p.Img, op)
}
