package game

import (
	"fmt"
	"strconv"
	"strings"
)

func DrawHeader(cols int) string {
	var sb strings.Builder

	infoLine1 := "\nVersion: 0.0.1"
	infoLine2 := "\nSayantan Ghosh (https://sayantanghosh.in)"

	sb.WriteString(asciiaTitleText)
	sb.WriteString(infoLine1)
	sb.WriteString(infoLine2)

	return sb.String()
}

func PrintEmptyLines(numberOfLines int) {
	for range numberOfLines  {
		fmt.Println("")
	}
}

func printDefaultFooter() string {
	fmt.Print("Press 's' to start a new game or 'q' to quit: ")
        var input string
       	input = readInput()

		return input
}

func printInGameFooter(gameData *GameData) string {
	switch gameData.Turn {
		case 1: {
			fmt.Print("Press 'a' to attack; 'h' to heal; 'r' to restart; 'q' to quit")
		}
		case 2: {
			// TODO - this will be replaced with an automatic Monster move
			fmt.Print("Press 'a' to attack; 'r' to restart; 'q' to quit")
		}
	}
    var input string
    input = readInput()

	return input
}

func DrawFooter(gameData *GameData) string {
	var userInput string
	switch gameData.State {
		case Init, Over: {
			userInput = printDefaultFooter()
			break
		}
		case InGame: {
			userInput = printInGameFooter(gameData)
			break
		}
		default: {
			userInput = printDefaultFooter()
			break
		}
	}
	return userInput
}

func drawIntro() {
    PrintEmptyLines(2)
    fmt.Print(introScreenText)
}

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

func drawCharacters(gameData *GameData) {
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

func drawLastMove(lastMove string) {
	PrintEmptyLines(2)
	fmt.Print(lastMove)
}