package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, i := range alphabet {
		if strings.Contains(strings.ToLower(input), i) {
			continue
		} else {
			return false
		}
	}
	return true
}
