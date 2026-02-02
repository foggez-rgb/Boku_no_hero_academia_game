package game

import (
	"bufio"
	"fmt"
	"os"

	"bnhagame/game/characters"
)

func StartDialogue(player *Player, c characters.Character) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nТы:", player.Name)
	fmt.Print("Скажи что-нибудь: ")

	text, _ := reader.ReadString('\n')
	response := c.React(text)

	fmt.Println(c.GetName()+":", response)
}