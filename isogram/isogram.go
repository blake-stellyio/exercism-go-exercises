package isogram

import "strings"

func IsIsogram(word string) bool {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var isogramExists bool
	if word == "" {
		isogramExists = true
	}
	for _, x := range alphabet {
		isogram := strings.Count(word, x)
		if isogram > 1 {
			isogramExists = false
		} else if isogram == 1 {
			isogramExists = true
		}
	}
	return isogramExists
}
