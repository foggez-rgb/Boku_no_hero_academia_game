package game

import (
	"yourmodule/ai"
)

func Talk(state *GameState, input string) string {
	char := state.CurrentCharacter

	// 1. Получаем СКРИПТОВЫЙ ответ
	baseReply := char.React(input, state)

	// 2. Решаем: нужен ли ИИ
	if shouldUseAI(char) {
		prompt := ai.BuildPrompt(
			char,
			state,
			input,
			baseReply,
		)

		return ai.GenerateResponse(prompt)
	}

	// 3. Иначе — обычный ответ
	return baseReply
}

func shouldUseAI(char Character) bool {
	return char.GetRomance() >= 60
}