package characters

import (
	"fmt"
	"yourmodule/game"
)

type Bakugo struct {
	Romance int
}

func NewBakugo() *Bakugo {
	return &Bakugo{
		Romance: 10,
	}
}

func (b *Bakugo) GetName() string {
	return "Бакуго Кацуки"
}

func (b *Bakugo) GetRomance() int {
	return b.Romance
}

func (b *Bakugo) AddRomance(value int) {
	b.Romance += value
	if b.Romance < 0 {
		b.Romance = 0
	}
	if b.Romance > 100 {
		b.Romance = 100
	}
}

func (b *Bakugo) React(input string, state *game.GameState) string {
	stage := game.GetRomanceStage(b.Romance)

	switch stage {
	case game.Hostile:
		return "Чё ты несёшь, придурок?!"
	case game.Annoyed:
		return "Тц… хватит меня доставать."
	case game.Neutral:
		return "Хм. Ну… нормально."
	case game.Attached:
		return "Эй… ты чё так смотришь?"
	case game.InLove:
		return "Заткнись… Я, блядь, волнуюсь."
	default:
		return fmt.Sprintf("…%s", input)
	}
}