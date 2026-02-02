package game

type Memory struct {
	Flags     map[string]bool
	Counters  map[string]int
}

func NewMemory() *Memory {
	return &Memory{
		Flags:    make(map[string]bool),
		Counters: make(map[string]int),
	}
}