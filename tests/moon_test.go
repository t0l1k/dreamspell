package main

import (
	"testing"

	"github.com/t0l1k/dreamspell/lib/metonic"
)

func TestAgeFor1042000(t *testing.T) {
	got := metonic.GetMoonAge(metonic.GetDate("2000.04.10"))
	want := 6
	if got != want {
		t.Errorf("error test got:%v want:%v", got, want)
	}
}

func TestAge(t *testing.T) {
	data := map[string]int{
		"2000.04.10": 6,
		"1999.08.11": 30,
		"2007.08.28": 15,
		"2004.01.28": 7,
		"1961.02.15": 0,
		"1936.06.19": 0,
		"1898.01.22": 0,
		"1882.05.17": 0,
		"1793.09.05": 0,
		// "1699.09.23": 0, // 30
		// "1590.06.31": 0, // panic day out of range
	}
	for k, v := range data {
		got := metonic.GetMoonAge(metonic.GetDate(k))
		want := v
		if got != want {
			t.Errorf("for %v got %v want %v", k, got, want)
		}
	}
}
