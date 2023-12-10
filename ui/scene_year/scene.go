package scene_dreamspell_year

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res"
	"github.com/t0l1k/dreamspell/ui/icons"
	"github.com/t0l1k/eui"
)

type SceneDreamSpellYear struct {
	eui.SceneBase
	layout *eui.GridLayoutDownRight
	dt     time.Time
}

func NewSceneDreamSpellYear() *SceneDreamSpellYear {
	sc := &SceneDreamSpellYear{}
	sc.layout = eui.NewGridLayoutDownRight(16, 30)
	sc.layout.SetCellMargin(1)
	sc.dt = time.Now()
	clrs := []color.Color{eui.Red, eui.White, eui.Blue, eui.Yellow}
	layout := "2006.01.02"
	tm0 := lib.NewConvert(sc.dt.Format(layout))
	dt0 := tm0.FindDreamspellYearBeginDate()
	weekDays := genWeekDays(dt0, sc.dt)
	yearKin := tm0.FindYearKin()
	emptyIcon := eui.NewIcon(res.GetSymbolAll()[3])

	layout2 := "02 Jan 2006"

	beg := tm0.FindDreamspellYearBeginDate()
	end := time.Date(beg.Year()+1, time.July, 25, 0, 0, 0, 0, time.Local)
	sBeg := beg.Format(layout2)
	sEnd := end.Format(layout2)

	fmt.Println(yearKin, sBeg, sEnd, weekDays)

	sc.layout.Add(icons.NewKinSealIcon(yearKin)) // кин года
	for _, clr := range clrs {                   // плазмы
		for _, plazma := range lib.GetPlazmas() {
			plazmaIcon := eui.NewIcon(res.GetPlazmaAll()[plazma-1])
			plazmaIcon.Bg(clr)
			sc.layout.Add(plazmaIcon)
		}
	}
	sc.layout.Add(emptyIcon)

	sc.layout.Add(eui.NewText(sBeg)) // дата начало года
	i := 1
	for _, clr := range clrs { // номер дня
		for j := 1; j <= 7; j++ {
			nrLbl := eui.NewText(strconv.Itoa(i))
			nrLbl.Bg(clr)
			nrLbl.Fg(eui.Black)
			sc.layout.Add(nrLbl)
			i++
		}
	}
	sc.layout.Add(emptyIcon)

	sc.layout.Add(eui.NewText(sEnd)) // дата конца года
	i = 1
	for _, clr := range clrs { // дни недели
		for _, v := range weekDays {
			nrLbl := eui.NewText(v)
			nrLbl.Bg(clr)
			nrLbl.Fg(eui.Black)
			sc.layout.Add(nrLbl)
		}
	}
	sc.layout.Add(emptyIcon)

	arrTons := res.GetTonAll()

	for i := 0; i < 13; i++ {
		ton := eui.NewIcon(arrTons[i])
		ton.Bg(eui.White)
		sc.layout.Add(ton)
		for i := 0; i < 28; i++ {
			if dt0.Month() == time.February && dt0.Day() == 29 {
				dt0 = dt0.Add(time.Duration(time.Hour * 24))
				i--
				continue
			}
			tm := lib.NewConvert(dt0.Format(layout))
			kin := icons.NewKinSealIcon(tm.FindKin())
			sc.layout.Add(kin)
			dt0 = dt0.Add(time.Duration(time.Hour * 24))
		}
		if i < 12 {
			sc.layout.Add(emptyIcon)
		} else {
			tm := lib.NewConvert(dt0.Format(layout))
			kin := icons.NewKinSealIcon(tm.FindKin())
			sc.layout.Add(kin)
		}
	}

	sc.Add(sc.layout)
	sc.Resize()
	return sc
}

func (sc *SceneDreamSpellYear) Update(dt int) {
	for _, v1 := range sc.layout.Container {
		v1.Update(dt)
	}
}

func (sc *SceneDreamSpellYear) Draw(surface *ebiten.Image) {
	for _, v1 := range sc.layout.Container {
		v1.Draw(surface)
	}
}

func (sc *SceneDreamSpellYear) Resize() {
	w0, h0 := eui.GetUi().Size()
	sc.layout.Resize([]int{0, 0, w0, h0})
}

func genWeekDays(dt0, dt time.Time) []string {
	dt1 := time.Date(dt.Year(), time.March, 7, 0, 0, 0, 0, time.Local)
	if dt1.Before(dt0) {
		y := dt.Year()
		y++
		dt1 = time.Date(y, time.March, 7, 0, 0, 0, 0, time.Local)
	}
	var weekDays []string
	if dt0.Weekday() == dt1.Weekday() {
		for i := 0; i < 7; i++ {
			w := dt0.Weekday().String()[:3]
			weekDays = append(weekDays, w)
			dt0 = dt0.Add(time.Duration(time.Hour * 24))
		}
	} else {
		for i := 0; i < 7; i++ {
			w := dt0.Weekday().String()[:2]
			w += "/"
			w += dt1.Weekday().String()[:2]
			weekDays = append(weekDays, w)
			dt0 = dt0.Add(time.Duration(time.Hour * 24))
			dt1 = dt1.Add(time.Duration(time.Hour * 24))
		}
	}
	return weekDays
}
