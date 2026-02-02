package game

type RomanceStage int

const (
	Hostile RomanceStage = iota
	Annoyed
	Neutral
	Attached
	InLove
)

func GetRomanceStage(value int) RomanceStage {
	switch {
	case value < 20:
		return Hostile
	case value < 40:
		return Annoyed
	case value < 60:
		return Neutral
	case value < 80:
		return Attached
	default:
		return InLove
	}
}