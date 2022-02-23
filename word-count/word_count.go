package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	var word []string
	wordCount := map[string]int{}
	for _, i := range phrase {
		// this will not handle apostrophes correctly in contractions.
		if unicode.IsPunct(i) {
			continue
		}
		if unicode.IsSpace(i) {
			wordString := strings.ToLower(strings.Join(word, ""))
			if val, ok := wordCount[wordString]; ok {
				wordCount[wordString] = val + 1
			} else {
				wordCount[wordString] = 1
			}

		} else {
			word = append(word, string(i))
		}

	}
	return wordCount
}
