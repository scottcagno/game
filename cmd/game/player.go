package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Entity
	img       *ebiten.Image
	direction int
	idle      bool
	speed     float64
}
