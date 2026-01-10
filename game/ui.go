package game

import "strings"

func DrawHeader(cols int) string {
	var sb strings.Builder

	infoLine1 := "\nVersion: 0.0.1"
	infoLine2 := "\nSayantan Ghosh (https://sayantanghosh.in)"

	sb.WriteString(asciiaTitleText)
	sb.WriteString(infoLine1)
	sb.WriteString(infoLine2)

	return sb.String()
}