package tons

import (
	"bytes"
	_ "embed"
	"image"
	"log"
	"sync"

	"github.com/t0l1k/dreamspell/lib"
)

var (
	//go:embed  01.png
	MagneticPng []byte
	//go:embed  02.png
	LunarPng []byte
	//go:embed  03.png
	ElectricPng []byte
	//go:embed  04.png
	SelfExistingPng []byte
	//go:embed  05.png
	OvertonePng []byte
	//go:embed  06.png
	RhtythmicPng []byte
	//go:embed  07.png
	ResonantPng []byte
	//go:embed  08.png
	GalacticPng []byte
	//go:embed  09.png
	SolarPng []byte
	//go:embed  10.png
	PlanetaryPng []byte
	//go:embed  11.png
	SpectralPng []byte
	//go:embed  12.png
	CrystalPng []byte
	//go:embed  13.png
	CosmicPng []byte
)

var tonPngsInstance *TonImageCashe
var once sync.Once

func GetTonPngs() *TonImageCashe {
	once.Do(func() {
		tonPngsInstance = newTonPngsCache()
	})
	return tonPngsInstance
}

type TonImageCashe struct {
	cache map[lib.Ton]*image.Image
}

func newTonPngsCache() *TonImageCashe {
	log.Println("Генерация иконок тонов")
	s := &TonImageCashe{}
	tonPngs := [][]byte{MagneticPng, LunarPng, ElectricPng, SelfExistingPng, OvertonePng, RhtythmicPng, ResonantPng, GalacticPng, SolarPng, PlanetaryPng, SpectralPng, CrystalPng, CosmicPng}
	s.cache = make(map[lib.Ton]*image.Image)
	for _, ton := range lib.GetTons() {
		img, _, err := image.Decode(bytes.NewReader(tonPngs[int(ton)-1]))
		if err != nil {
			panic(err)
		}
		s.cache[ton] = &img
	}
	return s
}

func (s *TonImageCashe) Get(ton lib.Ton) image.Image {
	if img, ok := s.cache[ton]; ok {
		return *img
	}
	return nil
}

func (s *TonImageCashe) GetAll() (arr []image.Image) {
	for _, ton := range lib.GetTons() {
		arr = append(arr, *s.cache[ton])
	}
	return arr
}

// func GetTonPngs() (arr []image.Image) {
// 	tonPngs := [][]byte{MagneticPng, LunarPng, ElectricPng, SelfExistingPng, OvertonePng, RhtythmicPng, ResonantPng, GalacticPng, SolarPng, PlanetaryPng, SpectralPng, CrystalPng, CosmicPng}
// 	for _, png := range tonPngs {
// 		img, _, err := image.Decode(bytes.NewReader(png))
// 		if err != nil {
// 			panic(err)
// 		}
// 		arr = append(arr, img)
// 	}
// 	return arr
// }
