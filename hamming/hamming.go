package hamming

import (
	"errors"
)

var ErrDiffStringLength = errors.New("strings are a differnt length")

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, ErrDiffStringLength
	}
	var diffTally int
	x := 0
	for _, i := range a {
		if i != rune(b[x]) {
			diffTally += 1
		}
		x++
	}
	return diffTally, nil
}
