package characters

func BakugoKatsuki() Character {
	return Character{
		Name:        "Кацуки Бакуго",
		Temperament: 9,
		SpeechStyle: "Грубо, агрессивно, орёт, часто матерится",
		Responses: map[string][]string{
			"hello": {
				"ЧЁ ТЕ НАДО?!",
				"ТЫ ЧЕГО ПРИПЁРСЯ?!",
			},
			"who_are_you": {
				"Я БУДУЩИЙ НОМЕР ОДИН, ЯСНО?!",
				"КАЦУКИ БАКУГО! ЗАПОМНИ!",
			},
			"praise": {
				"ЗАТКНИСЬ! Я И ТАК ЭТО ЗНАЮ!",
				"НЕ БЕСИ МЕНЯ!",
			},
			"fight": {
				"ХОЧЕШЬ ВЗРЫВОВ?!",
				"Я ТЕБЯ РАЗНЕСУ!",
			},
		},
	}
}