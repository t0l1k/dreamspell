package main

import (
	"fmt"
	"testing"

	"github.com/t0l1k/dreamspell/lib"
)

func TestKin72(t *testing.T) {
	got := lib.NewConvert("1976.04.10").FindKin().GetNr()
	want := 72
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestKinLeapYearKin(t *testing.T) {
	data := map[string]int{
		"2024.02.29":       132,
		"2024.02.29 11:00": 131,
		"2024.02.29 12:00": 132,
	}
	for k, v := range data {
		got := lib.NewConvert(k).FindKin().GetNr()
		want := v
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestDayNrInlibYear2019(t *testing.T) {
	gots := []string{"2019.07.24", "2019.07.25", "2019.07.26", "2019.07.27", "2020.01.01", "2020.02.28", "2020.02.29"}
	wants := []int{364, 365, 1, 2, 160, 218, 218}
	for i, dt := range gots {
		got := lib.NewConvert(dt).FindDayInYear()
		want := wants[i]
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestDayNrInlibYear2020(t *testing.T) {
	gots := []string{"2020.07.24", "2020.07.25", "2020.07.26", "2020.07.27", "2021.01.01", "2021.02.28"}
	wants := []int{364, 365, 1, 2, 160, 218}
	for i, dt := range gots {
		got := lib.NewConvert(dt).FindDayInYear()
		want := wants[i]
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestDayNrInlibYear2021(t *testing.T) {
	gots := []string{"2021.07.24", "2021.07.25", "2021.07.26", "2021.07.27", "2022.02.28"}
	wants := []int{364, 365, 1, 2, 218}
	for i, dt := range gots {
		got := lib.NewConvert(dt).FindDayInYear()
		want := wants[i]
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestDayNrInlibYear2022(t *testing.T) {
	gots := []string{"2022.07.24", "2022.07.25", "2022.07.26", "2022.07.27", "2023.02.28"}
	wants := []int{364, 365, 1, 2, 218}
	for i, dt := range gots {
		got := lib.NewConvert(dt).FindDayInYear()
		want := wants[i]
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestDayNrInlibYear2023(t *testing.T) {
	gots := []string{"2023.07.24", "2023.07.25", "2023.07.26", "2023.07.27", "2024.02.28", "2024.02.29"}
	wants := []int{364, 365, 1, 2, 218, 218}
	for i, dt := range gots {
		got := lib.NewConvert(dt).FindDayInYear()
		want := wants[i]
		fmt.Println(got, want)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestDayNrInlibLeapYear2024(t *testing.T) {
	gots := []string{"2024.07.24", "2024.07.25", "2024.07.26", "2024.07.27", "2025.02.28"}
	wants := []int{364, 365, 1, 2, 218}
	for i, dt := range gots {
		got := lib.NewConvert(dt).FindDayInYear()
		want := wants[i]
		fmt.Println(got, want)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestDayNrInlibLeapYear2025(t *testing.T) {
	gots := []string{"2025.07.24", "2025.07.25", "2025.07.26", "2025.07.27", "2026.02.28"}
	wants := []int{364, 365, 1, 2, 218}
	for i, dt := range gots {
		got := lib.NewConvert(dt).FindDayInYear()
		want := wants[i]
		fmt.Println(got, want)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestDayNrInlibLeapYear2026(t *testing.T) {
	gots := []string{"2026.07.24", "2026.07.25", "2026.07.26", "2026.07.27", "2027.02.28"}
	wants := []int{364, 365, 1, 2, 218}
	for i, dt := range gots {
		got := lib.NewConvert(dt).FindDayInYear()
		want := wants[i]
		fmt.Println(got, want)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestMoonDayNr(t *testing.T) {
	gots := []string{"2023.07.24", "2023.07.25", "2023.07.26", "2023.07.27", "2024.01.01", "2024.02.28", "2024.02.29"}
	wants := []int{28, 0, 1, 2, 20, 22, 22}
	for i, dt := range gots {
		got := lib.NewConvert(dt).FindMoonDayNr()
		want := wants[i]
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestMoonNr(t *testing.T) {
	gots := []string{"2023.07.24", "2023.07.25", "2023.07.26", "2023.07.27", "2024.01.01", "2024.02.28", "2024.02.29"}
	wants := []int{13, 0, 1, 1, 6, 8, 8}
	for i, dt := range gots {
		got := lib.NewConvert(dt).FindMoonNr()
		want := wants[i]
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestNSStr(t *testing.T) {
	data := map[string]string{
		"1987.07.26": "NS1.0.1.1 Kin 34",
		"1988.07.24": "NS1.0.13.28 Kin 137",
		"1988.07.25": "NS1.1.0.0 Kin 138",
		"2007.07.25": "NS1.20.0.0 Kin 53",
		"1993.07.26": "NS1.6.1.1 Kin 144",
		"2007.07.26": "NS1.20.1.1 Kin 54",
		"2012.12.21": "NS1.25.6.9 Kin 207",
		"2013.07.26": "NS1.26.1.1 Kin 164",
		"2007.01.01": "NS1.19.6.20 Kin 108",
		"2053.01.24": "NS2.13.7.15 Kin 21",
		"1939.01.24": "-NS1.3.7.15 Kin 11",
		"1987.07.25": "-NS1.52.0.0 Kin 33",
		"1987.07.24": "-NS1.51.13.28 Kin 32",
		"2023.08.27": "NS1.36.2.5 Kin 206",
		"2023.09.25": "NS1.36.3.6 Kin 235",
		"2024.04.10": "NS1.36.10.7 Kin 172",
		"2024.05.08": "NS1.36.11.7 Kin 200",
		"2023.07.03": "NS1.35.13.7 Kin 151",
		"2023.11.12": "NS1.36.4.26 Kin 23",
		"2023.07.12": "NS1.35.13.16 Kin 160",
		"2023.06.26": "NS1.35.12.28 Kin 144", // fail
	}
	for k, v := range data {
		got := lib.NewConvert(k).GetNS()
		want := v
		if got != want {
			t.Errorf("for %v got %v want %v", k, got, want)
		}
	}
}

func TestMoonKin(t *testing.T) {
	data := map[string]int{
		"2012.03.07": 191,
		"2012.04.27": 192,
		"2012.05.20": 193,
		"2012.06.17": 194,
		"2012.07.24": 195,
		"2012.07.26": 196,
		"2012.09.01": 197,
		"2012.10.02": 198,
		"2012.11.15": 200,
		"2012.12.12": 200,
		"2012.11.20": 200,
		"2012.12.13": 201,
		"2012.12.18": 201,
		"2013.01.01": 201,
		"2013.01.24": 202,
		"2013.01.25": 202,
		"2013.02.06": 202,
		"2013.02.13": 203,
		"2013.07.24": 208,
		"2013.07.26": 209,
		"1997.07.26": 1,
		"1997.09.01": 2,
	}
	for k, v := range data {
		got := lib.NewConvert(k).FindMoonKin().GetNr()
		want := v
		if got != want {
			t.Errorf("[%v] got %v want %v", k, got, want)
		}
	}
}

func Test28and29Feb(t *testing.T) {
	gots := []string{"2024.02.28", "2024.02.29", "2023.02.28"}
	wants := []string{"28/29 Feb", "28/29 Feb", "28 Feb"}
	for i, dt := range gots {
		got := lib.NewConvert(dt).Strings()[0]
		want := wants[i]
		fmt.Println(got, want)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	gots = []string{"2024.02.28", "2024.02.29", "2023.02.28"}
	wants = []string{"Wed/Thu", "Wed/Thu", "Tue"}
	for i, dt := range gots {
		got := lib.NewConvert(dt).Strings()[1]
		want := wants[i]
		fmt.Println(got, want)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

}
