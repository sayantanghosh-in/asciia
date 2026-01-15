package game

import "fmt"

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

func HandleInGameKeys(key string) {
	switch key {
        case "a": {
        	// attack
		}
		case "h": {
			// heal
		}
	}
}