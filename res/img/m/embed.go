package moon2

import (
	"bytes"
	_ "embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed moon03.png
	moonPng []byte
)

func GetMoonPng() image.Image {
	img, _, err := image.Decode(bytes.NewReader(moonPng))
	if err != nil {
		panic(err)
	}
	return img
}

func GetMoonPngs() []*ebiten.Image {
	var (
		imgs []*ebiten.Image
		sz   int = 47
		w, h int = sz, sz
		img  *ebiten.Image
		from *ebiten.Image = ebiten.NewImageFromImage(GetMoonPng())
		x, y int
	)
	for i := 0; i < 5; i++ {
		for j := 0; j < 7; j++ {
			img = ebiten.NewImage(w, h)
			op := &ebiten.DrawImageOptions{}
			x = j*w + j
			img.DrawImage(from.SubImage(image.Rect(x+1, y+1, x+w, y+h)).(*ebiten.Image), op)
			imgs = append(imgs, img)
		}
		y += sz + 1
	}
	return imgs
}
