package lib

type Ton int

const (
	MAGNETIC Ton = iota + 1
	LUNAR
	ELECTRIC
	SELF_EXISTING
	OVERTONE
	RHYTHMIC
	RESONANT
	GALACTIC
	SOLAR
	PLANETARY
	SPECTRAL
	CRYSTAL
	COSMIC
)

func GetTons() []Ton {
	return []Ton{MAGNETIC, LUNAR, ELECTRIC, SELF_EXISTING, OVERTONE, RHYTHMIC, RESONANT, GALACTIC, SOLAR, PLANETARY, SPECTRAL, CRYSTAL, COSMIC}
}

func (t *Ton) Inc() *Ton {
	*t++
	if *t > COSMIC {
		*t = MAGNETIC
	}
	return t
}

func (t Ton) Question() string {
	arr := []string{
		"DAY OUT OF TIME",
		"WHAT IS MY PURPOSE?",
		"WHAT IS MY CHALLENGE?",
		"HOW CAN I BEST SERVE?",
		"WHAT IS THE FORM MY SERVICE WILL TAKE?",
		"HOW CAN I BEST EMPOWER MYSELF?",
		"HOW CAN I EXTEND MY EQUALITY TO OTHERS?",
		"HOW CAN I ATTUNE MY SERVICE TO OTHERS?",
		"DO I LIVE WHAT I BELIEVE?",
		"HOW DO I ATTAIN MY PURPOSE?",
		"HOW DO I PERFECT WHAT I DO?",
		"HOW DO I RELEASE AND LET GO?",
		"HOW CAN I DEDICATE MYSELF TO ALL THAT LIVES?",
		"HOW CAN I EXPAND MY JOY AND LOVE?"}
	return arr[int(t)]
}

func (t Ton) String() string {
	arr := []string{
		"DAY OUT OF TIME",
		"MAGNETIC",
		"LUNAR",
		"ELECTRIC",
		"SELF EXISTING",
		"OVERTONE",
		"RHYTHMIC",
		"RESONANT",
		"GALACTIC",
		"SOLAR",
		"PLANETARY",
		"SPECTRAL",
		"CRYSTAL",
		"COSMIC"}
	return arr[int(t)]
}

func (t Ton) StringRus() string {
	arr := []string{
		"День вне времени",
		"Магнитная",
		"Лунная",
		"Электрическая",
		"Самосущная",
		"Обертонная",
		"Ритмическая",
		"Резонансная",
		"Галактическая",
		"Солнечная",
		"Планетарная",
		"Спектральная",
		"Кристаллическая",
		"Космическая"}
	return arr[int(t)]
}

func (t Ton) QuestionRus() string {
	arr := []string{
		" ",
		"В чём состоит моя цель?",
		"В чём вызов жизни для меня?",
		"В чём состоит моё служение?",
		"Какую форму примет моё служени?",
		"Где мне черпать мою силу?",
		"Каким образом я могу проявить моё равенство с другими?",
		"Как сонастроить мне себя в служение другими?",
		"Живу ли я согласно моим убежденям?",
		"Как достичь мне своей цели?",
		"Как мне совершенствовать свои действия?",
		"Как мне освободиться от привязаностей?",
		"Как мне посвятить себя всему живому?",
		"Как стать истосником радости и любви?"}
	return arr[int(t)]
}

func (t Ton) TotemRus() string {
	arr := []string{
		"День вне времени",
		"Летучая мышь",
		"Скорпион",
		"Олень",
		"Сова",
		"Павлин",
		"Ящерица",
		"Обезьяна",
		"Сокол",
		"Ягуар",
		"Собака",
		"Змея",
		"Кролиик",
		"Черепаха"}
	return arr[int(t)]
}

func (t Ton) FuncRus() string {
	arr := []string{
		"День вне времени",
		"Воссоединение цели",
		"Принятие вызова",
		"Активация служения",
		"Определение формы",
		"Наделение силой веления",
		"Установление равенства",
		"Проведение сонастройки",
		"Моделирование Целостности",
		"Осуществление намерения",
		"Совершенствование проявления",
		"Отпускание высвобождения",
		"Посвящение сотрудничеству",
		"Трансцендентность"}
	return arr[int(t)]
}

func (t Ton) MoonNrRus() string {
	arr := []string{
		"День вне времени",
		"Первая",
		"Вторая",
		"Третья",
		"Четвёртая",
		"Пятая",
		"Шестая",
		"Седьмая",
		"Восьмая",
		"Девятая",
		"Десятая",
		"Одинадцатая",
		"Двенадцатая",
		"Тринадцатая"}
	return arr[int(t)]
}

func (t Ton) MoonFunc3Rus() []string {
	arr := [][]string{
		{"", "", ""},
		{"Цель", "Объединение", "Привлечение"},
		{"Принятие вызова", "Поляризация", "Стабилизация"},
		{"Служение", "Активация", "Связь"},
		{"Форма", "Определение", "Соизмерение"},
		{"Сияние", "Наделение силой", "Управление"},
		{"Равенство", "Организация", "Уравновешивание"},
		{"Сонастройка", "Проведение", "Вдохновение"},
		{"Целостность", "Гармонизация", "Моделирование"},
		{"Намерение", "Пульсация", "Осуществление"},
		{"Проявленность", "Совершенствование", "Создание"},
		{"Освобождение", "Растворение", "Отпускание"},
		{"Сотрудничество", "Посвящение", "Универсализация"},
		{"Присутствие", "Стойкость", "Трансцендентность"},
	}
	return arr[int(t)]
}
