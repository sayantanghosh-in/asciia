package game

type CharacterState int

const (
	idle CharacterState = iota
	attack
	hit
	heal
	defeated
)