package scene_tzolkin

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res/img/symbols"
	"github.com/t0l1k/dreamspell/res/img/tons"
	"github.com/t0l1k/eui"
)

type TzolkinKinIcon struct {
	Image          *ebiten.Image
	rect           *eui.Rect
	Dirty, Visible bool
	kin            *lib.Kin
}

func NewTzolkinKinIcon(kin *lib.Kin, rect []int) *TzolkinKinIcon {
	return &TzolkinKinIcon{
		kin:     kin,
		rect:    eui.NewRect(rect),
		Dirty:   true,
		Visible: true,
	}
}

func (i *TzolkinKinIcon) Layout() {
	w, h := i.rect.Size()
	if i.Image == nil {
		i.Image = ebiten.NewImage(w, h)
	} else {
		i.Image.Clear()
	}
	clrs := []color.Color{eui.Red, eui.White, eui.Blue, eui.Yellow}
	var bg, fg color.Color
	if i.kin.IsPga() {
		bg = eui.Green
		fg = clrs[3]
	} else {
		nr := i.kin.GetColor() - 1
		bg = clrs[nr]
		nr++
		if int(nr) >= len(clrs) {
			nr = 0
		}
		fg = clrs[nr]
	}
	i.Image.Fill(bg)
	ww := float64(w) / 3
	tonImg := eui.NewIcon(
		ebiten.NewImageFromImage(
			tons.GetTonPngs().Get(i.kin.GetTon())),
		[]int{int(ww / 2), 0, int(ww * 2), h / 2})
	tonImg.Draw(i.Image)
	kinNr := eui.NewLabel(strconv.Itoa(int(i.kin.GetNr())), []int{0, h / 2, w, h / 2}, bg, fg)
	kinNr.Draw(i.Image)

	if i.kin.IsClearSign() {
		clearImg := eui.NewIcon(ebiten.NewImageFromImage(symbols.GetSignPngs()[0]), []int{0, 0, int(ww), int(ww)})
		clearImg.Draw(i.Image)
	}
	if i.kin.IsHiddenSign() {
		hiddenImg := eui.NewIcon(ebiten.NewImageFromImage(symbols.GetSignPngs()[1]), []int{0, 0, int(ww), int(ww)})
		hiddenImg.Draw(i.Image)
	}
	if i.kin.IsPolar() {
		polarImg := eui.NewIcon(ebiten.NewImageFromImage(symbols.GetSignPngs()[2]), []int{w - int(ww), 0, int(ww), int(ww)})
		polarImg.Draw(i.Image)
	}

	eui.DrawRect(i.Image, eui.NewRect([]int{0, 0, w, h}), fg)
	i.Dirty = false
}

func (i *TzolkinKinIcon) Size() (int, int) {
	return i.rect.Size()
}

func (i *TzolkinKinIcon) Update(dt int) {}

func (i *TzolkinKinIcon) Draw(surface *ebiten.Image) {
	if i.Dirty {
		i.Layout()
	}
	if i.Visible {
		op := &ebiten.DrawImageOptions{}
		x, y := i.rect.Pos()
		op.GeoM.Translate(float64(x), float64(y))
		surface.DrawImage(i.Image, op)
	}
}

func (i *TzolkinKinIcon) Resize(rect []int) {
	i.rect = eui.NewRect(rect)
	i.Dirty = true
	i.Image = nil
}

func (i *TzolkinKinIcon) Close() {
	i.Image.Dispose()
}
