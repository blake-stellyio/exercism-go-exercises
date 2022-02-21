package luhn

import (
	"strconv"
	"unicode"
)

func Valid(id string) bool {

	var digits []int

	for _, x := range id {
		if unicode.IsDigit(x) {
			y, _ := strconv.Atoi(string(x))
			digits = append(digits, y)
		} else if unicode.IsSpace(x) {
			continue
		} else if unicode.IsLetter(x) || unicode.IsPunct(x) || unicode.IsGraphic(x) {
			return false
		}
	}

	t := len(digits)

	if t <= 1 {
		return false
	}

	var newDigits []int

	for i := range digits {
		digit := digits[len(digits)-(i+1)]
		if i%2 == 1 {
			var doubledDigit int
			if digit*2 > 9 {
				doubledDigit = (digit * 2) - 9
			} else {
				doubledDigit = digit * 2
			}
			newDigits = append(newDigits, doubledDigit)
		} else {
			newDigits = append(newDigits, digit)
		}

	}

	var sum int

	for _, x := range newDigits {
		sum = sum + x
	}

	var ifValid bool

	if sum%10 == 0 {
		ifValid = true
	} else if sum%10 != 0 {
		ifValid = false
	}
	return ifValid
}
