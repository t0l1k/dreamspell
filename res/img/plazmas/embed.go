package plazmas

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
	Dali []byte
	//go:embed  02.png
	Seli []byte
	//go:embed  03.png
	Gamma []byte
	//go:embed  04.png
	Kali []byte
	//go:embed  05.png
	Alpha []byte
	//go:embed  06.png
	Limi []byte
	//go:embed  07.png
	Silio []byte
)

var plazmaPngsInstance *plazmaImageCashe
var once sync.Once

func GetPlazmaPngs() *plazmaImageCashe {
	once.Do(func() {
		plazmaPngsInstance = newTonPngsCache()
	})
	return plazmaPngsInstance
}

type plazmaImageCashe struct {
	cache map[lib.Plazma]*ebiten.Image
}

func newTonPngsCache() *plazmaImageCashe {
	log.Println("Генерация иконок плазм")
	s := &plazmaImageCashe{}
	plazmaPngs := [][]byte{Dali, Seli, Gamma, Kali, Alpha, Limi, Silio}
	s.cache = make(map[lib.Plazma]*ebiten.Image)
	for _, plazma := range lib.GetPlazmas() {
		img, _, err := image.Decode(bytes.NewReader(plazmaPngs[int(plazma)-1]))
		if err != nil {
			panic(err)
		}
		im := ebiten.NewImageFromImage(img)
		s.cache[plazma] = im
	}
	return s
}

func (s *plazmaImageCashe) Get(plazma lib.Plazma) *ebiten.Image {
	if img, ok := s.cache[plazma]; ok {
		return img
	}
	return nil
}

func (s *plazmaImageCashe) GetAll() (arr []*ebiten.Image) {
	for _, plazma := range lib.GetPlazmas() {
		arr = append(arr, s.cache[plazma])
	}
	return arr
}
