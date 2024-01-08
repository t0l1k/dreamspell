package app

import (
	"image/color"

	"github.com/t0l1k/eui"
)

func NewGame() *eui.Ui {
	u := eui.GetUi()
	u.SetTitle("Dreamspell")
	u.SetSize(800, 600)
	u.GetTheme().Set(eui.SceneBg, color.White)
	return u
}
