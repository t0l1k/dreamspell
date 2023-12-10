package icons

import (
	"time"

	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res"
	"github.com/t0l1k/eui"
)

type MoonBanner struct {
	eui.View
	dt                              time.Time
	lblTotem, moonLbl               *eui.Text
	moonQuestionLbl, moonFuncLbl    *eui.Text
	moonFunc0, moonFunc1, moonFunc2 *eui.Text
	moonTonImg                      *eui.Icon
	str3                            []string
	YearBanner                      *YearBanner
}

func NewMoonBanner(dt time.Time) *MoonBanner {
	i := &MoonBanner{}
	i.SetupView()
	i.dt = dt

	_, moonNr := i.calcTm()

	i.moonTonImg = eui.NewIcon(res.GetMoonTonAll()[lib.Ton(moonNr-1)])
	i.Add(i.moonTonImg)

	i.lblTotem = eui.NewText("")
	i.Add(i.lblTotem)
	i.moonLbl = eui.NewText("")
	i.Add(i.moonLbl)
	i.moonQuestionLbl = eui.NewText("")
	i.Add(i.moonQuestionLbl)
	i.moonFuncLbl = eui.NewText("")
	i.Add(i.moonFuncLbl)
	i.moonFunc0 = eui.NewText("")
	i.Add(i.moonFunc0)
	i.moonFunc1 = eui.NewText("")
	i.Add(i.moonFunc1)
	i.moonFunc2 = eui.NewText("")
	i.Add(i.moonFunc2)

	i.YearBanner = NewYearBanner(dt)
	i.Add(i.YearBanner)

	i.Setup(dt)
	return i
}

func (i *MoonBanner) Setup(dt0 time.Time) {
	i.dt = dt0

	_, moonNr := i.calcTm()

	bg0 := eui.White
	fg0 := eui.Black
	i.Bg(bg0)

	i.moonTonImg.SetIcon(res.GetMoonTonAll()[lib.Ton(moonNr-1)])

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

	i.YearBanner.Setup(dt0)
}

func (i *MoonBanner) calcTm() (*lib.Convert, int) {
	layout := "2006.01.02"
	tm0 := lib.NewConvert(i.dt.Format(layout))
	moonNr := tm0.FindMoonNr()
	return tm0, moonNr
}

func (i *MoonBanner) Resize(rect []int) {
	i.View.Resize(rect)
	x, y, w, _ := i.GetRect().X, i.GetRect().Y, i.GetRect().W, i.GetRect().H
	sz8 := i.GetRect().GetLowestSize() / 8
	cellSize := w / 8
	i.moonTonImg.Resize([]int{x, y, cellSize, sz8 * 6})
	i.lblTotem.Resize([]int{x, y + sz8*6, cellSize, sz8 * 2})
	i.moonLbl.Resize([]int{x + cellSize, y, cellSize * 4, cellSize / 2})
	i.moonQuestionLbl.Resize([]int{x + cellSize, y + cellSize/2, cellSize * 3, sz8 * 2})
	i.moonFuncLbl.Resize([]int{x + cellSize, y + sz8*6, cellSize * 3, sz8 * 2})

	sz8a := int(float64(sz8) * 1.2)
	i.moonFunc0.Resize([]int{x + cellSize*4, y + sz8*4, cellSize, sz8a})
	i.moonFunc1.Resize([]int{x + cellSize*4, y + sz8*4 + sz8a, cellSize, sz8a})
	i.moonFunc2.Resize([]int{x + cellSize*4, y + sz8*4 + sz8a*2, cellSize, sz8a})

	i.YearBanner.Resize([]int{x + cellSize*5, y, cellSize * 3, cellSize})
}
