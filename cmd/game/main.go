package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Welcome to the Starside")
	game := &Game{
		Level: loadLevelFromFile("cmd/game/maps/level1.map"),
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Printf("error running game: %s", err)
	}
}
