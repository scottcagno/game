package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1080
	screenHeight = 720

	frameOX     = 0
	frameOY     = 48
	frameWidth  = 33
	frameHeight = 48
	frameCount  = 3
)

//go:embed character.png
var playerImg []byte

//go:embed tiles.png
var tilesImg []byte

var (
	p1Image *ebiten.Image
	p1      player
)

const (
	dirIdle = iota
	dirUp
	dirDown
	dirLeft
	dirRight
	dirUpRight
	dirUpLeft
	dirDownRight
	dirDownLeft
)

var dirMap = map[int]string{
	dirIdle:      "Idle",
	dirUp:        "Up",
	dirDown:      "Down",
	dirLeft:      "Left",
	dirRight:     "Right",
	dirUpRight:   "Up-Right",
	dirUpLeft:    "Up-Left",
	dirDownRight: "Down-Right",
	dirDownLeft:  "Down-Left",
}

// Create the player class
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	direction  int
	idle       bool
	speed      float64
}

func getDirection(x, y int) int {
	switch {
	case x == 1 && y == 0:
		return dirRight
	case x == -1 && y == 0:
		return dirLeft
	case x == 0 && y == -1:
		return dirUp
	case x == 0 && y == 1:
		return dirDown
	case x == 1 && y == -1:
		return dirUpRight
	case x == 1 && y == 1:
		return dirDownRight
	case x == -1 && y == -1:
		return dirUpLeft
	case x == -1 && y == 1:
		return dirDownLeft
	}
	return -1
}

// Move the player depending on which key is pressed
func updatePlayer() {
	h := ebiten.StandardGamepadAxisValue(0, ebiten.StandardGamepadAxisLeftStickHorizontal)
	v := ebiten.StandardGamepadAxisValue(0, ebiten.StandardGamepadAxisLeftStickVertical)
	if int(h) == 0 && int(v) == 0 {
		p1.idle = true
	} else {
		p1.idle = false
	}
	p1.xPos += float64(int(h) * int(p1.speed))
	p1.yPos += float64(int(v) * int(p1.speed))
	dir := getDirection(int(h), int(v))
	if dir != -1 {
		p1.direction = dir
	}
}

func drawPlayer(p1Op *ebiten.DrawImageOptions, tick int, screen *ebiten.Image) {
	switch p1.direction {
	case dirUp:
		up(p1Op, p1.idle, p1.image, tick, screen)
	case dirDown:
		down(p1Op, p1.idle, p1.image, tick, screen)
	case dirLeft:
		left(p1Op, p1.idle, p1.image, tick, screen)
	case dirRight:
		right(p1Op, p1.idle, p1.image, tick, screen)
	case dirUpRight:
		upright(p1Op, p1.idle, p1.image, tick, screen)
	case dirUpLeft:
		upleft(p1Op, p1.idle, p1.image, tick, screen)
	case dirDownRight:
		downright(p1Op, p1.idle, p1.image, tick, screen)
	case dirDownLeft:
		downleft(p1Op, p1.idle, p1.image, tick, screen)
	default:
		down(p1Op, true, p1.image, tick, screen)
	}
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadImgFile(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	errCheck(err)
	return img
}

func loadImgRaw(src []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(src))
	errCheck(err)
	return ebiten.NewImageFromImage(img)
}

func init() {
	// plyrImg_idle = loadImgRaw(resources.Player_png)
	// plyrImg_run = loadImgRaw(resources.Player_png)
	// plyrImg_jump = loadImgRaw(resources.Player_png)
	// plyrImg_attack = loadImgRaw(resources.Player_png)

	p1 = player{
		loadImgRaw(playerImg),
		screenWidth / 2.0,
		screenHeight / 2.0,
		0,
		true,
		2,
	}
}

type Game struct {
	count int
	op    *ebiten.DrawImageOptions
	p1    *player
	level *Level
}

func (g *Game) Update() error {
	g.count++
	updatePlayer()
	return nil
}

// func subImage(img *ebiten.Image, x0, y0, x1, y1 int) *ebiten.Image {
// 	return img.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
// }
//
// // todo: make a universal sprite rendering function.
// func sprite(tick int, op *ebiten.DrawImageOptions, idle bool, img *ebiten.Image, screen *ebiten.Image) {
// 	if idle {
// 		sx, sy := frameOX+1*frameWidth, frameOY
// 		screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
// 		return
// 	}
// 	i := (tick / 5) % frameCount
// 	sx, sy := frameOX+i*frameWidth, frameOY
// 	screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
// }

