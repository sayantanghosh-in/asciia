package game

type CharacterState int
type GameState int

const (
	Idle CharacterState = iota
	Attack
	Hit
	Heal
	Defeated
)

const (
	Init GameState = iota
	InGame
	Over
)

// Character represents either the Hero or the Monster
type Character struct {
	Name      string
	MaxHP     int
	CurrentHP int
	State     CharacterState
}

// GameData holds the entire "snapshot" of the current game
type GameData struct {
	Player   Character
	Monster  Character
	Turn     int // 1 means Hero, 2 means player 2 (monster)
	lastMove string
	State    GameState
}