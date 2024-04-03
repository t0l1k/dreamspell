package lib

import (
	"strconv"
	"time"
)

type Convert struct {
	tm time.Time
}

func NewConvert(dt string) *Convert {
	layout := "2006.01.02 15:04"
	if len(dt) <= 10 {
		dt += " 12:00"
	}
	tm, err := time.Parse(layout, dt)
	if err != nil {
		panic(err)
	}
	return &Convert{tm: tm}
}

func (c *Convert) GetDate() time.Time {
	return c.tm
}

func (c *Convert) FindMoonKin() *Kin {
	y1, m1, d1 := c.tm.Date()
	if m1 == time.July && d1 == 25 {
		return nil
	}
	beginYear := time.Date(y1, time.July, 25, 00, 00, 0, 0, time.Local)
	if c.tm.Before(beginYear) {
		y1--
	}
	moonNr := c.FindMoonNr()
	i := 0
	beginMoonKin := time.Date(1997, time.July, 26, 00, 00, 0, 0, time.Local)
	y0, _, _ := beginMoonKin.Date()
	for y1 > y0 {
		i += 13
		y1--
	}
	return GetTzolkin().GetKin(i + moonNr)
}

func (c *Convert) getNSCycle() (cycle, y int) {
	year, m, d := c.tm.Date()
	rootYear := 1987
	beginYear := time.Date(year, time.July, 26, 00, 00, 0, 0, time.Local)
	if c.tm.Before(beginYear) {
		year--
	}
	if year >= rootYear {
		for rootYear <= year {
			cycle++
			rootYear += 52
		}
		y = 52 - (rootYear - year)
	} else if year <= rootYear {
		for rootYear >= year {
			cycle--
			rootYear -= 52
		}
		y = (rootYear - year) * -1
	}
	if m == time.July && d == 25 {
		y++
	}
	return cycle, y
}
func (c *Convert) GetNSShort() string {
	s := "NS"
	cycle, year := c.getNSCycle()
	if cycle > 0 {
		s = "NS" + strconv.Itoa(cycle)
	} else {
		s = "-NS" + strconv.Itoa(cycle*-1)
	}
	s += "."
	s += strconv.Itoa(year)
	s += "."
	s += strconv.Itoa(c.FindMoonNr())
	return s
}

func (c *Convert) GetNS() string {
	s := c.GetNSShort()
	s += "."
	dayNr := c.FindMoonDayNr()
	s += strconv.Itoa(dayNr)
	s += " "
	kin := c.FindKin()
	s += "Kin " + strconv.Itoa(kin.GetNr())
	return s
}

func (c *Convert) FindDayPlazma() Plazma {
	if c.FindDayInYear() == 365 {
		return 0
	}
	nr := c.FindMoonDayNr() % 7
	if nr == 0 {
		nr = 7
	}
	return Plazma(nr)
}

func (c *Convert) FindMoonNr() int {
	dayInYear := c.FindDayInYear()
	if dayInYear == 365 {
		return 0
	}
	var moonNr int
	if dayInYear%28 == 0 {
		moonNr = (dayInYear - 1) / 28
	} else {
		moonNr = dayInYear / 28
	}
	return moonNr + 1
}

func (c *Convert) FindMoonDayNr() int {
	dayInYear := c.FindDayInYear()
	if dayInYear == 365 {
		return 0
	}
	moonDayNr := dayInYear % 28
	if moonDayNr == 0 {
		moonDayNr = 28
	}
	return moonDayNr
}

func (c *Convert) FindDayInYear() int {
	dayBegindreamspellYear := c.FindDreamspellYearBeginDate()
	yearDayNr := c.tm.YearDay()
	if isLeapYear(c.tm.Year()) {
		// Если высокосный год, пропускаем счет дней до 29 февраля, находим 29 февраля, его пропускаем и считаем -1 день далее до конца года, сравнивая день в григорианском календаре, и это в високосный год по григорианскому календарю
		dayYearEnd := time.Date(c.tm.Year(), time.December, 31, 0, 0, 0, 0, time.Local)
		leapDay := time.Date(c.tm.Year(), time.February, 29, 0, 0, 0, 0, time.Local)
		if c.tm.After(leapDay) && c.tm.Before(dayYearEnd) && dayBegindreamspellYear.Year() != c.tm.Year() {
			yearDayNr--
		}
	} else if isLeapYear(dayBegindreamspellYear.Year()) {
		// Если високостный год от начала майянского года после нового григоранского года увеличиваем счет дней, до конца майянского года
		yearDayNr++
	}
	result := yearDayNr - dayBegindreamspellYear.YearDay()
	if result < 0 {
		result += 366
	} else if result > 0 {
		result++
	} else {
		result += 1
	}
	return result
}

func (c *Convert) FindYearKin() *Kin {
	layout := "2006.01.02"
	dtBeginYear := c.FindDreamspellYearBeginDate()
	return NewConvert(dtBeginYear.Format(layout)).FindKin()
}

func (c *Convert) FindDreamspellYearBeginDate() time.Time {
	layout := "2006.01.02"
	dt := strconv.Itoa(c.tm.Year()) + ".07.26"
	dtBeginYear, err := time.Parse(layout, dt)
	if err != nil {
		panic(err)
	}
	if c.tm.Before(dtBeginYear) {
		y := c.tm.Year() - 1
		dt := strconv.Itoa(y) + ".07.26"
		dtBeginYear, err = time.Parse(layout, dt)
		if err != nil {
			panic(err)
		}
	}
	return dtBeginYear
}

func (c *Convert) FindKin() *Kin {
	yearKin := func() *Kin {
		ton := PLANETARY
		seal := WIND
		y := 1858
		for y < c.tm.Year() {
			ton.Inc()
			seal += 5
			if seal > STORM {
				seal = WIND
			}
			y++
		}
		return GetTzolkin().GetKinByTonSeal(ton, seal)
	}()
	moonKin := func() int {
		arr := []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 13, 44, 74}
		return arr[c.tm.Month()-1]
	}()
	nr := yearKin.nr + moonKin + c.tm.Day()
	_, m, d := c.tm.Date()
	if m == time.February && d == 29 && c.tm.Hour() < 12 {
		nr--
	}
	return GetTzolkin().GetKin(nr)
}

var weekDaysShort = []string{"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"}

func (c Convert) Strings() []string {
	y, m, d := c.tm.Date()
	dtWeekDay := weekDaysShort[c.tm.Weekday()]
	if isLeapYear(y) && m == time.February && d >= 28 {
		s := dtWeekDay
		switch d {
		case 28:
			dt := c.tm.Add(time.Duration(time.Hour * 24))
			dtWeekDay2 := dt.Weekday().String()[:3]
			s += "/" + dtWeekDay2
		case 29:
			dt := c.tm.Add(time.Duration(time.Hour * -24))
			dtWeekDay2 := dt.Weekday().String()[:3]
			s = dtWeekDay2 + "/" + dtWeekDay
		}
		return []string{"28/29 Feb", s}
	}
	return []string{c.tm.Format("02 Jan"), dtWeekDay}
}
