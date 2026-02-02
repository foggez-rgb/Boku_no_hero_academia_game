package characters

import "strings"

/* ===== СОСТОЯНИЯ ===== */

type EmotionState struct {
	Anger        int
	Affection    int
	Trust        int
	Jealousy     int
	LoveUnlocked bool
}

type Bakugo struct {
	Name    string
	Emotion EmotionState
	Memory  map[string]string
}

/* ===== СОЗДАНИЕ ===== */

func NewBakugo() *Bakugo {
	return &Bakugo{
		Name: "Бакуго Кацуки",
		Emotion: EmotionState{
			Anger:     60,
			Affection: 10,
			Trust:     15,
		},
		Memory: make(map[string]string),
	}
}

/* ===== INTERFACE ===== */

func (b *Bakugo) GetName() string {
	return b.Name
}

/* ===== ПОДАРКИ ===== */

func (b *Bakugo) GiveGift(gift Gift) string {
	b.Emotion.Affection += gift.Affection
	b.Emotion.Trust += gift.Trust
	b.Emotion.Anger += gift.Anger

	if gift.ID == "letter" {
		b.Memory["letter"] = "Ты написал ему письмо"
		return "Ч-ЧТО ЭТО?! …Я не выброшу. Даже не надейся."
	}

	if gift.ID == "spicy_food" {
		return "Хм… норм. Ты не полный идиот."
	}

	return "Тц. Ладно."
}

/* ===== ДИАЛОГ ===== */

func (b *Bakugo) React(input string) string {
	text := strings.ToLower(input)

	if strings.Contains(text, "люблю") {
		b.Emotion.Affection += 15
		b.Emotion.Trust += 10
		return "ТЫ ЧТО НЕСЁШЬ?! …Идиот… но продолжай."
	}

	if strings.Contains(text, "другой") {
		b.Emotion.Jealousy += 15
		return "А МНЕ КАКОЕ ДЕЛО?! …Только попробуй."
	}

	if strings.Contains(text, "дурак") {
		b.Emotion.Anger += 20
		return "Я ТЕБЯ ВЗОРВУ!"
	}

	return "Говори нормально."
}

/* ===== РОМАНС ===== */

func (b *Bakugo) CheckRomance() string {
	if b.Emotion.Affection >= 80 && b.Emotion.Trust >= 70 {
		if !b.Emotion.LoveUnlocked {
			b.Emotion.LoveUnlocked = true
			return b.confess()
		}
	}
	return ""
}

func (b *Bakugo) confess() string {
	return `
Чёрт…
Я не умею в слова.

Но если ты уйдёшь —
мне будет хреново.

Так что… оставайся.`
}