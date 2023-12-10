package icons

import (
	"image/color"
	"strconv"

	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res"
	"github.com/t0l1k/eui"
)

type KinIcon struct {
	eui.View
	kin                     *lib.Kin
	nr                      *eui.Text
	ton, seal, signL, signR *eui.Icon
	sealOn                  bool
	bg, fg                  color.RGBA
}

func NewKinNrIcon(kin *lib.Kin) *KinIcon {
	i := &KinIcon{}
	i.sealOn = false
	i.Setup(kin)
	return i
}

func NewKinSealIcon(kin *lib.Kin) *KinIcon {
	i := &KinIcon{}
	i.sealOn = true
	i.Setup(kin)
	return i
}

func (i *KinIcon) GetBg() color.Color {
	return i.bg
}

func (i *KinIcon) GetFg() color.Color {
	return i.fg
}

func (i *KinIcon) Setup(kin *lib.Kin) {
	i.SetupView()
	i.kin = kin
	i.setColors()
	i.Bg(i.bg)

	tonImg := res.GetTonAll()[i.kin.GetTon()-1]
	if i.ton == nil {
		i.ton = eui.NewIcon(tonImg)
		i.Add(i.ton)
	} else {
		i.ton.SetIcon(tonImg)
	}
	i.ton.Bg(i.bg)

	sealImg := res.GetSealAll()[i.kin.GetSeal()]
	if i.sealOn && i.seal == nil {
		i.seal = eui.NewIcon(sealImg)
		i.Add(i.seal)
	}
	if i.nr == nil {
		i.nr = eui.NewText("")
		i.Add(i.nr)
	}

	if i.sealOn {
		i.seal.SetIcon(sealImg)
		i.seal.Bg(i.bg)
		i.nr.SetText(strconv.Itoa(int(i.kin.GetNr())))
	} else if !i.sealOn {
		i.nr.SetText(strconv.Itoa(int(i.kin.GetNr())))
	}
	i.nr.Bg(i.bg)
	i.nr.Fg(i.fg)

	img := res.GetSymbolAll()[3]
	if i.signL == nil {
		i.signL = eui.NewIcon(img)
		i.Add(i.signL)
	}
	if i.kin.IsClearSign() {
		i.signL.SetIcon(res.GetSymbolAll()[0])
	} else if i.kin.IsHiddenSign() {
		i.signL.SetIcon(res.GetSymbolAll()[1])
	} else {
		i.signL.SetIcon(img)
	}
	i.signL.Bg(i.bg)

	if i.signR == nil {
		i.signR = eui.NewIcon(img)
		i.Add(i.signR)
	}
	if i.kin.IsPolar() {
		i.signR.SetIcon(res.GetSymbolAll()[2])
	} else {
		i.signR.SetIcon(img)
	}
	i.signR.Bg(i.bg)
}

func (i *KinIcon) setColors() {
	clrs := []color.RGBA{eui.Red, eui.White, eui.Blue, eui.Yellow}
	nr := i.kin.GetColor() - 1
	if i.kin.IsPga() {
		i.bg = eui.Green
		i.fg = eui.Black
	} else {
		i.bg = clrs[nr]
		if i.bg == clrs[0] && i.bg == clrs[2] {
			i.fg = eui.White
		} else {
			i.fg = eui.Black
		}
	}
	if i.kin.IsCentral() {
		i.bg = clrs[nr]
		r := i.bg.R
		g := i.bg.G
		b := i.bg.B
		a := i.bg.A
		v := uint8(64)
		if r == 255 {
			r -= v
		}
		if g == 255 {
			g -= v
		}
		if b == 255 {
			b -= v
		}
		i.bg = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	}
}

func (i *KinIcon) Resize(rect []int) {
	i.View.Resize(rect)
	m := 8 * 16.0
	sz := float64(i.GetRect().GetLowestSize()) / m
	x, y := i.GetRect().Pos()
	marginX := (float64(i.GetRect().W) - sz*m) / 2
	marginY := (float64(i.GetRect().H) - sz*m) / 2
	x += int(marginX)
	y += int(marginY)
	i.ton.Resize([]int{x + int(sz*(m/4)), y, int(sz * (m / 2)), int(sz * (m / 4))})
	i.signL.Resize([]int{x, y, int(sz * (m / 4)), int(sz * (m / 4))})
	i.signR.Resize([]int{x + int(sz*(m-(m/4))), y, int(sz * (m / 4)), int(sz * (m / 4))})
	if i.sealOn {
		i.seal.Resize([]int{x + int(sz*(m/4)), y + int(sz*(m/4)), int(sz * (m / 2)), int(sz * (m / 2))})
		i.nr.Resize([]int{x + int(sz*(m/4)), y + int(sz*(m-(m/4))), int(sz * (m / 2)), int(sz * (m / 4))})
	} else {
		i.nr.Resize([]int{x, y + int(sz*(m/4)), int(sz * m), int(sz * (m - (m / 4)))})
	}
}
