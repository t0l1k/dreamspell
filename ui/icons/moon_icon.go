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
	i.Bg(i.icon.GetBg())
	i.Fg(i.icon.GetFg())
	i.lbl.Bg(i.icon.GetBg())
	i.lbl.Fg(i.icon.GetFg())
}

func (i *MoonIcon) Resize(rect []int) {
	i.View.Resize(rect)
	x, y := i.GetRect().X, i.GetRect().Y
	sz := i.GetRect().GetLowestSize() / 8
	i.icon.Resize([]int{x, y, sz * 8, sz * 6})
	i.lbl.Resize([]int{x + sz, y + sz*6, sz * 6, sz * 2})
}
