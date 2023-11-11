package lib

import (
	"strconv"
)

type Kin struct {
	nr                                         int
	ton                                        Ton
	seal                                       Seal
	color                                      SealColor
	central, pga, clearSign, hiddenSign, polar bool
}

func newKin(nr int) *Kin {
	for nr > 260 {
		nr -= 260
	}
	pgas := []int{1, 20, 22, 39, 43, 50, 51, 58, 64, 69, 72, 77, 85, 93, 96, 106,
		107, 108, 109, 110, 111, 112, 113, 114, 115, 146, 147, 148,
		149, 150, 151, 152, 153, 154, 155, 165, 168, 173, 176, 184,
		189, 192, 197, 203, 210, 211, 218, 222, 239, 241, 260}
	polars := []int{185, 205, 225, 245, 10, 30, 50, 250, 55, 75, 95, 115, 120,
		140, 160, 180}
	clears := []int{20, 26, 30, 40, 57, 58, 60, 87, 106, 132, 176, 211, 245}
	k := &Kin{nr: nr}
	if nr%4 == 0 {
		k.color = YELLOW
	} else {
		k.color = SealColor(nr % 4)
	}
	k.seal = Seal(nr % 20)
	if nr%13 == 0 {
		k.ton = 13
	} else {
		k.ton = Ton(nr % 13)
	}
	k.central = func() bool {
		return nr >= 121 && nr <= 140
	}()
	k.pga = contains(pgas, nr)
	k.polar = contains(polars, nr)
	k.clearSign = contains(clears, nr)
	k.hiddenSign = func() bool {
		return nr == 78
	}()
	return k
}

func (k *Kin) GetNr() int {
	return k.nr
}

func (k *Kin) GetTon() Ton {
	return k.ton
}

func (k *Kin) GetSeal() Seal {
	return k.seal
}

func (k *Kin) GetColor() SealColor {
	return k.color
}

func (k *Kin) IsCentral() bool {
	return k.central
}

func (k *Kin) IsPga() bool {
	return k.pga
}

func (k *Kin) IsPolar() bool {
	return k.polar
}

func (k *Kin) IsClearSign() bool {
	return k.clearSign
}

func (k *Kin) IsHiddenSign() bool {
	return k.hiddenSign
}

func (k *Kin) String() string {
	s := strconv.Itoa(k.nr) + " " + k.ton.String() + " " + k.color.String() + " " + k.seal.String()
	if k.pga {
		s += " PGA"
	}
	if k.polar {
		s += " POLAR"
	}
	if k.clearSign {
		s += " CLEAR SIGN"
	}
	if k.hiddenSign {
		s += " HIDDEN SIGN"
	}
	if k.central {
		s += " CENTRAL"
	}
	return s
}
