package main

import (
	"fmt"

	game "github.com/sayantanghosh-in/asciia/game"
)

func main() {
	for {
        // 2. Clear and Render
        fmt.Print("\033[H\033[2J")
        // @TODO - renderUI(rows, cols)
		rows, cols := game.GetTerminalDimensions()

		fmt.Print("\033[H\033[2J")
		fmt.Print(game.DrawHeader(cols))
		game.PrintEmptyLines(rows - 7)

		// 3. Wait for Input
        fmt.Print("Press 'q' to quit or 'Enter' to refresh: ")
        var input string
        fmt.Scanln(&input) // Simpler than Scanf for basic blocking

        // 4. Exit Logic
        if input == "q" {
            fmt.Println("Exiting Asciia... Goodbye!")
            break 
        }
	}
}