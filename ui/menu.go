package ui

import (
	"fmt"
	"yourmodule/game"
)

func ShowMainMenu(state *game.GameState) {
	fmt.Println("1. Поговорить")
	fmt.Println("2. Подарить подарок")
}