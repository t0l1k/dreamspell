package scene_moon

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/eui"
)

type MoonIcon struct {
	Image          *ebiten.Image
	rect           *eui.Rect
	Dirty, Visible bool
	moon           int
	moonKin        *lib.Kin
}

func NewMoonIcon(moon int, moonKin *lib.Kin, rect []int) *MoonIcon {
	m := &MoonIcon{
		moon:    moon,
		moonKin: moonKin,
		rect:    eui.NewRect(rect),
		Dirty:   true,
		Visible: true,
	}
	m.setup(moon, moonKin)
	return m
}

func (i *MoonIcon) setup(moon int, kin *lib.Kin) {
	i.moon = moon
	i.moonKin = kin
	i.Dirty = true
}

func (i *MoonIcon) Layout() {
	w, h := i.rect.Size()
	if i.Image == nil {
		i.Image = ebiten.NewImage(w, h)
	} else {
		i.Image.Clear()
	}

	bg := eui.White
	fg := eui.Black
	i.Image.Fill(bg)

	sz := i.rect.GetLowestSize() / 6
	kinIcon := NewKinIcon(i.moonKin, []int{sz / 2, 0, sz * 5, sz * 5}, bg, fg)
	kinIcon.Draw(i.Image)

	lblMoonKin := eui.NewLabel("Кин луны", []int{sz / 2, sz * 5, sz * 5, sz}, bg, fg)
	lblMoonKin.Draw(i.Image)

	eui.DrawRect(i.Image, eui.NewRect([]int{0, 0, w, h}), eui.Black)

}

func (i *MoonIcon) Size() (int, int) {
	return i.rect.Size()
}

func (i *MoonIcon) Update(dt int) {}

func (i *MoonIcon) Draw(surface *ebiten.Image) {
	if i.Dirty {
		i.Layout()
		i.Dirty = false
	}
	if i.Visible {
		op := &ebiten.DrawImageOptions{}
		x, y := i.rect.Pos()
		op.GeoM.Translate(float64(x), float64(y))
		surface.DrawImage(i.Image, op)
	}
}

func (i *MoonIcon) Resize(rect []int) {
	i.rect = eui.NewRect(rect)
	i.Dirty = true
	i.Image = nil
}

func (i *MoonIcon) Close() {
	i.Image.Dispose()
}
