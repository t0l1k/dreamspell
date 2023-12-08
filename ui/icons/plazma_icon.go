package icons

import (
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res/img/plazmas"
	"github.com/t0l1k/eui"
)

type PlazmaIcon struct {
	eui.View
	plazmaIcon          *eui.Icon
	lblPlazma, lblMotto *eui.Text
	plazma              lib.Plazma
}

func NewPlazmaIcon(plazma lib.Plazma) *PlazmaIcon {
	i := &PlazmaIcon{}
	i.SetupView()
	i.plazma = plazma

	i.plazmaIcon = eui.NewIcon(plazmas.GetPlazmaPngs().Get(i.plazma))
	i.Add(i.plazmaIcon)

	i.lblPlazma = eui.NewText(i.plazma.String())
	i.Add(i.lblPlazma)

	i.lblMotto = eui.NewText(i.plazma.MottoRus())
	i.Add(i.lblMotto)

	i.Setup(plazma)
	return i
}

func (i *PlazmaIcon) Setup(p lib.Plazma) {
	i.plazma = p
	bg0 := eui.White
	fg0 := eui.Black
	i.Bg(bg0)

	i.plazmaIcon.SetIcon(plazmas.GetPlazmaPngs().Get(i.plazma))
	i.plazmaIcon.Bg(bg0)

	i.lblPlazma.SetText(i.plazma.String())
	i.lblPlazma.Bg(bg0)
	i.lblPlazma.Fg(fg0)

	i.lblMotto.SetText(i.plazma.MottoRus())
	i.lblMotto.Bg(bg0)
	i.lblMotto.Fg(fg0)
}

func (i *PlazmaIcon) Resize(r []int) {
	i.View.Resize(r)
	x := i.GetRect().X
	y := i.GetRect().Y
	sz := int(float64(i.GetRect().GetLowestSize()) / 16)
	i.plazmaIcon.Resize([]int{x + sz*3, y, sz * 10, sz * 10})
	i.lblPlazma.Resize([]int{x, y + sz*10, sz * 16, sz * 3})
	i.lblMotto.Resize([]int{x, y + sz*13, sz * 16, sz * 3})
}
