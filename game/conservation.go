package game

import (
	"game/ai"
	"game/characters"
)

func GenerateCharacterReply(
	character *characters.Character,
	playerText string,
	baseReply string,
) string {

	// если романтики мало — обычный текст
	if character.Romance.Level < 60 {
		return baseReply
	}

	// собираем prompt
	prompt := ai.BuildPrompt(ai.PromptData{
		Style:      character.Style,
		Romance:    character.Romance.Level,
		PlayerText: playerText,
		BaseReply:  baseReply,
	})

	// генерируем ответ
	return ai.Generate(prompt)
}