package scene_moon

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/eui"
)

type DayIcon struct {
	Image          *ebiten.Image
	rect           *eui.Rect
	Dirty, Visible bool
	day            string
}

func NewDayIcon(day string, rect []int) *DayIcon {
	return &DayIcon{
		day:     day,
		rect:    eui.NewRect(rect),
		Dirty:   true,
		Visible: true,
	}
}

func (i *DayIcon) setup(day string) {
	i.day = day
	i.Dirty = true
}

func (i *DayIcon) Layout() {
	w, h := i.rect.Size()
	if i.Image == nil {
		i.Image = ebiten.NewImage(w, h)
	} else {
		i.Image.Clear()
	}
	dt := lib.NewConvert(i.day)

	clrs := []color.Color{eui.Red, eui.White, eui.Blue, eui.Yellow}
	mweek := (dt.FindMoonDayNr()-1)/7 + 1
	fg := clrs[int(lib.SealColor(mweek)-1)]
	r, g, b, _ := fg.RGBA()
	bg := color.RGBA{uint8(r), uint8(g), uint8(b), 255}
	i.Image.Fill(bg)

	szWidth2 := i.rect.GetLowestSize() / 2
	szWidth4 := i.rect.GetLowestSize() / 4
	szHeight := i.rect.GetLowestSize() / 2
	lblDt := eui.NewLabel(dt.Strings()[0], []int{szWidth2, 0, szWidth2, szHeight / 2}, bg, eui.Black)
	lblDt.Draw(i.Image)

	lblWeekDay := eui.NewLabel(dt.Strings()[1], []int{w - szWidth4, szHeight / 2, szWidth4, szHeight / 3}, bg, eui.Black)
	lblWeekDay.Draw(i.Image)

	lblYearDayNr := eui.NewLabel(strconv.Itoa(dt.FindDayInYear()), []int{szWidth2 + szWidth4, szHeight + szHeight/3*2, szWidth4, szHeight / 3}, bg, eui.Black)
	lblYearDayNr.Draw(i.Image)

	// moonAge := metonic.GetMoonAge(metonic.GetDate(i.day))

	// moonAgeImg := eui.NewIcon(
	// 	ebiten.NewImageFromImage(
	// 		moon2.GetMoonPngs()[moonAge-1]),
	// 	[]int{szWidth2, szHeight, szWidth2, szHeight})
	// moonAgeImg.Draw(i.Image)

	// lblMoonAge := eui.NewLabel(strconv.Itoa(moonAge), []int{szWidth2, szHeight, szWidth4, szHeight / 3}, eui.Black, eui.Yellow)
	// lblMoonAge.Draw(i.Image)

	dNr := (dt.FindDayPlazma() + 2) % 4
	fg = clrs[int(lib.SealColor(dNr))]
	if bg == fg {
		fg = eui.Black
	}
	lblMoonDay := eui.NewLabel(strconv.Itoa(dt.FindMoonDayNr()), []int{0, 0, szWidth2, szHeight}, bg, fg)
	lblMoonDay.Draw(i.Image)

	kinIcon := NewKinIcon(dt.FindKin(), []int{0, szHeight, szWidth2, szWidth2}, bg, fg)
	kinIcon.Draw(i.Image)

	eui.DrawRect(i.Image, eui.NewRect([]int{0, 0, w, h}), eui.Black)

	i.Dirty = false
}

func (i *DayIcon) Size() (int, int) {
	return i.rect.Size()
}

func (i *DayIcon) Update(dt int) {}

func (i *DayIcon) Draw(surface *ebiten.Image) {
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

func (i *DayIcon) Resize(rect []int) {
	i.rect = eui.NewRect(rect)
	i.Dirty = true
	i.Image = nil
}

func (i *DayIcon) Close() {
	i.Image.Dispose()
}
