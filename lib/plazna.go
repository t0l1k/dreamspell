package lib

type Plazma int

const (
	DALI Plazma = iota + 1
	SELI
	GAMMA
	KALI
	ALPHA
	LIMI
	SILIO
)

func GetPlazmas() []Plazma {
	return []Plazma{DALI, SELI, GAMMA, KALI, ALPHA, LIMI, SILIO}
}

func (p Plazma) String() string {
	arr := []string{
		"DAY OUT OF TIME",
		"DALI",  // Target Цель
		"SELI",  // Flow Поток
		"GAMMA", // Pacify Усмирять
		"KALI",  // Establish Организовать
		"ALPHA", // Release Выпускать
		"LIMI",  // Purify Очистить
		"SILIO"} // Discharge Освобождать
	return arr[int(int(p))]
}

func (p Plazma) Motto() string {
	arr := []string{
		"",
		"Target",     // Target Цель
		"Flow",       // Flow Поток
		"Pacify",     // Pacify Усмирять
		"Estabilish", // Establish Организовать
		"Release",    // Release Выпускать
		"Purity",     // Purify Очистить
		"Discharge"}  // Discharge Освобождать
	return arr[int(int(p))]
}

func (p Plazma) MottoRus() string {
	arr := []string{
		"",
		"Нацеливает",     // Target Цель
		"Струит",         // Flow Поток
		"Умиротворяет",   // Pacify Усмирять
		"Устанавлиивает", // Establish Организовать
		"Высвобождает",   // Release Выпускать
		"Очищает",        // Purify Очистить
		"Разряжает"}      // Discharge Освобождать
	return arr[int(int(p))]
}
