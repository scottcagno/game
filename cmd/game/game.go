package main

import (
	"math"
	"sync/atomic"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1080
	screenHeight = 720
)

type Pos struct {
	X, Y int
}

type Entity struct {
	Pos
}

type Game struct {
	tick  atomic.Int64
	Level *Level
}

func (g *Game) Tick() {
	if !g.tick.CompareAndSwap(math.MaxInt64-1, 0) {
		g.tick.Add(1)
	}
}

func (g *Game) GetTick() int {
	return int(g.tick.Load())
}

func (g *Game) Update() error {
	g.Tick()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Level.DrawLevel(screenWidth, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}
