package isogram

import (
	"unicode"
	//  "strings"
)

func IsIsogram(word string) bool {

	//	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var isogramExists bool
	if word == "" {
		isogramExists = true
	}
	m := make(map[rune]bool)
	isogramExists = true
	for _, x := range word {
		x = unicode.ToLower(x)
		if unicode.IsLetter(x) {
			if m[x] {
				isogramExists = false
			} else if !m[x] {
				m[x] = true
			} else {
				continue
			}
		}
	}
	return isogramExists
}

/*
	for _, x := range alphabet {
		isogram := strings.Count(strings.ToLower(word), x)
		if isogram > 1 {
			isogramExists = false
			break
		} else if isogram == 1 {
			isogramExists = true
		}
	}
	return isogramExists
}
*/
