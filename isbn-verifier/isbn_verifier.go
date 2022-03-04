package isbn

import (
	"unicode"
)

func IsValidISBN(isbn string) bool {
	var output, counter int

	for _, x := range isbn {
		if !unicode.IsDigit(x) && x != 'x' && x != 'X' && x != '-' {
			return false
		}

		if x != '-' {
			if (x == 'x' || x == 'X') && counter != 9 {
				return false
			} else if x == 'x' || x == 'X' {
				output += 10
			} else {
				i := int(x - '0')
				output += (i * (10 - counter))
			}
			counter++
		}
	}

	return output%11 == 0 && counter == 10
}
