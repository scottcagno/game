package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// ref: https://code.rocket9labs.com/tslocum/carotidartillery/src/branch/master/game.go

var startButtons = []ebiten.StandardGamepadButton{
	ebiten.StandardGamepadButtonRightBottom,
	ebiten.StandardGamepadButtonRightRight,
	ebiten.StandardGamepadButtonRightLeft,
	ebiten.StandardGamepadButtonRightTop,
	ebiten.StandardGamepadButtonFrontTopLeft,
	ebiten.StandardGamepadButtonFrontTopRight,
	ebiten.StandardGamepadButtonFrontBottomLeft,
	ebiten.StandardGamepadButtonFrontBottomRight,
	ebiten.StandardGamepadButtonCenterLeft,
	ebiten.StandardGamepadButtonCenterRight,
	ebiten.StandardGamepadButtonLeftStick,
	ebiten.StandardGamepadButtonRightStick,
	ebiten.StandardGamepadButtonLeftBottom,
	ebiten.StandardGamepadButtonLeftRight,
	ebiten.StandardGamepadButtonLeftLeft,
	ebiten.StandardGamepadButtonLeftTop,
	ebiten.StandardGamepadButtonCenterCenter,
}

type gamepad struct {
	gamepadIDs            []ebiten.GamepadID
	gamepadIDsBuf         []ebiten.GamepadID
	activeGamepad         ebiten.GamepadID
	initialButtonReleased bool
}

func (g *gamepad) update(p1 *player) (float64, float64) {

	gamepadDeadZone := 0.1

	g.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(g.gamepadIDsBuf[:0])
	for _, id := range g.gamepadIDsBuf {
		log.Printf("gamepad connected: %d", id)
		g.gamepadIDs = append(g.gamepadIDs, id)
	}
	for i, id := range g.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			log.Printf("gamepad disconnected: %d", id)
			g.gamepadIDs = append(g.gamepadIDs[:i], g.gamepadIDs[i+1:]...)
		}

		if g.activeGamepad == -1 {
			for _, button := range startButtons {
				if ebiten.IsStandardGamepadButtonPressed(id, button) {
					log.Printf("gamepad activated: %d", id)
					g.activeGamepad = id
					// g.updateCursor()
					break
				}
			}
		}
	}

	pan := 0.05
	px, py := p1.xPos, p1.yPos
	if g.activeGamepad != -1 {
		h := ebiten.StandardGamepadAxisValue(g.activeGamepad, ebiten.StandardGamepadAxisLeftStickHorizontal)
		v := ebiten.StandardGamepadAxisValue(g.activeGamepad, ebiten.StandardGamepadAxisLeftStickVertical)
		if v < -gamepadDeadZone || v > gamepadDeadZone || h < -gamepadDeadZone || h > gamepadDeadZone {
			px += h * pan
			py += v * pan
		}
	}
	return px, py
}
