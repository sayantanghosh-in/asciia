package game

import "fmt"

func attack(p1 Character, p2 Character) {
	/**
	* Function to simulate p1 attacking p2
	* A random HP between 10-20 will be deducted.
	* This function will mutate the characters in the state
	*/
	fmt.Println(p1.Name + " attacked " + p2.Name)
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
			attack(gameData.Player, gameData.Monster)
		}
		case "h": {
			// heal
		}
	}
}