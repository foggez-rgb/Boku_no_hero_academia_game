package characters

import (
	"fmt"
	"strings"

	"bnhagame/game"
)

type Bakugo struct {
	Name     string
	Emotions Emotions
	Memory   Memory
	Romance  game.RomanceStage
}

func NewBakugo() *Bakugo {
	return &Bakugo{
		Name: "Бакуго Кацуки",
		Emotions: Emotions{
			Anger:     60,
			Respect:  10,
			Affection: 0,
		},
		Romance: game.Neutral,
	}
}

func (b *Bakugo) GetName() string {
	return b.Name
}

func (b *Bakugo) updateRomance() {
	a := b.Emotions.Affection
	r := b.Emotions.Respect

	switch {
	case a >= 60 && r >= 40:
		b.Romance = game.Confession
	case a >= 50:
		b.Romance = game.Love
	case a >= 35:
		b.Romance = game.Attachment
	case r >= 30:
		b.Romance = game.Respect
	case a >= 15:
		b.Romance = game.UsedToYou
	case a >= 5:
		b.Romance = game.Interest
	default:
		b.Romance = game.Neutral
	}
}

func (b *Bakugo) React(text string) string {
	b.Memory.TimesTalked++
	b.updateRomance()

	lower := strings.ToLower(text)

	if b.Romance == game.Confession {
		return b.confessionDialogue()
	}

	if strings.Contains(lower, "люблю") {
		b.Emotions.Anger += 5
		b.Emotions.Affection += 2
		return "Тц... Чё за херню несёшь..."
	}

	switch b.Romance {

	case game.Neutral:
		return "Не беси меня. Отвали."

	case game.Interest:
		return "Ты чё опять тут?.."

	case game.UsedToYou:
		return "Если пришёл — говори уже."

	case game.Respect:
		return "Хм. Ты не такой бесполезный, как я думал."

	case game.Attachment:
		return "Не думай, что я тебя жду... Просто так вышло."

	case game.Love:
		return "Чёрт... Почему ты всегда лезешь в голову, а?!"

	default:
		return "..."
	}
}

func (b *Bakugo) confessionDialogue() string {
	if b.Memory.LastGift != "" {
		return fmt.Sprintf(
			"Слушай... Я не умею это дерьмо говорить. "+
				"Но ты важнее всех. И если ты свалишь — я взорвусь. "+
				"Понял? Я... люблю тебя, идиот.",
		)
	}

	return "Тц... Подойди ближе. Я не повторю."
}

func (b *Bakugo) ReceiveGift(g Gift) {
	b.Memory.LastGift = g.Name
	b.Emotions.Affection += g.Affection
	b.Emotions.Respect += g.Respect
	b.Emotions.Anger += g.Anger
	b.updateRomance()
}

func (b *Bakugo) PrintEmotions() {
	fmt.Println("\nЭмоции Бакуго:")
	fmt.Println("Злость:", b.Emotions.Anger)
	fmt.Println("Уважение:", b.Emotions.Respect)
	fmt.Println("Привязанность:", b.Emotions.Affection)
	fmt.Println("Стадия отношений:", b.Romance.String())
}