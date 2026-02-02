package characters

type Mood int

const (
	Neutral Mood = iota
	Angry
	Annoyed
	Friendly
	Flustered // смущён / влюблён
)

type Character interface {
	GetName() string
	GetMood() Mood
	Talk(input string) string
}