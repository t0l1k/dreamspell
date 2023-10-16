package lib

import "sync"

var instance *tzolkin

func GetTzolkin() *tzolkin {
	if instance == nil {
		instance = newTzolkin()
	}
	return instance
}

type tzolkin struct {
	kins []*Kin
	sync.Mutex
}

func newTzolkin() *tzolkin {
	tz := &tzolkin{}
	tz.Lock()
	defer tz.Unlock()
	nr := 1
	for nr <= 260 {
		tz.kins = append(tz.kins, newKin(nr))
		nr++
	}
	return tz
}

func (tz *tzolkin) GetAll() []*Kin {
	tz.Lock()
	defer tz.Unlock()
	return tz.kins
}

func (tz *tzolkin) GetKin(nr int) *Kin {
	tz.Lock()
	defer tz.Unlock()
	for nr > 260 {
		nr -= 260
	}
	return tz.kins[nr-1]
}

func (tz *tzolkin) GetKinByTonSeal(ton Ton, seal Seal) *Kin {
	tz.Lock()
	defer tz.Unlock()
	for _, kin := range tz.kins {
		if kin.ton == ton && kin.seal == seal {
			return kin
		}
	}
	return nil
}
