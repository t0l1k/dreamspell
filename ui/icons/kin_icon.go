package icons

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res/img/seals"
	"github.com/t0l1k/dreamspell/res/img/symbols"
	"github.com/t0l1k/dreamspell/res/img/tons"
	"github.com/t0l1k/eui"
)

type KinIcon struct {
	eui.View
	kin                     *lib.Kin
	nr                      *eui.Text
	ton, seal, signL, signR *eui.Icon
	sealOn                  bool
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

func (i *KinIcon) Setup(kin *lib.Kin) {
	i.SetupView()
	i.kin = kin
	clrs := []color.RGBA{eui.Red, eui.White, eui.Blue, eui.Yellow}
	nr := i.kin.GetColor() - 1
	var bg, fg color.RGBA
	if i.kin.IsPga() {
		bg = eui.Green
		fg = eui.Black
	} else {
		bg = clrs[nr]
		if bg == clrs[0] && bg == clrs[2] {
			fg = eui.White
		} else {
			fg = eui.Black
		}
	}
	if i.kin.IsCentral() {
		bg = clrs[nr]
		r := bg.R
		g := bg.G
		b := bg.B
		a := bg.A
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
		bg = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	}
	i.Bg(bg)

	tonImg := ebiten.NewImageFromImage(tons.GetTonPngs().Get(i.kin.GetTon()))
	if i.ton == nil {
		i.ton = eui.NewIcon(tonImg)
		i.Add(i.ton)
	} else {
		i.ton.SetIcon(tonImg)
	}
	i.ton.Bg(bg)

	sealImg := ebiten.NewImageFromImage(seals.GetSealPngs().Get(i.kin.GetSeal()))
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
		i.seal.Bg(bg)
		i.nr.SetText(strconv.Itoa(int(i.kin.GetNr())))
	} else if !i.sealOn {
		i.nr.SetText(strconv.Itoa(int(i.kin.GetNr())))
	}
	i.nr.Bg(bg)
	i.nr.Fg(fg)

	img := ebiten.NewImageFromImage(symbols.GetSignPngs()[3])
	if i.signL == nil {
		i.signL = eui.NewIcon(img)
		i.Add(i.signL)
	}
	if i.kin.IsClearSign() {
		i.signL.SetIcon(ebiten.NewImageFromImage(symbols.GetSignPngs()[0]))
	} else if i.kin.IsHiddenSign() {
		i.signL.SetIcon(ebiten.NewImageFromImage(symbols.GetSignPngs()[1]))
	} else {
		i.signL.SetIcon(img)
	}
	i.signL.Bg(bg)

	if i.signR == nil {
		i.signR = eui.NewIcon(img)
		i.Add(i.signR)
	}
	if i.kin.IsPolar() {
		i.signR.SetIcon(ebiten.NewImageFromImage(symbols.GetSignPngs()[2]))
	} else {
		i.signR.SetIcon(img)
	}
	i.signR.Bg(bg)
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
