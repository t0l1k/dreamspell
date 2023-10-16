package scene_moon

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res/img/seals"
	"github.com/t0l1k/dreamspell/res/img/symbols"
	"github.com/t0l1k/dreamspell/res/img/tons"
	"github.com/t0l1k/eui"
)

type KinIcon struct {
	Image          *ebiten.Image
	rect           *eui.Rect
	Dirty, Visible bool
	kin            *lib.Kin
	bg, fg         color.Color
}

func NewKinIcon(kin *lib.Kin, rect []int, bg, fg color.Color) *KinIcon {
	return &KinIcon{
		kin:     kin,
		bg:      bg,
		fg:      fg,
		rect:    eui.NewRect(rect),
		Dirty:   true,
		Visible: true,
	}
}

func (i *KinIcon) Layout() {
	w, h := i.rect.Size()
	if i.Image == nil {
		i.Image = ebiten.NewImage(w, h)
	} else {
		i.Image.Clear()
	}
	bg := i.bg
	if i.kin.IsPga() {
		bg = eui.Green
	}
	if i.kin.IsCentral() {
		bg = eui.Gray
	}
	i.Image.Fill(bg)
	sz4 := i.rect.GetLowestSize() / 4
	sz := sz4 * 2
	tonImg := eui.NewIcon(
		ebiten.NewImageFromImage(
			tons.GetTonPngs().Get(i.kin.GetTon())),
		[]int{sz4, sz4 / 6, sz, sz4})
	tonImg.Draw(i.Image)

	if i.kin.IsClearSign() {
		clearImg := eui.NewIcon(ebiten.NewImageFromImage(symbols.GetSignPngs()[0]), []int{0, 0, sz4, sz4})
		clearImg.Draw(i.Image)
	}
	if i.kin.IsHiddenSign() {
		hiddenImg := eui.NewIcon(ebiten.NewImageFromImage(symbols.GetSignPngs()[1]), []int{0, 0, sz4, sz4})
		hiddenImg.Draw(i.Image)
	}
	if i.kin.IsPolar() {
		polarImg := eui.NewIcon(ebiten.NewImageFromImage(symbols.GetSignPngs()[2]), []int{w - sz4, 0, sz4, sz4})
		polarImg.Draw(i.Image)
	}

	sealImg := eui.NewIcon(
		ebiten.NewImageFromImage(
			seals.GetSealPngs().Get(i.kin.GetSeal())),
		[]int{sz4, sz4, sz4 * 2, sz4 * 2})
	sealImg.Draw(i.Image)

	lblKinNr := eui.NewLabel(strconv.Itoa(i.kin.GetNr()), []int{sz4, sz4 * 3, sz, sz4}, bg, eui.Black)
	lblKinNr.Draw(i.Image)

	// eui.DrawRect(i.Image, eui.NewRect([]int{0, 0, w, h}), eui.Black)

	i.Dirty = false
}

func (i *KinIcon) Size() (int, int) {
	return i.rect.Size()
}

func (i *KinIcon) Update(dt int) {}

func (i *KinIcon) Draw(surface *ebiten.Image) {
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

func (i *KinIcon) Resize(rect []int) {
	i.rect = eui.NewRect(rect)
	i.Dirty = true
	i.Image = nil
}

func (i *KinIcon) Close() {
	i.Image.Dispose()
}
