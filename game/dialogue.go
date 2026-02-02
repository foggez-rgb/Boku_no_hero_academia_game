package game

import (
	"math/rand"
	"time"

	"Boku_no_hero_academia_game/game/characters"
)

func Talk(c characters.Character, topic string) string {
	rand.Seed(time.Now().UnixNano())

	if answers, ok := c.Responses[topic]; ok {
		return answers[rand.Intn(len(answers))]
	}

	return "Бакуго смотрит на тебя с раздражением..."
}