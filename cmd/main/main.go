package main

import (
	"bytes"
	"fmt"
	"image"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/scottcagno/game/cmd/main/resources"
)

const (
	screenWidth  = 960
	screenHeight = 540
	unit         = 16
	groundY      = 380
)

var (
	leftSprite      *ebiten.Image
	rightSprite     *ebiten.Image
	idleSprite      *ebiten.Image
	backgroundImage *ebiten.Image
)

func init() {
	// preload right player image
	img, _, err := image.Decode(bytes.NewReader(resources.Right_png))
	if err != nil {
		panic(err)
	}
	rightSprite = ebiten.NewImageFromImage(img)

	// preload left player image
	img, _, err = image.Decode(bytes.NewReader(resources.Left_png))
	if err != nil {
		panic(err)
	}
	leftSprite = ebiten.NewImageFromImage(img)

	// preload center player image
	img, _, err = image.Decode(bytes.NewReader(resources.MainChar_png))
	if err != nil {
		panic(err)
	}
	idleSprite = ebiten.NewImageFromImage(img)

	// preload background image
	img, _, err = image.Decode(bytes.NewReader(resources.Background_png))
	if err != nil {
		panic(err)
	}
	backgroundImage = ebiten.NewImageFromImage(img)
}

type char struct {
	x, y, vx, vy int
}

func (c *char) tryJump() {
	// Now the character can jump anytime, even when the character is not on the ground.
	// If you want to restrict the character to jump only when it is on the ground, you can add an 'if' clause:
	//
	//     if gopher.y == groundY * unit {
	//         ...
	c.vy = -10 * unit
}

func (c *char) update() {
	c.x += c.vx
	c.y += c.vy
	if c.y > groundY*unit {
		c.y = groundY * unit
	}
	if c.vx > 0 {
		c.vx -= 4
	} else if c.vx < 0 {
		c.vx += 4
	}
	if c.vy < 20*unit {
		c.vy += 8
	}
}

func (c *char) draw(screen *ebiten.Image) {
	s := idleSprite
	switch {
	case c.vx > 0:
		s = rightSprite
	case c.vx < 0:
		s = leftSprite
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	screen.DrawImage(s, op)
}

type Game struct {
	gopher *char
}

func (g *Game) Update() error {
	if g.gopher == nil {
		g.gopher = &char{x: 50 * unit, y: groundY * unit}
	}

	// Controls
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.gopher.vx = -4 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.gopher.vx = 4 * unit
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.gopher.tryJump()
	}
	g.gopher.update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draws Background Image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(backgroundImage, op)

	// Draws the Gopher
	g.gopher.draw(screen)

	// Show the message
	msg := fmt.Sprintf("TPS: %0.2f\nPress the space key to jump.", ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Platformer (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
