package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type TileKind = byte

const (
	Blank TileKind = iota
	Obstacle
	Wall
	Floor
	ClosedDoor
	OpenedDoor
)

type Tile struct {
	Kind byte
}

type Level struct {
	Map    [][]Tile
	Player *Player
	Items  []Item
}

func NewLevel(w, h int, player *Player, items []Item) *Level {
	level := &Level{
		Map:    make([][]Tile, w),
		Player: player,
		Items:  items,
	}
	for i := range level.Map {
		level.Map[i] = make([]Tile, h)
	}
	return level
}

func (l *Level) Update() {

	// up
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		newPos := Pos{X: l.Player.X, Y: l.Player.Y - l.Player.Speed}
		if canMove(l, newPos) {
			l.Player.MovePlayer(newPos)
		}
	}
	// down
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		newPos := Pos{X: l.Player.X, Y: l.Player.Y + l.Player.Speed}
		if canMove(l, newPos) {
			l.Player.MovePlayer(newPos)
		}
	}
	// left
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		newPos := Pos{X: l.Player.X - l.Player.Speed, Y: l.Player.Y}
		if canMove(l, newPos) {
			l.Player.MovePlayer(newPos)
		}
	}
	// right
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		newPos := Pos{X: l.Player.X + l.Player.Speed, Y: l.Player.Y}
		if canMove(l, newPos) {
			l.Player.MovePlayer(newPos)
		}
	}
}

func (l *Level) Draw(screen *ebiten.Image) {
	l.Player.Draw(screen)
	for _, item := range l.Items {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(item.Pos.X), float64(item.Pos.Y))
		screen.DrawImage(item.img, op)
	}
}

func inRange(level *Level, pos Pos) bool {
	return int(pos.X) < len(level.Map[0]) && int(pos.Y) < len(level.Map) && pos.X >= 0 && pos.Y >= 0
}

func canMove(level *Level, pos Pos) bool {
	if !inRange(level, pos) {
		return false
	}
	// t := level.Map[pos.Y][pos.X]
	// switch t.Kind {
	// case Blank, Wall, Obstacle, ClosedDoor:
	// 	return false
	// default:
	// 	return true
	// }
	return true
}

func checkDoor(level *Level, pos Pos) {
	t := level.Map[pos.Y][pos.X]
	if t.Kind == ClosedDoor {
		level.Map[pos.Y][pos.X].Kind = OpenedDoor
	}
}
