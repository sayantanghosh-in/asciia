package game

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
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

// readInput waits for a single keystroke and returns it immediately.
func readInput() string {
    // 1. Get File Descriptor
    fd := int(os.Stdin.Fd())

    // 2. Switch to Raw Mode (Disables Enter requirement, line editing, etc.)
    oldState, err := term.MakeRaw(fd)
    if err != nil {
        return ""
    }
    
    // 3. IMPORTANT: Ensure we restore the terminal when this function ends
    // If we don't do this, the user's terminal will look broken after the game.
    defer term.Restore(fd, oldState)

    // 4. Read exactly one byte
    b := make([]byte, 1)
    os.Stdin.Read(b)

    return string(b)
}

func printDefaultFooter() string {
	fmt.Print("Press 'q' to quit: ")
        var input string
       	input = readInput()

		return input
}

func printInitFooter() string {
	fmt.Print("Press 's' to start the game or 'q' to quit: ")
        var input string
       	input = readInput()

		return input
}

func DrawFooter(gameData GameData) string {
	var userInput string
	switch gameData.State {
		case Init: {
			userInput = printInitFooter()
			break
		}
		case InGame, Over: {
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