package game

const asciiaTitleText string = `
   __    ___   ___  ____  ____    __   
  /__\  / __) / __)(_  _)(_  _)  /__\  
 /(__)\ \__ \( (__  _)(_  _)(_  /(__)\ 
(__)(__)(___/ \___)(____)(____)(__)(__)
`

const characterGap = "          "
const characterWidth = 8

/*
 * map to figure out easily how many lines the game will atleast take per game state.
 * more lines for logs and other things will be added to this
 * this is necessary to calculate the number of empty lines to render after the game content
 * and before the footer
 */
var fixedLinesPerGameState map[GameState]int = map[GameState]int{
	Init:   5, // @TODO - we will figure this out when we implement the idle state
	InGame: 7,
	Over:   7, // Same screen as InGame will be shown
}

const heroIdle = `
   [o]  
   /|\  
   / \  
`
const heroAttack = `
   [>]  
   /|\⚡
   / \   
`

const heroHit = `
   [-]💢
   /|\  
   / \  
`

const heroHeal = `
   [+]  
   /|\❤️
   / \  
`

const heroDefeated = `  
   _ _  
  [x_x] 
  /   \ 
`

const monsterIdle = `
  /M\   
  (oo)  
  /MM\    
`

const monsterAttack = `
   /M\  
💥(<<)   
  /MM\    
`

const monsterHit = `
 💢/M\   
   (--) 
   /MM\    
`

const monsterDefeated = `
        
   _xx_ 
  /____\
`

var heroStates = map[CharacterState]string{
	Idle:     heroIdle,
	Attack:   heroAttack,
	Hit:      heroHit,
	Heal:     heroHeal,
	Defeated: heroDefeated,
}

var monsterStates = map[CharacterState]string{
	Idle:     monsterIdle,
	Attack:   monsterAttack,
	Hit:      monsterHit,
	Defeated: monsterDefeated,
}