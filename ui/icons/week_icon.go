package icons

import (
	"image/color"
	"strconv"

	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res"
	"github.com/t0l1k/eui"
)

type MoonWeekIcon struct {
	eui.View
	week, mweek, qweek int
	moonIcon, quarIcon *eui.Icon
	yearLbl            *eui.Text
}

func NewMoonWeekIcon(week int) *MoonWeekIcon {
	i := &MoonWeekIcon{}
	i.SetupView()
	i.SetVertical()
	i.moonIcon = eui.NewIcon(nil)
	i.Add(i.moonIcon)
	i.quarIcon = eui.NewIcon(nil)
	i.Add(i.quarIcon)
	i.yearLbl = eui.NewText("")
	i.Add(i.yearLbl)
	i.Setup(week)
	return i
}

func (i *MoonWeekIcon) Setup(week int) {
	i.calcWeekNrs(week)
	i.moonIcon.SetIcon(res.GetTonAll()[lib.Ton(i.mweek-1)])
	i.quarIcon.SetIcon(res.GetTonAll()[lib.Ton(i.qweek-1)])
	i.yearLbl.SetText(strconv.Itoa(i.week))
	clrs := []color.Color{eui.Red, eui.White, eui.Blue, eui.Yellow}
	bg := clrs[int(lib.SealColor(i.mweek)-1)]
	i.Bg(bg)
	i.moonIcon.Bg(bg)
	i.quarIcon.Bg(bg)
	i.yearLbl.Bg(bg)
	i.yearLbl.Fg(eui.Black)
}

func (i *MoonWeekIcon) calcWeekNrs(week int) {
	i.week = week
	i.mweek = i.week % 4
	if i.mweek == 0 {
		i.mweek = 4
	}
	i.qweek = i.week % 13
	if i.qweek == 0 {
		i.qweek = 13
	}
}
