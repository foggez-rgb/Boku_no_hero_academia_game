package ai

import "fmt"

type PromptData struct {
	Style       Style
	Romance     int
	PlayerText  string
	BaseReply   string
}

func BuildPrompt(data PromptData) string {
	return fmt.Sprintf(`
Ты — персонаж из аниме My Hero Academia.

Имя: %s
Описание: %s

Правила поведения:
%s

Романтический уровень: %d

Фраза игрока:
"%s"

Базовый ответ персонажа:
"%s"

Перепиши базовый ответ, усилив стиль персонажа.
`,
		data.Style.Name,
		data.Style.Description,
		formatRules(data.Style.Rules),
		data.Romance,
		data.PlayerText,
		data.BaseReply,
	)
}

func formatRules(rules []string) string {
	result := ""
	for _, r := range rules {
		result += "- " + r + "\n"
	}
	return result
}