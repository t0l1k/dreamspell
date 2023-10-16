package lib

type SealColor int

const (
	RED SealColor = iota + 1
	WHITE
	BLUE
	YELLOW
)

func (c SealColor) String() string {
	arr := []string{
		"RED",
		"WHITE",
		"BLUE",
		"YELLOW"}
	return arr[int(int(c)-1)]
}
