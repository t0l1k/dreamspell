package scene_moon

import (
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res/img/moon"
	"github.com/t0l1k/eui"
)

type MoonBanner struct {
	Image          *ebiten.Image
	rect           *eui.Rect
	Dirty, Visible bool
	dt             time.Time
}

func NewMoonBanner(dt time.Time, rect []int) *MoonBanner {
	m := &MoonBanner{
		dt:      dt,
		rect:    eui.NewRect(rect),
		Dirty:   true,
		Visible: true,
	}
	m.setup(dt)
	return m
}

func (i *MoonBanner) setup(dt time.Time) {
	i.dt = dt
	i.Dirty = true
}

func (i *MoonBanner) Layout() {
	w, h := i.rect.Size()
	if i.Image == nil {
		i.Image = ebiten.NewImage(w, h)
	} else {
		i.Image.Clear()
	}
	bg := eui.White
	fg := eui.Black
	i.Image.Fill(bg)

	layout := "2006.01.02"
	tm0 := lib.NewConvert(i.dt.Format(layout))

	sz8 := i.rect.GetLowestSize() / 8
	sz6 := i.rect.GetLowestSize() / 6
	yearKin := tm0.FindYearKin()
	kinIcon := NewKinIcon(yearKin, []int{sz6 * 2 / 2, sz8, sz6 * 4, sz6 * 4}, bg, fg)
	kinIcon.Draw(i.Image)

	lblYearKin := eui.NewLabel("Кин года", []int{sz6 / 2, sz6 * 5, sz6 * 5, sz6}, bg, fg)
	lblYearKin.Draw(i.Image)

	bg2 := eui.YellowGreen
	fg2 := eui.Black
	lblNs := eui.NewLabel(tm0.GetNSShort(), []int{sz6 / 2, 0, sz6 * 5, sz8}, bg2, fg2)
	lblNs.Draw(i.Image)

	// eui.DrawRect(i.Image, eui.NewRect([]int{sz6 / 2, 0, sz6 * 5, sz6 * 5}), eui.Black)

	moonNr := tm0.FindMoonNr()
	moonTonImg := eui.NewIcon(
		ebiten.NewImageFromImage(
			moon.GetMoonPngs().Get(lib.Ton(moonNr))), []int{h + sz6/2, 0, sz6 * 5, sz6 * 5})
	moonTonImg.Draw(i.Image)

	s := lib.Ton(moonNr).TotemRus()
	lblTotem := eui.NewLabel("Тотем "+s, []int{h + sz6/2, sz6 * 5, sz6 * 5, sz6}, bg, fg)
	lblTotem.Draw(i.Image)

	// eui.DrawRect(i.Image, eui.NewRect([]int{h, 0, h, h}), eui.Black)

	str := lib.Ton(moonNr).StringRus() + " Луна"
	moonLbl := eui.NewLabel(str, []int{h * 2, 0, w / 2, sz8 * 4}, bg, fg)
	moonLbl.Draw(i.Image)

	str = lib.Ton(moonNr).QuestionRus()
	moonQuestionLbl := eui.NewLabel(str, []int{h * 2, sz8 * 4, w / 2, sz8 * 2}, bg, fg)
	moonQuestionLbl.Draw(i.Image)

	str = lib.Ton(moonNr).FuncRus()
	moonFuncLbl := eui.NewLabel(str, []int{h * 2, sz8 * 6, w / 2, sz8 * 2}, bg, fg)
	moonFuncLbl.Draw(i.Image)

	str = lib.Ton(moonNr).MoonNrRus() + " Луна"
	moonNrLbl := eui.NewLabel(str, []int{w - w/2 + h*2, 0, w/2 - h*2, sz8}, bg2, fg2)
	moonNrLbl.Draw(i.Image)

	layout2 := "02 Jan 2006"

	dt := i.dt.Add(time.Duration(time.Hour * 24 * 27))
	str = i.dt.Format(layout2) + " - " + dt.Format(layout2)
	moonPeriodLbl := eui.NewLabel(str, []int{w - w/2 + h*2, sz8, w/2 - h*2, sz8}, bg2, fg2)
	moonPeriodLbl.Draw(i.Image)

	tm1 := lib.NewConvert(dt.Format(layout))
	str = "Дни " + strconv.Itoa(tm0.FindDayInYear()) + " - " + strconv.Itoa(tm1.FindDayInYear())
	moonPeriodDaysLbl := eui.NewLabel(str, []int{w - w/2 + h*2, sz8 * 2, w/2 - h*2, sz8}, bg2, fg2)
	moonPeriodDaysLbl.Draw(i.Image)

	str = "Кины " + strconv.Itoa(tm0.FindKin().GetNr()) + " - " + strconv.Itoa(tm1.FindKin().GetNr())
	moonPeriodKinsLbl := eui.NewLabel(str, []int{w - w/2 + h*2, sz8 * 3, w/2 - h*2, sz8}, bg2, fg2)
	moonPeriodKinsLbl.Draw(i.Image)

	str3 := lib.Ton(moonNr).MoonFunc3Rus()

	sz8a := int(float64(sz8) * 1.5)
	for dx, s := range str3 {
		lbl := eui.NewLabel(s, []int{w - w/2 + h*2, sz8*3 + dx*sz8a, w/2 - h*2, sz8a}, bg, fg)
		lbl.Draw(i.Image)
	}

	eui.DrawRect(i.Image, eui.NewRect([]int{0, 0, w, h}), eui.Black)
}

func (i *MoonBanner) Size() (int, int) {
	return i.rect.Size()
}

func (i *MoonBanner) Update(dt int) {}

func (i *MoonBanner) Draw(surface *ebiten.Image) {
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

func (i *MoonBanner) Resize(rect []int) {
	i.rect = eui.NewRect(rect)
	i.Dirty = true
	i.Image = nil
}

func (i *MoonBanner) Close() {
	i.Image.Dispose()
}
