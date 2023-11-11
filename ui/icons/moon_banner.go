package icons

import (
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res/img/moon"
	"github.com/t0l1k/eui"
)

type MoonBanner struct {
	eui.View
	dt                                                  time.Time
	yearKinIcon                                         *KinIcon
	lblYearKin, lblTotem, moonLbl                       *eui.Text
	moonQuestionLbl, moonFuncLbl, moonNrLbl             *eui.Text
	moonPeriodLbl, moonPeriodDaysLbl, moonPeriodKinsLbl *eui.Text
	moonFunc0, moonFunc1, moonFunc2                     *eui.Text
	moonTonImg                                          *eui.Icon
	str3                                                []string
}

func NewMoonBanner(dt time.Time) *MoonBanner {
	i := &MoonBanner{}
	i.SetupView()
	i.dt = dt

	_, yearKin, moonNr := i.calcTm()

	i.yearKinIcon = NewKinSealIcon(yearKin)
	i.Add(i.yearKinIcon)

	i.moonTonImg = eui.NewIcon(ebiten.NewImageFromImage(moon.GetMoonPngs().Get(lib.Ton(moonNr))))
	i.Add(i.moonTonImg)

	i.lblYearKin = eui.NewText("Кин года")
	i.Add(i.lblYearKin)

	i.lblTotem = eui.NewText("")
	i.Add(i.lblTotem)
	i.moonLbl = eui.NewText("")
	i.Add(i.moonLbl)
	i.moonQuestionLbl = eui.NewText("")
	i.Add(i.moonQuestionLbl)
	i.moonFuncLbl = eui.NewText("")
	i.Add(i.moonFuncLbl)
	i.moonNrLbl = eui.NewText("")
	i.Add(i.moonNrLbl)
	i.moonPeriodLbl = eui.NewText("")
	i.Add(i.moonPeriodLbl)
	i.moonPeriodDaysLbl = eui.NewText("")
	i.Add(i.moonPeriodDaysLbl)
	i.moonPeriodKinsLbl = eui.NewText("")
	i.Add(i.moonPeriodKinsLbl)
	i.moonFunc0 = eui.NewText("")
	i.Add(i.moonFunc0)
	i.moonFunc1 = eui.NewText("")
	i.Add(i.moonFunc1)
	i.moonFunc2 = eui.NewText("")
	i.Add(i.moonFunc2)

	i.Setup(dt)
	return i
}

func (i *MoonBanner) Setup(dt0 time.Time) {
	i.dt = dt0

	layout := "2006.01.02"
	layout2 := "02 Jan 2006"

	tm0, yearKin, moonNr := i.calcTm()

	bg0 := eui.White
	fg0 := eui.Black
	i.Bg(bg0)

	i.yearKinIcon.Setup(yearKin)

	i.lblYearKin.Bg(bg0)
	i.lblYearKin.Fg(fg0)

	i.moonTonImg.SetIcon(ebiten.NewImageFromImage(moon.GetMoonPngs().Get(lib.Ton(moonNr))))

	s := lib.Ton(moonNr).TotemRus()
	i.lblTotem.SetText("Тотем " + s)
	i.lblTotem.Bg(bg0)
	i.lblTotem.Fg(fg0)

	str := lib.Ton(moonNr).StringRus() + " Луна"
	i.moonLbl.SetText(str)
	i.moonLbl.Bg(bg0)
	i.moonLbl.Fg(fg0)

	str = lib.Ton(moonNr).QuestionRus()
	i.moonQuestionLbl.SetText(str)
	i.moonQuestionLbl.Bg(bg0)
	i.moonQuestionLbl.Fg(fg0)

	str = lib.Ton(moonNr).FuncRus()
	i.moonFuncLbl.SetText(str)
	i.moonFuncLbl.Bg(bg0)
	i.moonFuncLbl.Fg(fg0)

	bg1 := eui.GreenYellow
	fg1 := eui.Black

	str = lib.Ton(moonNr).MoonNrRus() + " Луна"

	i.moonNrLbl.SetText(str)
	i.moonNrLbl.Bg(bg1)
	i.moonNrLbl.Fg(fg1)

	dt := i.dt.Add(time.Duration(time.Hour * 24 * 27))
	str = i.dt.Format(layout2) + " - " + dt.Format(layout2)
	i.moonPeriodLbl.SetText(str)
	i.moonPeriodLbl.Bg(bg1)
	i.moonPeriodLbl.Fg(fg1)

	tm1 := lib.NewConvert(dt.Format(layout))
	str = "Дни " + strconv.Itoa(tm0.FindDayInYear()) + " - " + strconv.Itoa(tm1.FindDayInYear())
	i.moonPeriodDaysLbl.SetText(str)
	i.moonPeriodDaysLbl.Bg(bg1)
	i.moonPeriodDaysLbl.Fg(fg1)

	str = "Кины " + strconv.Itoa(tm0.FindKin().GetNr()) + " - " + strconv.Itoa(tm1.FindKin().GetNr())
	i.moonPeriodKinsLbl.SetText(str)
	i.moonPeriodKinsLbl.Bg(bg1)
	i.moonPeriodKinsLbl.Fg(fg1)

	i.str3 = lib.Ton(moonNr).MoonFunc3Rus()
	i.moonFunc0.SetText(i.str3[0])
	i.moonFunc0.Bg(bg0)
	i.moonFunc0.Fg(fg0)
	i.moonFunc1.SetText(i.str3[1])
	i.moonFunc1.Bg(bg0)
	i.moonFunc1.Fg(fg0)
	i.moonFunc2.SetText(i.str3[2])
	i.moonFunc2.Bg(bg0)
	i.moonFunc2.Fg(fg0)
}

