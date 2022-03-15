package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {

	if input == "" {
		return ""
	}

	var output strings.Builder
	var counter int
	var currentLetter rune

	for i, x := range input {

		if i == 0 {
			currentLetter = x
			counter = 1
		}

		if currentLetter == x && i != 0 {
			counter++
		}

		if x != currentLetter {
			output.WriteString(AppendCounterAndLetter(counter, currentLetter))
			counter = 1
			currentLetter = x
		}

	}

	output.WriteString(AppendCounterAndLetter(counter, currentLetter))

	return output.String()

}

func RunLengthDecode(input string) string {

	if input == "" {
		return ""
	}

	var previousIsNumber bool
	var previousNumber int
	var output strings.Builder

	for _, x := range input {

		if unicode.IsDigit(x) && !previousIsNumber {
			previousIsNumber = true
			previousNumber = int(x - '0')
			continue
		}

		if unicode.IsDigit(x) && previousIsNumber {
			previousNumber = int(x-'0') + previousNumber*10
			previousIsNumber = true
			continue
		}

		if previousIsNumber && !unicode.IsDigit(x) {
			for j := 1; j <= previousNumber; j++ {
				output.WriteRune(x)
			}
			previousIsNumber = false
			continue
		}

		if !previousIsNumber && (unicode.IsLetter(x) || unicode.IsSpace(x)) {
			output.WriteRune(x)
		}

	}

	return output.String()
}

func AppendCounterAndLetter(counter int, letter rune) string {
	var output string
	if counter > 1 {
		number := strconv.Itoa(counter)
		output = number + string(letter)
	} else {
		output = string(letter)
	}
	return output
}
