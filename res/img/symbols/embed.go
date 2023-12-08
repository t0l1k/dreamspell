package symbols

import (
	"bytes"
	_ "embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed  star.png
	ClearSignSymbolPng []byte
	//go:embed  hidden.png
	HiddenSignSymbolPng []byte
	//go:embed  romb.png
	PolarSymbolPng []byte
	//go:embed  empty.png
	EmptySymbolPng []byte
)

func GetSignPngs() (arr []*ebiten.Image) {
	signPngs := [][]byte{ClearSignSymbolPng, HiddenSignSymbolPng, PolarSymbolPng, EmptySymbolPng}
	for _, png := range signPngs {
		img, _, err := image.Decode(bytes.NewReader(png))
		if err != nil {
			panic(err)
		}
		im := ebiten.NewImageFromImage(img)
		arr = append(arr, im)
	}
	return arr
}
