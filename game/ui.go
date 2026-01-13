package game

import (
	"fmt"
	"os"
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