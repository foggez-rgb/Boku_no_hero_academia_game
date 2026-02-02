package characters

import "yourmodule/game"

type Character interface {
	GetName() string
	React(input string, state *game.GameState) string
	GetRomance() int
	AddRomance(value int)
}