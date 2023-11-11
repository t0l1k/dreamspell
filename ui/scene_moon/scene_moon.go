package scene_moon

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/ui/icons"
	"github.com/t0l1k/eui"
)

type SceneMoon struct {
	eui.SceneBase
	dt             time.Time
	moonBanner     *icons.MoonBanner
	moonIcon       *icons.MoonIcon
	plazmas        []*icons.PlazmaIcon
	weeks          []*icons.MoonWeekIcon
	days           []*icons.DayIcon
	dirty          bool
	lblYear, lblNs *eui.Text
}

func NewSceneMoon() *SceneMoon {
	return &SceneMoon{
		dt:    time.Now(),
		dirty: true,
	}
}

func (sc *SceneMoon) setup() {
	if sc.plazmas == nil {
		for _, plazma := range lib.GetPlazmas() {
			p := icons.NewPlazmaIcon(plazma)
			sc.plazmas = append(sc.plazmas, p)
			sc.Add(p)
		}
	}

	layout := "2006.01.02"
	dayNr := lib.NewConvert(sc.dt.Format(layout)).FindMoonDayNr()
	dtMoonBegin := sc.dt.Add(time.Duration(time.Hour * -24 * time.Duration(dayNr-1)))
	tm0 := lib.NewConvert(dtMoonBegin.Format(layout))
	day := tm0.FindMoonDayNr()

	if day > 1 {
		dayNr := lib.NewConvert(sc.dt.Format(layout)).FindMoonDayNr()
		dtMoonBegin = sc.dt.Add(time.Duration(time.Hour * -24 * time.Duration(dayNr)))
		tm0 = lib.NewConvert(dtMoonBegin.Format(layout))
	}

	layout2 := "02 Jan 2006"

	beg := tm0.FindDreamspellYearBeginDate()
	end := time.Date(beg.Year()+1, time.July, 25, 0, 0, 0, 0, time.Local)
	sBeg := beg.Format(layout2)
	sEnd := end.Format(layout2)
	sYear := sBeg + " - " + sEnd
	if sc.lblYear == nil {
		bg := eui.GreenYellow
		fg := eui.Black
		sc.lblYear = eui.NewText(sYear)
		sc.lblYear.Bg(bg)
		sc.lblYear.Fg(fg)
		sc.Add(sc.lblYear)
		sc.lblNs = eui.NewText(tm0.GetNSShort())
		sc.lblNs.Bg(bg)
		sc.lblNs.Fg(fg)
		sc.Add(sc.lblNs)

	} else {
		sc.lblYear.SetText(sYear)
		sc.lblNs.SetText(tm0.GetNSShort())
	}

	if sc.moonBanner == nil {
		sc.moonBanner = icons.NewMoonBanner(dtMoonBegin)
		sc.Add(sc.moonBanner)
	} else {
		sc.moonBanner.Setup(dtMoonBegin)
	}

	moonKin := tm0.FindMoonKin()
	if sc.moonIcon == nil {
		sc.moonIcon = icons.NewMoonIcon(moonKin)
		sc.Add(sc.moonIcon)
	} else {
		sc.moonIcon.Setup(moonKin)
	}

	mweek := tm0.FindMoonNr()*4 - 3
	if sc.weeks == nil {
		for week := 0; week < 4; week++ {
			w := icons.NewMoonWeekIcon(week + mweek)
			sc.weeks = append(sc.weeks, w)
			sc.Add(w)
		}
	} else {
		for week := 0; week < 4; week++ {
			sc.weeks[week].Setup(week + mweek)
		}
	}

	if sc.days == nil {
		for i := 0; i < 28; i++ {
			day := dtMoonBegin.Format(layout)
			d := icons.NewDayIcon(day)
			sc.days = append(sc.days, d)
			sc.Add(d)
			dtMoonBegin = sc.nextDay(dtMoonBegin)
		}
	} else {
		for i := 0; i < 28; i++ {
			day := dtMoonBegin.Format(layout)
			sc.days[i].Setup(day)
			dtMoonBegin = sc.nextDay(dtMoonBegin)
		}
	}
	sc.dirty = false
}

func (*SceneMoon) nextDay(dtMoonBegin time.Time) time.Time {
	dtMoonBegin = dtMoonBegin.Add(time.Duration(time.Hour * 24))
	_, m, d := dtMoonBegin.Date()
	if m == time.February && d == 29 {
		dtMoonBegin = dtMoonBegin.Add(time.Duration(time.Hour * 24))
	}
	return dtMoonBegin
}

func (sc *SceneMoon) Entered() {
	sc.setup()
	sc.Resize()
}

func (sc *SceneMoon) Update(dt int) {
	if sc.dirty {
		sc.setup()
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyF1) {
		sc.dt = sc.dt.Add(time.Duration(time.Hour * 24 * 28))
		sc.dirty = true
		log.Println("F1", sc.dt)
	} else if inpututil.IsKeyJustReleased(ebiten.KeyF2) {
		sc.dt = sc.dt.Add(time.Duration(time.Hour * -24 * 28))
		sc.dirty = true
		log.Println("F2", sc.dt)
	}
	for _, v := range sc.Container {
		v.Update(dt)
	}
}

func (sc *SceneMoon) Resize() {
	w, h := eui.GetUi().Size()
	r, c := 8, 6
	sz := r
	size := int(float64(w) * 0.95)
	if w > h {
		size = int(float64(h) * 0.95)
		sz = c
	}
	cellSize := int(size / sz)
	x0 := w/2 - cellSize*r/2
	y0 := h/2 - cellSize*c/2
	h2 := int(float64(h) * 0.05)
	sc.lblYear.Resize([]int{x0 + cellSize*(r-2), y0 - h2/2, cellSize * 2, h2 / 2})
	sc.lblNs.Resize([]int{x0, y0 - h2/2, cellSize, h2 / 2})
	sc.moonBanner.Resize([]int{x0, y0, cellSize * r, cellSize})
	y0 += cellSize
	sc.moonIcon.Resize([]int{x0, y0, cellSize, cellSize})

	x, y := x0+cellSize, y0
	for _, icon := range sc.plazmas {
		icon.Resize([]int{x, y, cellSize, cellSize})
		x += cellSize
	}
	x, y = x0, y0+cellSize
	for _, week := range sc.weeks {
		week.Resize([]int{x, y, cellSize, cellSize})
		y += cellSize
	}

	x, y = x0+cellSize, y0+cellSize
	for i, day := range sc.days {
		i++
		day.Resize([]int{x, y, cellSize, cellSize})
		if i > 0 && i%7 == 0 {
			y += cellSize
			x = x0
		}
		x += cellSize
	}
}
