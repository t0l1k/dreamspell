package lib

func contains(arr []int, value int) bool {
	for _, v := range arr {
		if value == v {
			return true
		}
	}
	return false
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
