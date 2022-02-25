package isbn

import (
	"strconv"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	var isbnINT []int
	for _, x := range isbn {
		if !unicode.IsDigit(x) && string(x) != "x" || string(x) != "X" {
			return false
		} else if unicode.IsPunct(x) && string(x) != "-" {
			return false
		} else if string(x) == "-" {
			continue
		} else {
			var ifT int
			// ? not sure if this is how the formula works. I think it acts as a multipler instead.
			if string(x) == "x" || string(x) == "X" {
				ifT = 10
				isbnINT = append(isbnINT, ifT)
			} else {
				i, _ := strconv.Atoi(string(x))
				isbnINT = append(isbnINT, i)
			}

		}
	}

}
