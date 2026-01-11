package game

import (
	"fmt"
	"strconv"
	"strings"
)

func drawHealthBars(player1Health int, player2Health int) {
	var healthBars strings.Builder
	healthBars.WriteString("  ")
	healthBars.WriteString("❤️")
	healthBars.WriteString(strconv.Itoa(player1Health))
	healthBars.WriteString(characterGap)
	healthBars.WriteString("❤️")
	healthBars.WriteString(strconv.Itoa(player2Health))
	fmt.Print(healthBars.String())
}

func drawCharacters() {
	PrintEmptyLines(2)
	var hero = heroStates[idle]
	var monster = monsterStates[defeated]
	var characters strings.Builder

	var heroLines = strings.Split(hero, "\n")
	var monsterLines = strings.Split(monster, "\n")
	for i := 0; i < len(heroLines); i++ {
		characters.WriteString(heroLines[i])
		const numberOfSpaces = len(characterGap) - characterWidth/2
		for range numberOfSpaces {
			characters.WriteString(" ")
		}
		characters.WriteString(monsterLines[i])
		characters.WriteRune('\n')
	}
	fmt.Print(characters.String())
}

func DrawGame() {
	drawCharacters()
	drawHealthBars(100, 80)
}