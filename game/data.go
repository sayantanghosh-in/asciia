package game

const asciiaTitleText string = `
   __    ___   ___  ____  ____    __   
  /__\  / __) / __)(_  _)(_  _)  /__\  
 /(__)\ \__ \( (__  _)(_  _)(_  /(__)\ 
(__)(__)(___/ \___)(____)(____)(__)(__)
`

const characterGap = "          "
const characterWidth = 8

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
	idle:     heroIdle,
	attack:   heroAttack,
	hit:      heroHit,
	heal:     heroHeal,
	defeated: heroDefeated,
}

var monsterStates = map[CharacterState]string{
	idle:     monsterIdle,
	attack:   monsterAttack,
	hit:      monsterHit,
	defeated: monsterDefeated,
}