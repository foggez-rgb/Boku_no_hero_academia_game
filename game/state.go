package game

import "yourmodule/game/characters"

type GameState struct {
	Player            *Player
	CurrentCharacter  characters.Character
	Flags             map[string]bool
	Variables         map[string]int
}

func NewGameState(player *Player) *GameState {
	return &GameState{
		Player:    player,
		Flags:    make(map[string]bool),
		Variables: make(map[string]int),
	}
}