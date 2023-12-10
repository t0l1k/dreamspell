package res

import (
	"bytes"
	"io"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/t0l1k/dreamspell/lib"
	"github.com/t0l1k/dreamspell/res/img/moon"
	"github.com/t0l1k/dreamspell/res/img/plazmas"
	"github.com/t0l1k/dreamspell/res/img/seals"
	"github.com/t0l1k/dreamspell/res/img/symbols"
	"github.com/t0l1k/dreamspell/res/img/tons"
)

var instance *resource.Loader
var once sync.Once

func getResourceInstance() *resource.Loader {
	once.Do(func() {
		instance = loadRes()
	})
	return instance
}

const (
	SunImg resource.ImageID = iota + 20
	DragonImg
	WindImg
	NightImg
	SeedImg
	SerpentImg
	WorldBridgerImg
	HandImg
	StarImg
	MoonImg
	DogImg
	MonkeyImg
	HumanImg
	SkyWalkerImg
	WizzardImg
	EagleImg
	WarriorImg
	EarthImg
	MirrorImg
	StormImg
)

const (
	MagneticImg resource.ImageID = iota + 1
	LunarImg
	ElectricImg
	SelfExistingImg
	OvertoneImg
	RhtythmicImg
	ResonantImg
	GalacticImg
	SolarImg
	PlanetaryImg
	SpectralImg
	CrystalImg
	CosmicImg
)

const (
	DaliImg resource.ImageID = iota + 70
	SeliImg
	GammaImg
	KaliImg
	AlphaImg
	LimiImg
	SilioImg
)

const (
	ClearSignSymbolImg resource.ImageID = iota + 50
	HiddenSignSymbolImg
	PolarSymbolImg
	EmptySymbolImg
)

const (
	MagneticMoonImg resource.ImageID = iota + 91
	LunarMoonImg
	ElectricMoonImg
	SelfExistingMoonImg
	OvertoneMoonImg
	RhtythmicMoonImg
	ResonantMoonImg
	GalacticMoonImg
	SolarMoonImg
	PlanetaryMoonImg
	SpectralMoonImg
	CrystalMoonImg
	CosmicMoonImg
)

func loadRes() *resource.Loader {
	audioContext := audio.NewContext(44100)

	l := resource.NewLoader(audioContext)

	l.OpenAssetFunc = func(path string) io.ReadCloser {
		return io.NopCloser(bytes.NewReader(resdata[path]))
	}

	tonImgs := map[resource.ImageID]resource.ImageInfo{
		MagneticImg:     {Path: "tons/01.png"},
		LunarImg:        {Path: "tons/02.png"},
		ElectricImg:     {Path: "tons/03.png"},
		SelfExistingImg: {Path: "tons/04.png"},
		OvertoneImg:     {Path: "tons/05.png"},
		RhtythmicImg:    {Path: "tons/06.png"},
		ResonantImg:     {Path: "tons/07.png"},
		GalacticImg:     {Path: "tons/08.png"},
		SolarImg:        {Path: "tons/09.png"},
		PlanetaryImg:    {Path: "tons/10.png"},
		SpectralImg:     {Path: "tons/11.png"},
		CrystalImg:      {Path: "tons/12.png"},
		CosmicImg:       {Path: "tons/13.png"},

		SunImg:          {Path: "seals/00.png"},
		DragonImg:       {Path: "seals/01.png"},
		WindImg:         {Path: "seals/02.png"},
		NightImg:        {Path: "seals/03.png"},
		SeedImg:         {Path: "seals/04.png"},
		SerpentImg:      {Path: "seals/05.png"},
		WorldBridgerImg: {Path: "seals/06.png"},
		HandImg:         {Path: "seals/07.png"},
		StarImg:         {Path: "seals/08.png"},
		MoonImg:         {Path: "seals/09.png"},
		DogImg:          {Path: "seals/10.png"},
		MonkeyImg:       {Path: "seals/11.png"},
		HumanImg:        {Path: "seals/12.png"},
		SkyWalkerImg:    {Path: "seals/13.png"},
		WizzardImg:      {Path: "seals/14.png"},
		EagleImg:        {Path: "seals/15.png"},
		WarriorImg:      {Path: "seals/16.png"},
		EarthImg:        {Path: "seals/17.png"},
		MirrorImg:       {Path: "seals/18.png"},
		StormImg:        {Path: "seals/19.png"},

		DaliImg:  {Path: "plazmas/01.png"},
		SeliImg:  {Path: "plazmas/02.png"},
		GammaImg: {Path: "plazmas/03.png"},
		KaliImg:  {Path: "plazmas/04.png"},
		AlphaImg: {Path: "plazmas/05.png"},
		LimiImg:  {Path: "plazmas/06.png"},
		SilioImg: {Path: "plazmas/07.png"},

		ClearSignSymbolImg:  {Path: "symbols/star.png"},
		HiddenSignSymbolImg: {Path: "symbols/hidden.png"},
		PolarSymbolImg:      {Path: "symbols/romb.png"},
		EmptySymbolImg:      {Path: "symbols/empty.png"},

		MagneticMoonImg:     {Path: "moon/01.png"},
		LunarMoonImg:        {Path: "moon/02.png"},
		ElectricMoonImg:     {Path: "moon/03.png"},
		SelfExistingMoonImg: {Path: "moon/04.png"},
		OvertoneMoonImg:     {Path: "moon/05.png"},
		RhtythmicMoonImg:    {Path: "moon/06.png"},
		ResonantMoonImg:     {Path: "moon/07.png"},
		GalacticMoonImg:     {Path: "moon/08.png"},
		SolarMoonImg:        {Path: "moon/09.png"},
		PlanetaryMoonImg:    {Path: "moon/10.png"},
		SpectralMoonImg:     {Path: "moon/11.png"},
		CrystalMoonImg:      {Path: "moon/12.png"},
		CosmicMoonImg:       {Path: "moon/13.png"},
	}

	l.ImageRegistry.Assign(tonImgs)
	return l
}

