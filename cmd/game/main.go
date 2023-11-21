package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/tileset.png
var tileset []byte

//go:embed assets/character.png
var character []byte

//go:embed assets/tiles.png
var tiles []byte

var (
	tilesetImg   *ebiten.Image
	characterImg *ebiten.Image
	tilesImg     *ebiten.Image
)

func init() {
	//	tilesetImg = loadImgRaw(tileset)
	//	characterImg = loadImgRaw(character)
	//	tilesImg = loadImgRaw(tiles)
}

func loadImgRaw(src []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(src))
	if err != nil {
		log.Printf("error decoding image: %s", err)
	}
	return ebiten.NewImageFromImage(img)
}

func main() {

	tilesImg, _, err := ebitenutil.NewImageFromFile("cmd/game/assets/tiles.png")
	if err != nil {
		panic(err)
	}
	characterImg, _, err := ebitenutil.NewImageFromFile("cmd/game/assets/character.png")
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Welcome to the Starside")
	game := &Game{
		Level: OpenLevel("cmd/game/maps/level1.map", tilesImg, &Player{
			Entity: Entity{
				Pos: Pos{0, 0},
			},
			img:       characterImg,
			direction: 0,
			idle:      false,
			speed:     0,
		}),
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Printf("error running game: %s", err)
	}
}
