package metonic

import (
	"math"
	"time"
)

// Получить дату из строки в формате год.месяц.день
func GetDate(value string) time.Time {
	layout := "2006.01.02"
	dt, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return dt
}

// Определить используя метонов цикл возраст луны с поправками вычисляющий точно на несколько сотен лет вперед-назад, используется не 19 летний цикл луны, а 23х19=437 лет цикл синхронизации луны и солнца источник:ISSN 1607–2855,
func GetMoonAge(dt time.Time) int {
	l := getMoonNr(dt.Year())
	k := getKoef(dt.Month())
	result := l + int(dt.Month()) + k + dt.Day()
	for result > 30 {
		result -= 30
	}
	return result
}

func getMoonNr(year int) int {
	y0 := 1998      // Первый год 19 летнего цикла
	u := 21 / 437.0 // Вековая поправкв
	result := math.Mod(((11 + u) * float64(year-y0)), 30)
	return int(math.Round(result))
}

func getKoef(month time.Month) (result int) {
	switch month {
	case time.January:
		result = 2
	case time.February:
		result = 2
	case time.September:
		result = 1
	case time.November:
		result = 1
	}
	return result
}
