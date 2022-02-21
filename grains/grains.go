package grains

import (
	"errors"
	"math"
)

var ErrOutOfRange = errors.New("number is out of range")

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return uint64(0), ErrOutOfRange
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	for x := 1; x <= number; x++ {
		// unsure of how to implement this function when it takes no arguments.
	}
}
