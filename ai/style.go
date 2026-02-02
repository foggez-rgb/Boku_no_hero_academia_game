package ai

import "fmt"

func Stylize(
	character string,
	baseText string,
	styleDescription string,
) string {

	prompt := fmt.Sprintf(
		"[Персонаж: %s]\n[Стиль: %s]\nТекст: %s",
		character,
		styleDescription,
		baseText,
	)

	return GenerateReply(prompt)
}