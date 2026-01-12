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

func drawCharacters(gameData GameData) {
	PrintEmptyLines(2)
	var hero = heroStates[gameData.Player.State]
	var monster = monsterStates[gameData.Monster.State]
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

func DrawGame(gameData GameData) {
	totalLinesWithoutHeaderAndFooter := GetTotalLinesWithoutHeaderAndFooter()
	
	// render to the screen
	
	switch gameData.State {
		case Init: {
			// TODO
			fmt.Print("\nGame INIT")
			PrintEmptyLines(4)
			break;
		}
		case InGame:
		case Over: {
			drawCharacters(gameData)
				drawHealthBars(gameData.Player.CurrentHP, gameData.Monster.CurrentHP)
				break
		}
		default: {
			fmt.Print("Invalid Game State: Please raise in an issue in Github project: https://github.com/sayantanghosh-in/asciia")
		}
	}
    
	// render empty lines after the game screen content is loaded
	PrintEmptyLines(totalLinesWithoutHeaderAndFooter-fixedLinesPerGameState[gameData.State])
}