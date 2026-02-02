package characters

import (
	"fmt"
	"strings"
)

type Bakugo struct {
	Name    string
	Emotions Emotions
	Memory   Memory
}

func NewBakugo() *Bakugo {
	return &Bakugo{
		Name: "Бакуго Кацуки",
		Emotions: Emotions{
			Anger: 60,
			Respect: 10,
			Affection: 0,
		},
	}
}

func (b *Bakugo) GetName() string {
	return b.Name
}

func (b *Bakugo) React(text string) string {
	b.Memory.TimesTalked++

	lower := strings.ToLower(text)

	if strings.Contains(lower, "привет") {
		b.Emotions.Respect += 1
		return "ЧЁ ТЫ ПРИПЁРСЯ?! ...Ладно, привет."
	}

	if strings.Contains(lower, "люблю") {
		b.Emotions.Affection += 2
		b.Emotions.Anger += 5
		return "ЧТО ТЫ НЕСЁШЬ, ДЕБИЛ?! ...Тц."
	}

	if b.Emotions.Affection > 10 {
		return "Не думай, что я стал мягче, ясно?!"
	}

	return "Хватит болтать. У меня тренировка."
}

func (b *Bakugo) ReceiveGift(g Gift) {
	b.Memory.LastGift = g.Name
	b.Emotions.Affection += g.Affection
	b.Emotions.Respect += g.Respect
	b.Emotions.Anger += g.Anger
}

func (b *Bakugo) PrintEmotions() {
	fmt.Println("\nЭмоции Бакуго:")
	fmt.Println("Злость:", b.Emotions.Anger)
	fmt.Println("Уважение:", b.Emotions.Respect)
	fmt.Println("Привязанность:", b.Emotions.Affection)
}