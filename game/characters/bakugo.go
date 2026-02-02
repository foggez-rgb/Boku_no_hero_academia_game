package characters

import "strings"

type Bakugo struct {
	mood       Mood
	annoyance  int
	bond       int  // уровень привязанности
	conversations int
}

func NewBakugo() *Bakugo {
	return &Bakugo{
		mood: Neutral,
	}
}

func (b *Bakugo) GetName() string {
	return "Бакуго Кацуки"
}

func (b *Bakugo) GetMood() Mood {
	return b.mood
}

func (b *Bakugo) Talk(input string) string {
	b.conversations++
	input = strings.ToLower(input)

	// 💥 ОСКОРБЛЕНИЯ
	if strings.Contains(input, "дурак") || strings.Contains(input, "тупой") {
		b.annoyance++
		b.bond -= 2
		b.mood = Angry
		return "ЧЁ ТЫ СКАЗАЛ?! ХОЧЕШЬ, ЧТОБ Я ТЕБЯ ВЗОРВАЛ?!"
	}

	// 💬 ПОХВАЛА
	if strings.Contains(input, "крутой") || strings.Contains(input, "сильный") {
		b.bond += 2

		if b.bond >= 6 {
			b.mood = Flustered
			return "Т-ТЫ ЧЁ НЕСЁШЬ, ИДИОТ?! Я НЕ ПРОСИЛ ТАКОГО!"
		}

		b.mood = Annoyed
		return "Тц… Хватит нести хрень."
	}

	// ❤️ МЯГКИЕ ФРАЗЫ
	if strings.Contains(input, "мне нравится") || strings.Contains(input, "я люблю тебя") {
		b.bond += 3
		b.mood = Flustered
		return "Ч-ЧЁ?! СОВСЕМ С УМА СОШЁЛ?! НЕ ВЗДУМАЙ ПОВТОРЯТЬ!"
	}

	// 😡 ПЕРЕГРЕВ
	if b.annoyance >= 3 {
		b.mood = Angry
		return "ХВАТИТ МНЕ ТУТ МОРГИ ЗАСОРЯТЬ! СВАЛИ!"
	}

	// 💖 СКРЫТАЯ ПРИВЯЗАННОСТЬ
	if b.bond >= 8 {
		b.mood = Flustered
		return "…Чё ты всё ещё тут? Если уйдёшь — я не… не обрадуюсь."
	}

	// 🗯 ОБЫЧНО
	return "Чё уставился? Говори уже."
}