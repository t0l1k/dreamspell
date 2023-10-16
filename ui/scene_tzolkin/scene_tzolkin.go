package scene_tzolkin

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res/img/seals"
	"github.com/t0l1k/eui"
)

type SceneTzolkin struct {
	seals []*eui.Icon
	kins  []*TzolkinKinIcon
	eui.ContainerDefault
}

func NewSceneTzolkin() *SceneTzolkin {
	sc := &SceneTzolkin{}
	rect := []int{0, 0, 1, 1}
	for _, img := range seals.GetSealPngs().GetAll() {
		seal := eui.NewIcon(ebiten.NewImageFromImage(img), rect)
		sc.seals = append(sc.seals, seal)
		sc.Add(seal)
	}
	for _, kin := range lib.GetTzolkin().GetAll() {
		kinIcon := NewTzolkinKinIcon(kin, rect)
		sc.kins = append(sc.kins, kinIcon)
		sc.Add(kinIcon)
	}

	return sc
}

func (sc *SceneTzolkin) Entered() {
	sc.Resize()
}
func (sc *SceneTzolkin) Update(dt int) {
	for _, v := range sc.Container {
		v.Update(dt)
	}
}

func (sc *SceneTzolkin) Draw(surface *ebiten.Image) {
	for _, v := range sc.Container {
		v.Draw(surface)
	}
}

func (sc *SceneTzolkin) Resize() {
	w, h := ebiten.WindowSize()
	sz := 20
	size := w
	if w > h {
		size = h
	} else {
		size = w
	}
	cellSize := size / sz
	x0 := (w - cellSize*14) / 2
	y0 := (h - cellSize*20) / 2
	x, y := x0, y0
	w1 := cellSize
	h1 := cellSize
	for _, icon := range sc.seals {
		icon.Resize([]int{x, y, w1, h1})
		y += cellSize
	}
	y = y0
	x += cellSize
	i := 0
	for _, icon := range sc.kins {
		icon.Resize([]int{x, y, w1, h1})
		y += cellSize
		i++
		if i > 0 && i%20 == 0 {
			y = y0
			x += cellSize
		}
	}
}

func (sc *SceneTzolkin) Close() {
	for _, v := range sc.Container {
		v.Close()
	}
}
