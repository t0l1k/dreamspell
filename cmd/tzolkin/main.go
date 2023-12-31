package main

import (
	"github.com/t0l1k/dreamspell/ui/app"
	"github.com/t0l1k/dreamspell/ui/scene_tzolkin"
	"github.com/t0l1k/eui"
)

func main() {
	eui.Init(app.NewGame())
	eui.Run(scene_tzolkin.NewSceneTzolkin())
	eui.Quit()
}
