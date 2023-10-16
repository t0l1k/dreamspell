package scene_moon

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	plazmas "github.com/t0l1k/dreamspell/res/img/plasmas"
	"github.com/t0l1k/eui"
)

type PlazmaIcon struct {
	Image          *ebiten.Image
	rect           *eui.Rect
	Dirty, Visible bool
	plazma         lib.Plazma
}

func NewPlazmaIcon(plazma lib.Plazma, rect []int) *PlazmaIcon {
	return &PlazmaIcon{
		plazma:  plazma,
		rect:    eui.NewRect(rect),
		Dirty:   true,
		Visible: true,
	}
}

func (i *PlazmaIcon) Layout() {
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

	plazmaImg := eui.NewIcon(
		ebiten.NewImageFromImage(plazmas.GetPlazmaPngs().Get(i.plazma)),
		[]int{sz, 0, sz * 4, sz * 4})
	plazmaImg.Draw(i.Image)
	lblPlazma := eui.NewLabel(i.plazma.String(), []int{sz / 2, sz * 4, sz * 5, sz}, bg, fg)
	lblPlazma.Draw(i.Image)
	lblMotto := eui.NewLabel(i.plazma.MottoRus(), []int{sz / 2, sz * 5, sz * 5, sz}, bg, fg)
	lblMotto.Draw(i.Image)

	eui.DrawRect(i.Image, eui.NewRect([]int{0, 0, w, h}), fg)
	i.Dirty = false
}

func (i *PlazmaIcon) Size() (int, int) {
	return i.rect.Size()
}

func (i *PlazmaIcon) Update(dt int) {}

func (i *PlazmaIcon) Draw(surface *ebiten.Image) {
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

func (i *PlazmaIcon) Resize(rect []int) {
	i.rect = eui.NewRect(rect)
	i.Dirty = true
	i.Image = nil
}

func (i *PlazmaIcon) Close() {
	i.Image.Dispose()
}
