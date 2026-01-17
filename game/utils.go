package game

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func GetTerminalDimensions() (rows int, cols int) {
	// Get the file descriptor for standard output.
	fd := int(os.Stdout.Fd())

	// Check if the file descriptor is a terminal.
	if !term.IsTerminal(fd) {
		fmt.Println("Not running in a terminal. Cannot get size.")
		return -1, -1
	}

	// Get the terminal size (cols is columns, rows is rows).
	cols, rows, err := term.GetSize(fd)
	if err != nil {
		fmt.Printf("Error getting terminal size: %v\n", err)
		return -1, -1
	}

	return rows, cols
}

func GetTotalLinesWithoutHeaderAndFooter() (lines int) {
	rows, _ := GetTerminalDimensions()
	return rows - 9
}

// Note: The readInput function is purely AI generated
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