func (i *MoonBanner) calcTm() (*lib.Convert, *lib.Kin, int) {
	layout := "2006.01.02"
	tm0 := lib.NewConvert(i.dt.Format(layout))
	yearKin := tm0.FindYearKin()
	moonNr := tm0.FindMoonNr()
	return tm0, yearKin, moonNr
}

func (i *MoonBanner) Resize(rect []int) {
	i.View.Resize(rect)
	x, y, w, h := i.GetRect().X, i.GetRect().Y, i.GetRect().W, i.GetRect().H
	sz8 := i.GetRect().GetLowestSize() / 8
	i.yearKinIcon.Resize([]int{x, y, sz8 * 8, sz8 * 6})
	i.lblYearKin.Resize([]int{x + sz8, y + sz8*6, sz8 * 6, sz8 * 2})
	i.moonTonImg.Resize([]int{x + sz8*8, y, sz8 * 8, sz8 * 6})
	i.lblTotem.Resize([]int{x + sz8*9, y + sz8*6, sz8 * 6, sz8 * 2})
	i.moonLbl.Resize([]int{x + h*2, y, w / 2, sz8 * 4})
	i.moonQuestionLbl.Resize([]int{x + h*2, y + sz8*4, w / 2, sz8 * 2})
	i.moonFuncLbl.Resize([]int{x + h*2, y + sz8*6, w / 2, sz8 * 2})
	i.moonNrLbl.Resize([]int{x + w - w/2 + h*2, y, w/2 - h*2, sz8})
	i.moonPeriodLbl.Resize([]int{x + (w - w/2 + h*2), y + sz8, w/2 - h*2, sz8})
	i.moonPeriodDaysLbl.Resize([]int{x + (w - w/2 + h*2), y + sz8*2, w/2 - h*2, sz8})
	i.moonPeriodKinsLbl.Resize([]int{x + (w - w/2 + h*2), y + sz8*3, w/2 - h*2, sz8})

	sz8a := int(float64(sz8) * 1.2)

	i.moonFunc0.Resize([]int{x + w - w/2 + h*2, y + sz8*4, w/2 - h*2, sz8a})
	i.moonFunc1.Resize([]int{x + w - w/2 + h*2, y + sz8*4 + sz8a, w/2 - h*2, sz8a})
	i.moonFunc2.Resize([]int{x + w - w/2 + h*2, y + sz8*4 + sz8a*2, w/2 - h*2, sz8a})
}
