package app

import "github.com/t0l1k/eui"

func NewGame() *eui.Ui {
	u := eui.GetUi()
	theme := eui.NewTheme()
	theme.Set("bg", eui.Gray)
	u.ApplyTheme(&theme)
	loc := eui.NewLocale()
	loc.Set("lblUpTm", "Up")
	u.ApplyLocale(&loc)
	u.SetTitle("Dreamspell")
	return u
}
