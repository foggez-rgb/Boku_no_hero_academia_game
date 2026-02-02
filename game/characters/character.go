package characters

type Emotions struct {
	Anger       int
	Respect     int
	Affection   int
}

type Memory struct {
	LastGift string
	TimesTalked int
}

type Character interface {
	GetName() string
	React(text string) string
	ReceiveGift(gift Gift)
	PrintEmotions()
}