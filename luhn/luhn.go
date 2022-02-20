package luhn

import (
	"unicode"
)

func Valid(id string) bool {

	var digits []int

	for _, x := range id {
		if unicode.IsDigit(x) {
			y := int(x)
			digits = append(digits, y)
		}
	}

	t := len(digits)

	if t <= 1 {
		return false
	}

	var secondDigit []int
	z := 0
	for {
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
		doubledSecondDigit = append(doubledSecondDigit, x*2)
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
