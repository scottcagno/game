package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed mainchar.png
var playerImg []byte

type Pos struct {
	X, Y float64
}

type Game struct {
	tick         int
	screenWidth  int
	screenHeight int
	player       *Player
}

func (g *Game) DrawOverlay(screen *ebiten.Image) {
	// Show the message
	msg := fmt.Sprintf("TPS: %0.2f\n", ebiten.ActualTPS())
	msg += fmt.Sprintf("Player position: X: %0.2f, Y: %0.2f\n", g.player.Pos.X, g.player.Pos.Y)
	// msg += fmt.Sprintf("Player actions: idel=%v\n", g.player.IsIdle)
	// msg += fmt.Sprintf("Direction: %s Idle: %v\n", dirMap[p1.direction], p1.idle)
	// msg += fmt.Sprintf("X: %0.2f, Y: %0.2f\n")
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Tick() {
	if g.tick == math.MaxInt-1 {
		g.tick = 0
		return
	}
	g.tick++
}

func (g *Game) Update() error {
	g.Tick()
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	g.DrawOverlay(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}

func main() {

	g := Game{
		tick:         0,
		screenWidth:  800,
		screenHeight: 600,
	}

	img := ebiten.NewImage(32, 32)
	img.Fill(color.RGBA{0x39, 0x38, 0x25, 0xff})
	g.player = NewPlayer(img, float64(g.screenWidth/2), float64(g.screenHeight/2))

	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle("Example game")
	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}
