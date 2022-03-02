package isbn

import (
	"strconv"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	var isbnINT []int
	for _, x := range isbn {
		if !unicode.IsDigit(x) && string(x) != "x" && string(x) != "X" && string(x) != "-" {
			return false
		} else if string(x) == "-" {
			continue
		} else {
			var ifT int
			if string(x) == "x" || string(x) == "X" {
				ifT = 10
				isbnINT = append(isbnINT, ifT)
			} else {
				i, _ := strconv.Atoi(string(x))
				isbnINT = append(isbnINT, i)
			}
		}
	}

	if len(isbnINT) != 10 {
		return false
	}

	for i, x := range isbnINT {
		if x == 10 {
			if i != 9 {
				return false
			}
		}
	}

	output := ((isbnINT[0]*10)+(isbnINT[1]*9)+(isbnINT[2]*8)+(isbnINT[3]*7)+(isbnINT[4]*6)+(isbnINT[5]*5)+(isbnINT[6]*4)+(isbnINT[7]*3)+(isbnINT[8]*2)+isbnINT[9])%11 == 0

	return output
}
