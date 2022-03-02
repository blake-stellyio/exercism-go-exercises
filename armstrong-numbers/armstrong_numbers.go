package armstrong

import (
	"fmt"
	"math"
)

func IsNumber(n int) bool {

	nString := fmt.Sprintln(n)
	var digits []float64

	for _, x := range nString {
		if float64(x)-'0' >= 0 {
			digits = append(digits, float64(x)-'0')
		}
	}

	var sum int

	for _, x := range digits {
		sum += int(math.Pow(x, float64(len(digits))))
	}

	if sum == n {
		return true
	}

	return false
}
