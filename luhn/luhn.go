package luhn

import (
	"strconv"
	"unicode"
)

func Valid(id string) bool {

	var digits []int

	for _, x := range id {
		if unicode.IsDigit(x) {
			// obviously the below line is pretty ugly. I was trying to get the Atoi function to work, but it needed a string rather than a rune
			y, _ := strconv.Atoi(strconv.QuoteRuneToASCII(x))
			digits = append(digits, y)
		}
	}

	t := len(digits)

	if t <= 1 {
		return false
	}

	var secondDigit []int
	z := 0
	for _, x := range digits {
		y := digits[(len(digits) - (z + 1))]

		if y >= 0 {
			if z%2 != 0 {
				secondDigit = append(secondDigit, y)
			}
			z++
		} else {
			break
		}
	}

	var doubledSecondDigit []int
	for _, x := range secondDigit {
		var product int
		if x*2 > 9 {
			product = (x * 2) - 9
		} else {
			product = x * 2
		}
		doubledSecondDigit = append(doubledSecondDigit, product)
	}

	var sum int

	for _, x := range doubledSecondDigit {
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
