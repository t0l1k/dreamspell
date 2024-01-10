package icons

import (
	"strconv"

	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/eui"
)

type DayIcon struct {
	eui.View
	day                                         string
	kinIcon                                     *KinIcon
	lblDt, lblMoonDay, lblWeekDay, lblYearDayNr *eui.Text
}

func NewDayIcon(day string) *DayIcon {
	i := &DayIcon{}
	i.SetupView()

	dt := i.calcDt(day)

	i.lblDt = eui.NewText(dt.Strings()[0])
	i.Add(i.lblDt)

	i.lblWeekDay = eui.NewText(dt.Strings()[1])
	i.Add(i.lblWeekDay)

	i.lblYearDayNr = eui.NewText(strconv.Itoa(dt.FindDayInYear()))
	i.Add(i.lblYearDayNr)

	i.lblMoonDay = eui.NewText(strconv.Itoa(dt.FindMoonDayNr()))
	i.Add(i.lblMoonDay)

	i.kinIcon = NewKinSealIcon(dt.FindKin())
	i.Add(i.kinIcon)

	i.Setup(day)
	return i
}

func (i *DayIcon) Setup(day string) {
	dt := i.calcDt(day)
	mweek := (dt.FindMoonDayNr()-1)/7 + 1
	bg, fg := lib.SealColor(mweek).Color()
	// fg2 := eui.White
	// if bg == fg2 || bg == eui.Yellow {
	// 	fg2 = eui.Black
	// }
	i.Bg(bg)

	i.lblDt.SetText(dt.Strings()[0])
	i.lblDt.Bg(bg)
	i.lblDt.Fg(fg)

	i.lblWeekDay.SetText(dt.Strings()[1])
	i.lblWeekDay.Bg(bg)
	i.lblWeekDay.Fg(fg)

	i.lblYearDayNr.SetText(strconv.Itoa(dt.FindDayInYear()))
	i.lblYearDayNr.Bg(bg)
	i.lblYearDayNr.Fg(fg)

	dNr := (dt.FindDayPlazma() + 2) % 4
	fg, _ = lib.SealColor(dNr + 1).Color()
	if bg == fg {
		fg = eui.Black
	}

	i.lblMoonDay.SetText(strconv.Itoa(dt.FindMoonDayNr()))
	i.lblMoonDay.Bg(bg)
	i.lblMoonDay.Fg(fg)

	i.kinIcon.Setup(dt.FindKin())
}

func (i *DayIcon) calcDt(day string) *lib.Convert {
	i.day = day
	dt := lib.NewConvert(i.day)
	return dt
}

func (i *DayIcon) Resize(rect []int) {
	i.View.Resize(rect)
	x, y := i.GetRect().X, i.GetRect().Y
	szWidth8 := i.GetRect().GetLowestSize() / 8
	i.lblMoonDay.Resize([]int{x, y, szWidth8 * 4, szWidth8 * 4})
	i.kinIcon.Resize([]int{x, y + szWidth8*4, szWidth8 * 4, szWidth8 * 4})
	i.lblDt.Resize([]int{x + szWidth8*4, y, szWidth8 * 4, szWidth8 * 2})
	i.lblWeekDay.Resize([]int{x + szWidth8*6, y + szWidth8*2, szWidth8 * 2, szWidth8})
	i.lblYearDayNr.Resize([]int{x + szWidth8*6, y + szWidth8*7, szWidth8 * 2, szWidth8})
}
