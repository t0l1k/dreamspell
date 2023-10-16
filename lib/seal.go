package lib

type Seal int

const (
	SUN Seal = iota
	DRAGON
	WIND
	NIGHT
	SEED
	SERPENT
	WORLD_BRIDGER
	HAND
	STAR
	MOON
	DOG
	MONKEY
	HUMAN
	SKYWALKER
	WIZARD
	EAGLE
	WARRIOR
	EARTH
	MIRROR
	STORM
)

func GetSeals() []Seal {
	return []Seal{SUN, DRAGON, WIND, NIGHT, SEED, SERPENT, WORLD_BRIDGER, HAND, STAR, MOON, DOG, MONKEY, HUMAN, SKYWALKER, WIZARD, EAGLE, WARRIOR, EARTH, MIRROR, STORM}
}

func (s Seal) String() string {
	arr := []string{
		"SUN",
		"DRAGON",
		"WIND",
		"NIGHT",
		"SEED",
		"SERPENT",
		"WORLD BRIDGER",
		"HAND",
		"STAR",
		"MOON",
		"DOG",
		"MONKEY",
		"HUMAN",
		"SKYWALKER",
		"WIZARD",
		"EAGLE",
		"WARRIOR",
		"EARTH",
		"MIRROR",
		"STORM"}
	return arr[int(s)]
}
