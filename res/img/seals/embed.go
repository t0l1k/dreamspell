package seals

import (
	"bytes"
	_ "embed"
	"image"
	"log"
	"sync"

	"github.com/t0l1k/dreamspell/lib"
)

var (
	//go:embed  00.png
	SunPng []byte
	//go:embed  01.png
	DragonPng []byte
	//go:embed  02.png
	WindPng []byte
	//go:embed  03.png
	NightPng []byte
	//go:embed  04.png
	SeedPng []byte
	//go:embed  05.png
	SerpentPng []byte
	//go:embed  06.png
	WorldBridgerPng []byte
	//go:embed  07.png
	HandPng []byte
	//go:embed  08.png
	StarPng []byte
	//go:embed  09.png
	MoonPng []byte
	//go:embed  10.png
	DogPng []byte
	//go:embed  11.png
	MonkeyPng []byte
	//go:embed  12.png
	HumanPng []byte
	//go:embed  13.png
	SkyWalkerPng []byte
	//go:embed  14.png
	WizzardPng []byte
	//go:embed  15.png
	EaglePng []byte
	//go:embed  16.png
	WarriorPng []byte
	//go:embed  17.png
	EarthPng []byte
	//go:embed  18.png
	MirrorPng []byte
	//go:embed  19.png
	StormPng []byte
)

var sealPngsInstance *SealsImageCashe
var once sync.Once

func GetSealPngs() *SealsImageCashe {
	once.Do(func() {
		sealPngsInstance = newSealPngsCache()
	})
	return sealPngsInstance
}

type SealsImageCashe struct {
	cache map[lib.Seal]*image.Image
}

func newSealPngsCache() *SealsImageCashe {
	log.Println("Генерация иконок печатей")
	s := &SealsImageCashe{}
	pngs := [][]byte{SunPng, DragonPng, WindPng, NightPng, SeedPng, SerpentPng, WorldBridgerPng, HandPng, StarPng, MoonPng, DogPng, MonkeyPng, HumanPng, SkyWalkerPng, WizzardPng, EaglePng, WarriorPng, EarthPng, MirrorPng, StormPng}
	s.cache = make(map[lib.Seal]*image.Image)
	for _, seal := range lib.GetSeals() {
		img, _, err := image.Decode(bytes.NewReader(pngs[int(seal)]))
		if err != nil {
			panic(err)
		}
		s.cache[seal] = &img
	}
	return s
}

func (s *SealsImageCashe) Get(seal lib.Seal) image.Image {
	if img, ok := s.cache[seal]; ok {
		return *img
	}
	return nil
}

func (s *SealsImageCashe) GetAll() (arr []image.Image) {
	for _, seal := range lib.GetSeals() {
		if seal == lib.SUN {
			continue
		}
		arr = append(arr, *s.cache[seal])
	}
	arr = append(arr, *s.cache[lib.SUN])
	return arr
}
