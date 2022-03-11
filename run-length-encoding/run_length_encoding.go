package encode

import (
	"strconv"
	"unicode"
)

func RunLengthEncode(input string) string {

	if input == "" {
		return ""
	}

	var output []rune
	var counter int
	var currentLetter rune
	var lastLetter rune

	for i, x := range input {

		if i == 0 {
			currentLetter = x
			counter = 1
		}

		if currentLetter == x && i != 0 {
			counter++
		}

		if x != currentLetter {
			if counter == 1 {
				output = append(output, currentLetter)
			} else {
				output = AppendCounterAndLetter(counter, currentLetter, output)
			}
			counter = 1
			currentLetter = x
			continue
		}

		if i == len(input)-1 {
			lastLetter = x
		}
	}

	if currentLetter != lastLetter {
		if counter == 1 {
			output = append(output, currentLetter)
		} else {
			output = AppendCounterAndLetter(counter, currentLetter, output)
		}
	} else {
		output = AppendCounterAndLetter(counter, lastLetter, output)
	}
	return string(output)
}

func RunLengthDecode(input string) string {

	if input == "" {
		return ""
	}

	var previousIsNumber bool
	var previousNumber int
	var output []rune

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
				output = append(output, x)
			}
			previousIsNumber = false
			continue
		}

		if !previousIsNumber && (unicode.IsLetter(x) || unicode.IsSpace(x)) {
			output = append(output, x)
		}

	}

	return string(output)
}

func AppendCounterAndLetter(counter int, letter rune, output []rune) []rune {
	if counter >= 10 {
		number := strconv.Itoa(counter)
		var numberSlice []rune
		for _, x := range number {
			numberSlice = append(numberSlice, x)
		}
		output = append(output, numberSlice...)
	} else {
		output = append(output, rune(counter+'0'))
	}

	output = append(output, letter)

	return output
}
