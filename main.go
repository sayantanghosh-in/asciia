package main

import (
	"fmt"

	game "github.com/sayantanghosh-in/asciia/game"
)

func initialize() game.GameData {
    return game.GameData{
        Player: game.Character{
            Name: "Player 1",
            MaxHP: 100,
            CurrentHP: 100,
            State: game.Idle,
        },
        Monster: game.Character{
            Name: "Monster",
            MaxHP: 100,
            CurrentHP: 100,
            State: game.Idle,
        },
        Turn: 1,
        State: game.Init,
    }
}

func main() {
    // initialize the game state
    var gameData game.GameData = initialize()

    // calculate important things for rendering in the terminal
	_, cols := game.GetTerminalDimensions()

	for {
        // 1. Clear and Render

		fmt.Print("\033[H\033[2J")
		fmt.Print(game.DrawHeader(cols))

        // 2. Draw the Game
        game.DrawGame(gameData)

		// 3. Wait for Input
        input := game.DrawFooter(gameData)

        // 4. Exit Logic
        if input == "q" {
            fmt.Println("Exiting Asciia... Goodbye!")
            break 
        }
	}
}