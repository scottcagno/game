package resources

import (
	_ "embed"
)

var (
	//go:embed background.png
	Background_png []byte

	//go:embed left.png
	Left_png []byte

	//go:embed mainchar.png
	MainChar_png []byte

	//go:embed right.png
	Right_png []byte
)
