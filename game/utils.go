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

func PrintEmptyLines(numberOfLines int) {
	for range numberOfLines  {
		fmt.Println("")
	}
}

func printDefaultFooter() string {
	fmt.Print("Press 'q' to quit or 'Enter' to refresh: ")
        var input string
        fmt.Scanln(&input)

		return input
}

func PrintFooter() string {
	// @TODO implement full logic for dynamic footer
	input := printDefaultFooter()
	return input
}