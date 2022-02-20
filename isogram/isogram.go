package isogram

import "unicode"

func IsIsogram(word string) bool {

	if word == "" {
		return true
	}
	m := make(map[rune]bool)
	for _, x := range word {
		x = unicode.ToLower(x)
		if unicode.IsLetter(x) {
			if m[x] {
				return false
			} else if !m[x] {
				m[x] = true
			} else {
				continue
			}
		}
	}
	return true
}
