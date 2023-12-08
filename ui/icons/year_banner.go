package icons

import (
	"strconv"
	"time"

	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/eui"
)

type YearBanner struct {
	eui.View
	dt                                   time.Time
	lblYear, lblNs, lblYearKin           *eui.Text
	moonNrLbl, moonPeriodLbl             *eui.Text
	moonPeriodDaysLbl, moonPeriodKinsLbl *eui.Text
	yearKinIcon                          *KinIcon
}

func NewYearBanner(dt time.Time) *YearBanner {
	y := &YearBanner{}
	y.SetupView()
	y.dt = dt
	y.lblNs = eui.NewText("")
	y.Add(y.lblNs)
	y.lblYear = eui.NewText("")
	y.Add(y.lblYear)
	y.moonNrLbl = eui.NewText("")
	y.Add(y.moonNrLbl)
	y.moonPeriodLbl = eui.NewText("")
	y.Add(y.moonPeriodLbl)
	y.moonPeriodDaysLbl = eui.NewText("")
	y.Add(y.moonPeriodDaysLbl)
	y.moonPeriodKinsLbl = eui.NewText("")
	y.Add(y.moonPeriodKinsLbl)
	y.yearKinIcon = NewKinSealIcon(lib.GetTzolkin().GetKin(1))
	y.Add(y.yearKinIcon)
	y.lblYearKin = eui.NewText("Кин года")
	y.Add(y.lblYearKin)
	y.Setup(dt)
	return y
}

func (y *YearBanner) Setup(dt0 time.Time) {
	y.dt = dt0

	layout := "2006.01.02"
	dayNr := lib.NewConvert(y.dt.Format(layout)).FindMoonDayNr()
	dtMoonBegin := y.dt.Add(time.Duration(time.Hour * -24 * time.Duration(dayNr-1)))
	tm0 := lib.NewConvert(dtMoonBegin.Format(layout))
	yearKin := tm0.FindYearKin()
	moonNr := tm0.FindMoonNr()
	layout2 := "02 Jan 2006"

	beg := tm0.FindDreamspellYearBeginDate()
	end := time.Date(beg.Year()+1, time.July, 25, 0, 0, 0, 0, time.Local)
	sBeg := beg.Format(layout2)
	sEnd := end.Format(layout2)
	sYear := sBeg + " - " + sEnd

	y.yearKinIcon.Setup(yearKin)
	y.Bg(y.yearKinIcon.GetBg())
	y.lblYearKin.Bg(y.yearKinIcon.GetBg())
	y.lblYearKin.Fg(y.yearKinIcon.GetFg())

	y.lblYear.SetText(sYear)
	y.lblNs.SetText(tm0.GetNSShort())
	y.lblYear.Bg(y.yearKinIcon.GetBg())
	y.lblYear.Fg(y.yearKinIcon.GetFg())
	y.lblNs.Bg(y.yearKinIcon.GetBg())
	y.lblNs.Fg(y.yearKinIcon.GetFg())

	bg1 := eui.YellowGreen
	fg1 := eui.Black

	str := lib.Ton(moonNr).MoonNrRus() + " Луна"

	y.moonNrLbl.SetText(str)
	y.moonNrLbl.Bg(bg1)
	y.moonNrLbl.Fg(fg1)

	dt := y.dt.Add(time.Duration(time.Hour * 24 * 27))
	str = y.dt.Format(layout2) + " - " + dt.Format(layout2)
	y.moonPeriodLbl.SetText(str)
	y.moonPeriodLbl.Bg(bg1)
	y.moonPeriodLbl.Fg(fg1)

	tm1 := lib.NewConvert(dt.Format(layout))
	str = "Дни " + strconv.Itoa(tm0.FindDayInYear()) + " - " + strconv.Itoa(tm1.FindDayInYear())
	y.moonPeriodDaysLbl.SetText(str)
	y.moonPeriodDaysLbl.Bg(bg1)
	y.moonPeriodDaysLbl.Fg(fg1)

	str = "Кины " + strconv.Itoa(tm0.FindKin().GetNr()) + " - " + strconv.Itoa(tm1.FindKin().GetNr())
	y.moonPeriodKinsLbl.SetText(str)
	y.moonPeriodKinsLbl.Bg(bg1)
	y.moonPeriodKinsLbl.Fg(fg1)
}

func (yb *YearBanner) Resize(rect []int) {
	yb.View.Resize(rect)
	x0, y0 := yb.GetRect().Pos()
	w0, h0 := yb.GetRect().Size()
	w := w0 / 3
	h := h0 / 6
	x, y := x0, y0
	yb.lblNs.Resize([]int{x, y, w * 2, h})
	y += h
	yb.lblYear.Resize([]int{x, y, w * 2, h})
	y += h
	yb.moonNrLbl.Resize([]int{x, y, w * 2, h})
	y += h
	yb.moonPeriodLbl.Resize([]int{x, y, w * 2, h})
	y += h
	yb.moonPeriodDaysLbl.Resize([]int{x, y, w * 2, h})
	y += h
	yb.moonPeriodKinsLbl.Resize([]int{x, y, w * 2, h})
	sz8 := yb.GetRect().GetLowestSize() / 8
	x += w * 2
	y = y0
	yb.yearKinIcon.Resize([]int{x, y, w, sz8 * 6})
	y = y + sz8*6
	yb.lblYearKin.Resize([]int{x, y, w, sz8 * 2})
}
