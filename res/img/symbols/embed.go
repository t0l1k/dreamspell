package symbols

import (
	"bytes"
	_ "embed"
	"image"
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

func GetSignPngs() (arr []image.Image) {
	signPngs := [][]byte{ClearSignSymbolPng, HiddenSignSymbolPng, PolarSymbolPng, EmptySymbolPng}
	for _, png := range signPngs {
		img, _, err := image.Decode(bytes.NewReader(png))
		if err != nil {
			panic(err)
		}
		arr = append(arr, img)
	}
	return arr
}
