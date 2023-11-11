package icons

import (
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/eui"
)

type MoonIcon struct {
	eui.View
	icon *KinIcon
	lbl  *eui.Text
}

func NewMoonIcon(moonKin *lib.Kin) *MoonIcon {
	i := &MoonIcon{}
	i.SetupView()
	i.icon = NewKinSealIcon(moonKin)
	i.Add(i.icon)
	i.lbl = eui.NewText("Кин луны")
	i.Add(i.lbl)
	i.Setup(moonKin)
	return i
}

func (i *MoonIcon) Setup(kin *lib.Kin) {
	i.icon.Setup(kin)
	bg0 := eui.White
	fg0 := eui.Black
	i.Bg(bg0)
	i.lbl.Bg(bg0)
	i.lbl.Fg(fg0)
}

func (i *MoonIcon) Resize(rect []int) {
	i.View.Resize(rect)
	x, y := i.GetRect().X, i.GetRect().Y
	sz := i.GetRect().GetLowestSize() / 8
	i.icon.Resize([]int{x, y, sz * 8, sz * 6})
	i.lbl.Resize([]int{x + sz, y + sz*6, sz * 6, sz * 2})
}
