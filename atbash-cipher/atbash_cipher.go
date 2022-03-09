package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {

	var output []rune
	counter := 0

	for _, i := range strings.ToLower(s) {

		if counter == 5 && !unicode.IsPunct(i) && !unicode.IsSpace(i) {
			output = append(output, ' ')
			counter = 0
		}

		if unicode.IsLetter(i) || unicode.IsDigit(i) {
			if i >= 'a' && i <= 'z' {
				output = append(output, ('a' + 'z' - i))
			} else {
				output = append(output, i)
			}
			counter++
		}
	}

	return string(output)

}
