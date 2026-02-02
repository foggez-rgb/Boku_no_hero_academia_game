package characters

import "strings"

type Bakugo struct {
	mood      Mood
	annoyance int
}

func NewBakugo() *Bakugo {
	return &Bakugo{
		mood:      Neutral,
		annoyance: 0,
	}
}

func (b *Bakugo) GetName() string {
	return "Бакуго Кацуки"
}

func (b *Bakugo) GetMood() Mood {
	return b.mood
}

func (b *Bakugo) Talk(input string) string {
	input = strings.ToLower(input)

	// Если его бесят
	if strings.Contains(input, "дурак") || strings.Contains(input, "тупой") {
		b.annoyance++
		b.mood = Angry
		return "ЧЁ ТЫ СКАЗАЛ, УБЛЮДОК?! Я ТЕБЯ ВЗОРВУ!"
	}

	// Если его хвалят
	if strings.Contains(input, "крутой") || strings.Contains(input, "сильный") {
		b.mood = Annoyed
		return "Тц... Я и так это знаю. Не беси."
	}

	// Если его раздражение выросло
	if b.annoyance >= 3 {
		b.mood = Angry
		return "ХВАТИТ МНЕ ТУТ ТРЫНДЕТЬ! ИСЧЕЗНИ!"
	}

	// Обычная реакция
	return "Чё уставился? Говори нормально."
}