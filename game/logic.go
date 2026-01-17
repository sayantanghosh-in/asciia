package game

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func calculateAttackHpReductionQuantity() int {
	min := 10
	max := 20
	return min + rand.IntN(max-min+1)
}

func attack(gameData GameData) {
	/**
	* Function to simulate player 1 attacking player 2
	* Currently player 1 means the current player and player 2 is the monster
	* A random HP between 10-20 will be deducted.
	* This function will mutate the characters in the state
	*/
	switch gameData.Turn {
		case 1: {
			// Turn = 1 means player 1 turn
			// reducing the hp of the player 2
			hpToReduce := calculateAttackHpReductionQuantity()
			reducedHp := gameData.Monster.CurrentHP - hpToReduce
			if reducedHp < 0 {
				gameData.Monster.CurrentHP = 0
			} else {
				gameData.Monster.CurrentHP = reducedHp
			}
			gameData.Turn = 2
			gameData.lastMove = "Player 1 attacked Monster. Monster: -" + strconv.Itoa(hpToReduce) + "HP"
			break
		} 
		case 2: {
			// Turn = 2 means player 2 turn
			// reducing the hp of the player 1
			hpToReduce := calculateAttackHpReductionQuantity()
			reducedHp := gameData.Player.CurrentHP - hpToReduce
			if reducedHp < 0 {
				gameData.Player.CurrentHP = 0
			} else {
				gameData.Player.CurrentHP = reducedHp
			}
			gameData.Turn = 1
			gameData.lastMove = "Monster attacked Player 1. Player: -" + strconv.Itoa(hpToReduce) + "HP"
			break
		}
	}
	DrawGame(gameData)
}

func DrawGame(gameData GameData) {
	totalLinesWithoutHeaderAndFooter := GetTotalLinesWithoutHeaderAndFooter()
	
	// render to the screen
	
	switch gameData.State {
		case Init: {
			drawIntro()
			break;
		}
		case InGame, Over: {
			drawCharacters(gameData)
			drawHealthBars(gameData.Player.CurrentHP, gameData.Monster.CurrentHP)
			drawLastMove(gameData.lastMove)
			break
		}
		default: {
			fmt.Print("Invalid Game State: Please raise an issue in Github project: https://github.com/sayantanghosh-in/asciia")
		}
	}
    
	// render empty lines after the game screen content is loaded
	PrintEmptyLines(totalLinesWithoutHeaderAndFooter-fixedLinesPerGameState[gameData.State])
}

func HandleInGameKeys(gameData GameData, key string) {
	switch key {
        case "a": {
        	// attack
			attack(gameData)
		}
		case "h": {
			// heal
		}
	}
}