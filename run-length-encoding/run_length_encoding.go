package encode

import (
	"fmt"
	"strings"
)

func RunLengthEncode(input string) string {
	var counter int
	var currentRune rune
	var encodedString []string
	for i, x := range input {
		if i == 0 {
			currentRune = x
			counter++
			continue
		}

		if x == currentRune && i != len(input)-1 {
			counter++
			continue
		}

		var singleUnit string

		// if i == len(input)-1 {
		// 	singleUnit = fmt.Sprintf("%d", counter) + string(currentRune)
		// 	encodedString = append(encodedString, singleUnit)

		// }

		if counter == 1 {
			singleUnit = string(currentRune)
		} else {
			singleUnit = fmt.Sprintf("%d", counter) + string(currentRune)
		}

		fmt.Println(singleUnit, "SU")
		encodedString = append(encodedString, singleUnit)
		fmt.Println(encodedString)
		currentRune = x
	}

	return strings.Join(encodedString, "")
}

func RunLengthDecode(input string) string {
	panic("Please implement the RunLengthDecode function")
}
