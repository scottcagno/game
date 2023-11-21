package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// ComponentType defines the supported component types
type ComponentType string

// ComponentTyper returns the type of a component
type ComponentTyper interface {
	Type() ComponentType
}

const SpriteType ComponentType = "SPRITE"

// SpriteComponent holds a reference to an image to be drawn
type SpriteComponent struct {
	Image *ebiten.Image
}

func (t *SpriteComponent) Type() ComponentType {
	return SpriteType
}

const TransformType ComponentType = "TRANSFORM"

// TransformComponent describes the current position of an entity
type TransformComponent struct {
	PosX, PosY int
}

func (t *TransformComponent) Type() ComponentType {
	return TransformType
}

const AnimationType ComponentType = "ANIMATION"

// AnimationComponent holds the data necessary for the animation of frames based on the animation speed
type AnimationComponent struct {
	Frames            []*ebiten.Image
	CurrentFrameIndex int
	Count             float64
	AnimationSpeed    float64
}

func (a AnimationComponent) Type() ComponentType {
	return AnimationType
}
