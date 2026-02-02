package main

import (
	"bufio"
	"fmt"
	"os"

	"bnhagame/game"
	"bnhagame/game/characters"
)

func main() {
	player := game.NewPlayer("Игрок")

	bakugo := characters.NewBakugo()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Меню ---")
		fmt.Println("1. Поговорить с Бакуго")
		fmt.Println("2. Подарить подарок Бакуго")
		fmt.Println("3. Показать эмоции Бакуго")
		fmt.Println("0. Выход")

		fmt.Print("> ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			game.StartDialogue(player, bakugo)
		case 2:
			characters.GiveGiftMenu(bakugo)
		case 3:
			bakugo.PrintEmotions()
		case 0:
			return
		}
		reader.ReadString('\n')
	}
}