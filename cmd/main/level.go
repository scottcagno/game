package main

import (
	"bufio"
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var atlasMap = map[byte]int{
	' ':  384,
	'\t': 384,
	'\n': 384,
	'#':  145,  // wall
	'.':  406,  // floor
	'|':  126,  // door
	'/':  122,  // door (open)
	'R':  3788, // enemy 1
	'S':  3812, // enemy 2
	'@':  821,  // player spawn
}
var tileFromFile = newTilesFromFile("cmd/main/levelFile.txt", atlasMap)

type Level struct {
	ID       int
	TileSize int
	TilesImg *ebiten.Image
	TilesMap [][]int
}

func (l *Level) String() string {
	var ss string
	for _, row := range l.TilesMap {
		for _, col := range row {
			ss += fmt.Sprintf(" %d ", col)
		}
		ss += fmt.Sprintf(" \n ")
	}
	return ss
}

func newTilesFromFile(levelFile string, atlas map[byte]int) [][]int {
	// open the file
	fd, err := os.Open(levelFile)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	// read the file
	scanner := bufio.NewScanner(fd)
	levelLines := make([]string, 0)
	longestRow := 0
	var index int
	for scanner.Scan() {
		levelLines = append(levelLines, scanner.Text())
		if len(levelLines[index]) > longestRow {
			longestRow = len(levelLines[index])
		}
		index++
	}
	// initialize tiles matrix
	tiles := make([][]int, len(levelLines))
	for i := range tiles {
		tiles[i] = make([]int, longestRow)
	}
	// let's fill out our tiles with the proper
	// id's using the provided atlas mappings
	for y := 0; y < len(tiles); y++ {
		var tile int
		var found bool
		for x, ch := range levelLines[y] {
			tile, found = atlas[byte(ch)]
			if !found {
				tile = 384
			}
			tiles[y][x] = tile
		}
	}

	return tiles
}

func NewLevel(id int, tilesImg *ebiten.Image, tilesSize int, tilesMap [][]int) *Level {
	return &Level{
		ID:       id,
		TileSize: tilesSize,
		TilesImg: tilesImg,
		TilesMap: tilesMap,
	}
}

func (l *Level) DrawLevel(screenWidth int, screen *ebiten.Image) {
	w := l.TilesImg.Bounds().Dx()
	tileXCount := w / l.TileSize

	// Draw each tile with each DrawImage call.
	// As the source images of all DrawImage calls are always same,
	// this rendering is done very efficiently.
	// For more detail, see https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Image.DrawImage
	xCount := screenWidth / l.TileSize
	for _, m := range l.TilesMap {
		for i, t := range m {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xCount)*l.TileSize), float64((i/xCount)*l.TileSize))

			sx := (t % tileXCount) * l.TileSize
			sy := (t / tileXCount) * l.TileSize
			screen.DrawImage(l.TilesImg.SubImage(image.Rect(sx, sy, sx+l.TileSize, sy+l.TileSize)).(*ebiten.Image), op)
		}
	}
}
