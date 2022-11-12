package acronym

import "unicode"

// Abbreviate generates an acronym from a given a string
func Abbreviate(s string) string {

	var result []rune
	for i := range s {

		if i == 0 {
			result = append(result, unicode.ToUpper(rune(s[0])))
			continue
		}

		if unicode.IsSpace(rune(s[i])) || string(s[i]) == "-" || string(s[i]) == "_" {
			j := i
			j++
			if !unicode.IsSpace(rune(s[j])) && !unicode.IsPunct(rune(s[j])) {
				result = append(result, unicode.ToUpper(rune(s[j])))
				continue
			}
		}
	}

	return string(result)
}
