package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {

	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	if shiftKey == 0 || shiftKey == 26 {
		return plain
	}

	var plainOutput []string

	for _, x := range plain {

		if unicode.IsSpace(x) || unicode.IsNumber(x) || unicode.IsPunct(x) {
			plainOutput = append(plainOutput, string(x))
			continue
		}

		for p, l := range alphabet {
			if unicode.IsUpper(x) {
				if l == strings.ToLower(string(x)) {
					if p+shiftKey < 25 {
						plainOutput = append(plainOutput, strings.ToUpper((alphabet[p+shiftKey])))
						continue
					} else {
						plainOutput = append(plainOutput, strings.ToUpper((alphabet[p+shiftKey-26])))
						continue
					}
				}
			}
			if l == strings.ToLower(string(x)) {
				if p+shiftKey <= 25 {
					plainOutput = append(plainOutput, alphabet[p+shiftKey])
				} else {
					plainOutput = append(plainOutput, alphabet[p+shiftKey-26])
				}
			}
		}
	}
	return strings.Join(plainOutput, "")
}
