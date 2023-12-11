package icons

import (
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res"
	"github.com/t0l1k/eui"
)

type MoonIcon struct {
	eui.View
	icon *eui.Icon
	lbl  *eui.Text
}

func NewMoonIcon(moonNr int) *MoonIcon {
	i := &MoonIcon{}
	i.SetupView()

	i.icon = eui.NewIcon(res.GetMoonTonAll()[lib.Ton(moonNr-1)])
	i.Add(i.icon)
	i.lbl = eui.NewText("Тотем луны")
	i.Add(i.lbl)
	i.Setup(moonNr)
	return i
}

func (i *MoonIcon) Setup(moonNr int) {
	i.icon.SetIcon(res.GetMoonTonAll()[lib.Ton(moonNr-1)])
	str := lib.Ton(moonNr).TotemRus()
	i.lbl.SetText("Тотем " + str)
	bg := eui.White
	fg := eui.Black
	i.Bg(bg)
	i.Fg(fg)
	i.lbl.Bg(bg)
	i.lbl.Fg(fg)
}

func (i *MoonIcon) Resize(rect []int) {
	i.View.Resize(rect)
	x, y := i.GetRect().X, i.GetRect().Y
	sz := i.GetRect().GetLowestSize() / 8
	i.icon.Resize([]int{x, y, sz * 8, sz * 6})
	i.lbl.Resize([]int{x, y + sz*6, sz * 8, sz * 2})
}
