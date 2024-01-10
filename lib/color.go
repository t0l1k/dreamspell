package lib

import (
	"image/color"

	"github.com/t0l1k/eui"
)

type SealColor int

const (
	RED SealColor = iota + 1
	WHITE
	BLUE
	YELLOW
)

func (c SealColor) String() string {
	arr := []string{
		"RED",
		"WHITE",
		"BLUE",
		"YELLOW"}
	return arr[int(int(c)-1)]
}

func (c SealColor) Color() (bg color.RGBA, fg color.RGBA) {
	var (
		red    = color.RGBA{255, 0, 0, 255}
		white  = color.RGBA{255, 255, 255, 255}
		blue   = color.RGBA{0, 0, 255, 255}
		yellow = color.RGBA{255, 255, 0, 255}
	)
	clrs := []color.RGBA{red, white, blue, yellow}
	if c == RED || c == BLUE {
		fg = white
	} else {
		fg = eui.Black
	}
	bg = clrs[int(int(c)-1)]
	return bg, fg
}
