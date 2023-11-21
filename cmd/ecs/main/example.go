package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/scottcagno/game/cmd/ecs"
)

type AnimationExample struct {
	spriteRenderSystem *ecs.SpriteRenderSystem
	animationSystem    *ecs.AnimationSystem
}

func (a *AnimationExample) Draw(screen *ebiten.Image) {
	a.spriteRenderSystem.Draw(screen)
}

func (a *AnimationExample) Update() error {
	return a.animationSystem.Update()
}

func (a *AnimationExample) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
