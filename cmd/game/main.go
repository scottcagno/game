package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
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
	characterImg = loadImgRaw(character)
	tilesImg = loadImgRaw(tiles)
}

func loadImgRaw(src []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(src))
	if err != nil {
		log.Printf("error decoding image: %s", err)
	}
	return ebiten.NewImageFromImage(img)
}

func main() {

	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Welcome to the Starside")
	game := &Game{
		Level: OpenLevel("cmd/game/maps/level1.map", tilesImg, &Player{
			img: characterImg,
		}),
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Printf("error running game: %s", err)
	}
}