func GetTonAll() (arr []*ebiten.Image) {
	l := getResourceInstance()
	for _, ton := range lib.GetTons() {
		img := l.LoadImage(resource.ImageID(ton)).Data
		arr = append(arr, img)
	}
	return arr
}

func GetSealAll() (arr []*ebiten.Image) {
	l := getResourceInstance()
	for _, seal := range lib.GetSeals() {
		img := l.LoadImage(resource.ImageID(seal + 20)).Data
		arr = append(arr, img)
	}
	return arr
}

func GetPlazmaAll() (arr []*ebiten.Image) {
	l := getResourceInstance()
	for _, plazma := range lib.GetPlazmas() {
		img := l.LoadImage(resource.ImageID(plazma - 1 + 70)).Data
		arr = append(arr, img)
	}
	return arr
}

func GetSymbolAll() (arr []*ebiten.Image) {
	l := getResourceInstance()
	ids := []resource.ImageID{ClearSignSymbolImg, HiddenSignSymbolImg, PolarSymbolImg, EmptySymbolImg}
	for _, id := range ids {
		img := l.LoadImage(resource.ImageID(id)).Data
		arr = append(arr, img)
	}
	return arr
}

func GetMoonTonAll() (arr []*ebiten.Image) {
	l := getResourceInstance()
	for _, ton := range lib.GetTons() {
		img := l.LoadImage(resource.ImageID(ton - 1 + 91)).Data
		arr = append(arr, img)
	}
	return arr
}

var resdata = map[string][]byte{
	"tons/01.png": tons.MagneticPng,
	"tons/02.png": tons.LunarPng,
	"tons/03.png": tons.ElectricPng,
	"tons/04.png": tons.SelfExistingPng,
	"tons/05.png": tons.OvertonePng,
	"tons/06.png": tons.RhtythmicPng,
	"tons/07.png": tons.ResonantPng,
	"tons/08.png": tons.GalacticPng,
	"tons/09.png": tons.SolarPng,
	"tons/10.png": tons.PlanetaryPng,
	"tons/11.png": tons.SpectralPng,
	"tons/12.png": tons.CrystalPng,
	"tons/13.png": tons.CosmicPng,

	"seals/00.png": seals.SunPng,
	"seals/01.png": seals.DragonPng,
	"seals/02.png": seals.WindPng,
	"seals/03.png": seals.NightPng,
	"seals/04.png": seals.SeedPng,
	"seals/05.png": seals.SerpentPng,
	"seals/06.png": seals.WorldBridgerPng,
	"seals/07.png": seals.HandPng,
	"seals/08.png": seals.StarPng,
	"seals/09.png": seals.MoonPng,
	"seals/10.png": seals.DogPng,
	"seals/11.png": seals.MonkeyPng,
	"seals/12.png": seals.HumanPng,
	"seals/13.png": seals.SkyWalkerPng,
	"seals/14.png": seals.WizzardPng,
	"seals/15.png": seals.EaglePng,
	"seals/16.png": seals.WarriorPng,
	"seals/17.png": seals.EarthPng,
	"seals/18.png": seals.MirrorPng,
	"seals/19.png": seals.StormPng,

	"plazmas/01.png": plazmas.Dali,
	"plazmas/02.png": plazmas.Seli,
	"plazmas/03.png": plazmas.Gamma,
	"plazmas/04.png": plazmas.Kali,
	"plazmas/05.png": plazmas.Alpha,
	"plazmas/06.png": plazmas.Limi,
	"plazmas/07.png": plazmas.Silio,

	"symbols/star.png":   symbols.ClearSignSymbolPng,
	"symbols/romb.png":   symbols.PolarSymbolPng,
	"symbols/hidden.png": symbols.HiddenSignSymbolPng,
	"symbols/empty.png":  symbols.EmptySymbolPng,

	"moon/01.png": moon.MagneticPng,
	"moon/02.png": moon.LunarPng,
	"moon/03.png": moon.ElectricPng,
	"moon/04.png": moon.SelfExistingPng,
	"moon/05.png": moon.OvertonePng,
	"moon/06.png": moon.RhtythmicPng,
	"moon/07.png": moon.ResonantPng,
	"moon/08.png": moon.GalacticPng,
	"moon/09.png": moon.SolarPng,
	"moon/10.png": moon.PlanetaryPng,
	"moon/11.png": moon.SpectralPng,
	"moon/12.png": moon.CrystalPng,
	"moon/13.png": moon.CosmicPng,
}
