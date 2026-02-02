package game

type RomanceStage int

const (
	Neutral RomanceStage = iota
	Interest
	UsedToYou
	Respect
	Attachment
	Love
	Confession
	AIMode
)

func (r RomanceStage) String() string {
	switch r {
	case Neutral:
		return "Нейтралитет"
	case Interest:
		return "Раздражённый интерес"
	case UsedToYou:
		return "Привык"
	case Respect:
		return "Уважение"
	case Attachment:
		return "Привязанность"
	case Love:
		return "Влюблён"
	case Confession:
		return "Признание"
	case AIMode:
		return "ИИ-режим"
	default:
		return "Неизвестно"
	}
}