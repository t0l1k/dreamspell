package app

import (
	"github.com/t0l1k/eui"
)

func NewGame() *eui.Ui {
	u := eui.GetUi()
	u.SetTitle("Dreamspell")
	u.SetSize(640, 400)
	return u
}
