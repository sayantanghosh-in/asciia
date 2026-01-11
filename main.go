package main

import (
	"fmt"

	game "github.com/sayantanghosh-in/asciia/game"
)

func main() {
	_, cols := game.GetTerminalDimensions()
	for {
        // 1. Clear and Render

		fmt.Print("\033[H\033[2J")
		fmt.Print(game.DrawHeader(cols))

        totalLinesWithoutHeaderAndFooter := game.GetTotalLinesWithoutHeaderAndFooter()

        // 2. Draw the Game
        game.DrawGame()
        game.PrintEmptyLines(totalLinesWithoutHeaderAndFooter-7)

		// 3. Wait for Input
        input := game.PrintFooter()

        // 4. Exit Logic
        if input == "q" {
            fmt.Println("Exiting Asciia... Goodbye!")
            break 
        }
	}
}