func left(op *ebiten.DrawImageOptions, idle bool, img *ebiten.Image, tick int, screen *ebiten.Image) {
	if idle {
		sx, sy := frameOX+1*frameWidth, frameOY
		screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
		return
	}
	i := (tick / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func right(op *ebiten.DrawImageOptions, idle bool, img *ebiten.Image, tick int, screen *ebiten.Image) {
	if idle {
		sx, sy := frameOX+7*frameWidth, frameOY
		screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
		return
	}
	i := (tick/5)%frameCount + 6
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func up(op *ebiten.DrawImageOptions, idle bool, img *ebiten.Image, tick int, screen *ebiten.Image) {
	if idle {
		sx, sy := frameOX+4*frameWidth, frameOY-frameOY
		screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
		return
	}
	i := (tick/5)%frameCount + 3
	sx, sy := frameOX+i*frameWidth, frameOY-frameOY
	screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func down(op *ebiten.DrawImageOptions, idle bool, img *ebiten.Image, tick int, screen *ebiten.Image) {
	if idle {
		sx, sy := frameOX+4*frameWidth, frameOY+frameOY
		screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
		return
	}

	i := (tick/5)%frameCount + 3
	sx, sy := frameOX+i*frameWidth, frameOY+frameOY
	screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func upright(op *ebiten.DrawImageOptions, idle bool, img *ebiten.Image, tick int, screen *ebiten.Image) {
	if idle {
		sx, sy := frameOX+7*frameWidth, frameOY-frameOY
		screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
		return
	}
	i := (tick/5)%frameCount + 6
	sx, sy := frameOX+i*frameWidth, frameOY-frameOY
	screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func upleft(op *ebiten.DrawImageOptions, idle bool, img *ebiten.Image, tick int, screen *ebiten.Image) {
	if idle {
		sx, sy := frameOX+1*frameWidth, frameOY-frameOY
		screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
		return
	}
	i := (tick / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY-frameOY
	screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func downright(op *ebiten.DrawImageOptions, idle bool, img *ebiten.Image, tick int, screen *ebiten.Image) {
	if idle {
		sx, sy := frameOX+7*frameWidth, frameOY+frameOY
		screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
		return
	}
	i := (tick/5)%frameCount + 6
	sx, sy := frameOX+i*frameWidth, frameOY+frameOY
	screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func downleft(op *ebiten.DrawImageOptions, idle bool, img *ebiten.Image, tick int, screen *ebiten.Image) {
	if idle {
		sx, sy := frameOX+1*frameWidth, frameOY+frameOY
		screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
		return
	}
	i := (tick / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY+frameOY
	screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.level.DrawLevel(screenWidth, screen)

	//screen.Fill(color.RGBA{0x39, 0x38, 0x52, 0xff})
	p1Op := &ebiten.DrawImageOptions{}
	p1Op.GeoM.Translate(p1.xPos, p1.yPos)
	// screen.DrawImage(p1.image, p1Op)
	// drawPlayer(p1Op, g.count, screen)

	switch p1.direction {
	case dirUp:
		up(p1Op, p1.idle, p1.image, g.count, screen)
	case dirDown:
		down(p1Op, p1.idle, p1.image, g.count, screen)
	case dirLeft:
		left(p1Op, p1.idle, p1.image, g.count, screen)
	case dirRight:
		right(p1Op, p1.idle, p1.image, g.count, screen)
	case dirUpRight:
		upright(p1Op, p1.idle, p1.image, g.count, screen)
	case dirUpLeft:
		upleft(p1Op, p1.idle, p1.image, g.count, screen)
	case dirDownRight:
		downright(p1Op, p1.idle, p1.image, g.count, screen)
	case dirDownLeft:
		downleft(p1Op, p1.idle, p1.image, g.count, screen)
	default:
		down(p1Op, true, p1.image, g.count, screen)
	}

	// Show the message
	msg := fmt.Sprintf("TPS: %0.2f\n", ebiten.ActualTPS())
	msg += fmt.Sprintf("Direction: %s Idle: %v\n", dirMap[p1.direction], p1.idle)
	// msg += fmt.Sprintf("X: %0.2f, Y: %0.2f\n")
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// return screenWidth / SCALE, screenHeight / SCALE
	return screenWidth, screenHeight
}

func main() {

	level := NewLevel(1, loadImgRaw(tilesImg), 32, tileFromFile)
	fmt.Println(level)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Upwell")
	if err := ebiten.RunGame(&Game{
		level: level,
	}); err != nil {
		panic(err)
	}
}
