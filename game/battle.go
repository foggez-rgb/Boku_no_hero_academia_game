package game

func Battle(player Player, enemy Player) string {
	if player.Power > enemy.Power {
		return player.Name + " wins!"
	}
	return enemy.Name + " wins!"
}