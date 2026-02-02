package ai

import (
	"fmt"
	"yourmodule/game"
	"yourmodule/game/characters"
)

func BuildPrompt(
	character characters.Character,
	state *game.GameState,
	playerInput string,
	baseReply string,
) string {

	stage := game.GetRomanceStage(character.GetRomance())

	return fmt.Sprintf(`
Ты — персонаж из аниме My Hero Academia.

Имя: %s
Характер: агрессивный, вспыльчивый, грубый, но честный.
Романтическая стадия: %v

Правила:
- Не будь милым
- Не извиняйся
- Говори грубо
- Если влюблён — скрывай это злостью

Фраза игрока:
"%s"

Твой базовый ответ:
"%s"

Перепиши базовый ответ, усилив стиль персонажа.
`,
		character.GetName(),
		stage,
		playerInput,
		baseReply,
	)
}