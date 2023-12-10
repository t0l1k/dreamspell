package scene_tzolkin

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res"
	"github.com/t0l1k/dreamspell/ui/icons"
	"github.com/t0l1k/eui"
)

type SceneTzolkin struct {
	seals []*eui.Icon
	kins  []*icons.KinIcon
	eui.SceneBase
	layout *eui.GridLayoutDownRight
}

func NewSceneTzolkin() *SceneTzolkin {
	sc := &SceneTzolkin{}
	sc.layout = eui.NewGridLayoutDownRight(14, 20)
	sc.layout.SetCellMargin(1)
	for _, seal := range lib.GetSeals() {
		if seal == lib.SUN {
			continue
		}
		img := res.GetSealAll()[seal]
		icon := eui.NewIcon(img)
		sc.seals = append(sc.seals, icon)
		sc.layout.Add(icon)
	}
	img := res.GetSealAll()[lib.SUN]
	icon := eui.NewIcon(img)
	sc.seals = append(sc.seals, icon)
	sc.layout.Add(icon)

	for _, kin := range lib.GetTzolkin().GetAll() {
		kinIcon := icons.NewKinNrIcon(kin)
		sc.kins = append(sc.kins, kinIcon)
		sc.layout.Add(kinIcon)
	}
	sc.Add(sc.layout)
	sc.Resize()
	return sc
}

func (s *SceneTzolkin) Update(dt int) {
	for _, v := range s.Container {
		v.Update(dt)
		vv, ok := v.(*eui.GridLayoutDownRight)
		if ok {
			for _, v := range vv.Container {
				v.Update(dt)
			}
		}
	}
}

func (s *SceneTzolkin) Draw(surface *ebiten.Image) {
	for _, v := range s.Container {
		v.Draw(surface)
		vv, ok := v.(*eui.GridLayoutDownRight)
		if ok {
			for _, v := range vv.Container {
				v.Draw(surface)
			}
		}
	}
}

func (sc *SceneTzolkin) Resize() {
	w0, h0 := eui.GetUi().Size()
	sc.layout.Resize([]int{0, 0, w0, h0})
}
