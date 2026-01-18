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

func attack(gameData *GameData) {
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
			gameData.Monster.CurrentHP = max(0, reducedHp)
			// updating player states
			gameData.Player.State = Attack
			gameData.Monster.State = Hit
			// updating game turn
			gameData.Turn = 2
			// updating last move
			gameData.lastMove = "You attacked Monster. Monster: -" + strconv.Itoa(hpToReduce) + "HP"
			// checking if game is over
			if gameData.Monster.CurrentHP == 0 {
				gameData.State = Over
				gameData.Monster.State = Defeated
				gameData.lastMove = "You won!!!"
			}
			break
		} 
		case 2: {
			// Turn = 2 means player 2 turn
			// reducing the hp of the player 1
			hpToReduce := calculateAttackHpReductionQuantity()
			reducedHp := gameData.Player.CurrentHP - hpToReduce
			gameData.Player.CurrentHP = max(0, reducedHp)
			// updating player states
			gameData.Monster.State = Attack
			gameData.Player.State = Hit
			// updating game turn
			gameData.Turn = 1
			// updating last move
			gameData.lastMove = "Monster attacked You. Player: -" + strconv.Itoa(hpToReduce) + "HP"
			// checking if game is over
			if gameData.Player.CurrentHP == 0 {
				gameData.State = Over
				gameData.Player.State = Defeated
				gameData.lastMove = "You Lost :("
			}
			break
		}
	}
}

func heal(gameData *GameData) {
	/**
	* Function to simulate player 1 taking a heal
	* Currently only player 1 can heal
	* Healing is only possible 1 time
	* Healing increases HP of player 1 by a fixed amount = 15HP
	*/
	gameData.Player.CurrentHP = min(gameData.Player.CurrentHP + 15, 100)
	gameData.Turn = 2
	gameData.Player.State = Heal
	gameData.Monster.State = Idle
	gameData.lastMove = "You healed yourself: +15HP"
}

func DrawGame(gameData *GameData) {
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

func HandleInGameKeys(gameData *GameData, key string) {
	switch key {
        case "a": {
        	// attack
			attack(gameData)
		}
		case "h": {
			// heal
			heal(gameData)
		}
	}
}