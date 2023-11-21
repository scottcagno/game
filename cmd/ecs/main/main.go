package main

import (
	"bytes"
	_ "embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/scottcagno/game/cmd/ecs"
)

// https://co0p.github.io/posts/ecs-animation/
// https://github.com/co0p/ebiten-ecs-animation/blob/master/cmd/animation/main.go

const screenWidth = 500
const screenHeight = 500

//go:embed spritesheet.png
var Spritesheet []byte

func main() {

	// loading assets
	frames := LoadSpritesheet(Spritesheet, 4, 250, 260)

	// entities
	registry := ecs.Registry{}

	e := registry.NewEntity()
	e.AddComponent(
		&ecs.AnimationComponent{
			Frames:            frames,
			CurrentFrameIndex: 0,
			Count:             0,
			AnimationSpeed:    0.125,
		},
	)
	e.AddComponent(&ecs.SpriteComponent{Image: frames[0]})
	e.AddComponent(&ecs.TransformComponent{PosX: 100, PosY: 100})

	// systems
	animationSystem := ecs.AnimationSystem{Registry: &registry}
	spriteRenderSystem := ecs.SpriteRenderSystem{Registry: &registry}

	// the game
	example := AnimationExample{
		spriteRenderSystem: &spriteRenderSystem,
		animationSystem:    &animationSystem,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("animation example")
	ebiten.RunGame(&example) // omitted error handling
}

// LoadSpritesheet returns n sub images from the given input image
func LoadSpritesheet(input []byte, n int, width int, height int) []*ebiten.Image {
	var sprites []*ebiten.Image

	spritesheet, _, err := image.Decode(bytes.NewReader(input))
	if err != nil {
		panic(err)
	}
	ebitenImage := ebiten.NewImageFromImage(spritesheet)

	for i := 0; i < n; i++ {
		dimensions := image.Rect(i*width, 0, (i+1)*width, height)
		sprite := ebitenImage.SubImage(dimensions).(*ebiten.Image)
		sprites = append(sprites, sprite)
	}

	return sprites
}
