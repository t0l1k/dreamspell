package scene_moon

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res/img/tons"
	"github.com/t0l1k/eui"
)

type WeekIcon struct {
	Image              *ebiten.Image
	rect               *eui.Rect
	Dirty, Visible     bool
	week, mweek, qweek int
}

func NewWeekIcon(week int, rect []int) *WeekIcon {
	w := &WeekIcon{
		week:    week,
		rect:    eui.NewRect(rect),
		Dirty:   true,
		Visible: true,
	}
	w.setup(week)
	return w
}

func (i *WeekIcon) setup(week int) {
	i.week = week
	i.mweek = i.week % 4
	if i.mweek == 0 {
		i.mweek = 4
	}
	i.qweek = i.week % 13
	if i.qweek == 0 {
		i.qweek = 13
	}
	i.Dirty = true
}

func (i *WeekIcon) Layout() {
	w, h := i.rect.Size()
	if i.Image == nil {
		i.Image = ebiten.NewImage(w, h)
	} else {
		i.Image.Clear()
	}

	clrs := []color.Color{eui.Red, eui.White, eui.Blue, eui.Yellow}
	bg := clrs[int(lib.SealColor(i.mweek)-1)]
	i.Image.Fill(bg)

	szWidth := i.rect.GetLowestSize() / 2
	szHeight := i.rect.GetLowestSize() / 3
	mWeekTonImg := eui.NewIcon(
		ebiten.NewImageFromImage(
			tons.GetTonPngs().Get(lib.Ton(i.mweek))),
		[]int{szWidth / 2, 0, szWidth, szHeight})
	mWeekTonImg.Draw(i.Image)
	qWeekTonImg := eui.NewIcon(
		ebiten.NewImageFromImage(
			tons.GetTonPngs().Get(lib.Ton(i.qweek))),
		[]int{szWidth / 2, szHeight, szWidth, szHeight})
	qWeekTonImg.Draw(i.Image)
	lblYearWeek := eui.NewLabel(strconv.Itoa(i.week), []int{szWidth / 2, szHeight * 2, szWidth, szHeight}, bg, eui.Black)
	lblYearWeek.Draw(i.Image)

	eui.DrawRect(i.Image, eui.NewRect([]int{0, 0, w, h}), eui.Black)

}

func (i *WeekIcon) Size() (int, int) {
	return i.rect.Size()
}

func (i *WeekIcon) Update(dt int) {}

func (i *WeekIcon) Draw(surface *ebiten.Image) {
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

func (i *WeekIcon) Resize(rect []int) {
	i.rect = eui.NewRect(rect)
	i.Dirty = true
	i.Image = nil
}

func (i *WeekIcon) Close() {
	i.Image.Dispose()
}
