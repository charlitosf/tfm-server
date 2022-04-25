package stringutilities

import "strings"

// Split a string by the dot separator, reverse its elements and join them with the dot separator
func ReverseSplitJoin(s string) string {
	// Split string by dot separator
	splitted := strings.Split(s, ".")
	// Reverse elements
	for i, j := 0, len(splitted)-1; i < j; i, j = i+1, j-1 {
		splitted[i], splitted[j] = splitted[j], splitted[i]
	}
	// Join elements with dot separator
	return strings.Join(splitted, ".")
}
