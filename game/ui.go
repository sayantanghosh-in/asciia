package game

import (
	"fmt"
	"strings"
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

func printDefaultFooter() string {
	fmt.Print("Press 'q' to quit or 'Enter' to refresh: ")
        var input string
        fmt.Scanln(&input)

		return input
}

func DrawFooter() string {
	// @TODO implement full logic for dynamic footer
	input := printDefaultFooter()
	return input
}