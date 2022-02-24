package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {

	var word []string

	wordCount := Frequency{
		"": 0,
	}
	var wordString string
	for _, i := range phrase {

		wordString = strings.ToLower(strings.Join(word, ""))

		if string(i) != "," && string(i) != "'" && unicode.IsPunct(i) || unicode.IsSymbol(i) {
			continue
		}

		if unicode.IsSpace(i) || string(i) == "," {

			if val, ok := wordCount[wordString]; ok {
				wordCount[wordString] = val + 1
				wordString = ""
				word = nil
				continue
			} else {
				wordCount[wordString] = 1
				wordString = ""
				word = nil
				continue
			}

		} else {

			word = append(word, string(i))

		}
	}

	wordString = strings.ToLower(strings.Join(word, ""))

	if val, ok := wordCount[wordString]; ok {
		wordCount[wordString] = val + 1
	} else {
		wordCount[wordString] = 1
	}

	delete(wordCount, "")

	if len(phrase) == len(word) {
		wordCount[phrase] = 1
		return wordCount
	}

	for x := range wordCount {
		var quoteRemoval []string
		for _, j := range x {
			quoteRemoval = append(quoteRemoval, string(j))
		}
		if quoteRemoval[0] == "'" && quoteRemoval[len(x)-1] == "'" {
			quoteRemoval = quoteRemoval[1:(len(quoteRemoval) - 1)]
			quoteRemovedString := strings.ToLower(strings.Join(quoteRemoval, ""))
			delete(wordCount, x)
			wordCount[quoteRemovedString] = wordCount[quoteRemovedString] + 1
		}
	}

	return wordCount
}
