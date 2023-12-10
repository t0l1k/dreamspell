package symbols

import (
	_ "embed"
)

var (
	//go:embed  star.png
	ClearSignSymbolPng []byte
	//go:embed  hidden.png
	HiddenSignSymbolPng []byte
	//go:embed  romb.png
	PolarSymbolPng []byte
	//go:embed  empty.png
	EmptySymbolPng []byte
	//go:embed  crystal.png
	CrystalPng []byte
)
