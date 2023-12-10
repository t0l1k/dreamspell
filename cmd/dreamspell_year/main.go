package main

import (
	"github.com/t0l1k/dreamspell/ui/app"
	scene_dreamspell_year "github.com/t0l1k/dreamspell/ui/scene_year"
	"github.com/t0l1k/eui"
)

func main() {
	eui.Init(app.NewGame())
	eui.Run(scene_dreamspell_year.NewSceneDreamSpellYear())
	eui.Quit()
}
