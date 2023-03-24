package tools

import (
	"strings"
)

// ExtractNumber to extract a number from a string after a marker.
// returns "" if no number is found.
func ExtractNumber(parString string, parMarker string) string {
	if strings.Contains(parString, parMarker) {
		// Extract the number after parPrefix in the task.description
		startIndex := strings.Index(parString, parMarker) + len(parMarker)
		endIndex := startIndex
		for endIndex < len(parString) && parString[endIndex] >= '0' && parString[endIndex] <= '9' {
			endIndex++
		}
		if startIndex != endIndex {
			return parString[startIndex:endIndex]
		}
	}
	return ""
}
