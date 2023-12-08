package moon

import (
	"bytes"
	_ "embed"
	"image"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
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

var moonPngsInstance *moonImageCashe
var once sync.Once

func GetMoonPngs() *moonImageCashe {
	once.Do(func() {
		moonPngsInstance = newMoonPngsCache()
	})
	return moonPngsInstance
}

type moonImageCashe struct {
	cache map[lib.Ton]*ebiten.Image
}

func newMoonPngsCache() *moonImageCashe {
	log.Println("Генерация иконок тон луны")
	s := &moonImageCashe{}
	moonPngs := [][]byte{MagneticPng, LunarPng, ElectricPng, SelfExistingPng, OvertonePng, RhtythmicPng, ResonantPng, GalacticPng, SolarPng, PlanetaryPng, SpectralPng, CrystalPng, CosmicPng}

	s.cache = make(map[lib.Ton]*ebiten.Image)
	for _, ton := range lib.GetTons() {
		img, _, err := image.Decode(bytes.NewReader(moonPngs[int(ton)-1]))
		if err != nil {
			panic(err)
		}
		im := ebiten.NewImageFromImage(img)
		s.cache[ton] = im
	}
	return s
}

func (s *moonImageCashe) Get(ton lib.Ton) *ebiten.Image {
	if img, ok := s.cache[ton]; ok {
		return img
	}
	return nil
}

func (s *moonImageCashe) GetAll() (arr []*ebiten.Image) {
	for _, ton := range lib.GetTons() {
		arr = append(arr, s.cache[ton])
	}
	return arr
}
