package beer

import (
	"errors"
	"fmt"
	"strings"
)

func Song() string {
	y, _ := Verses(99, 0)
	return y
}

func Verses(start, stop int) (string, error) {

	if stop < 0 || start > 99 {
		return "", errors.New("Out of range")
	}

	if start <= stop {
		return "", errors.New("Start must be greater than stop")
	}

	var output strings.Builder
	for x := start; x >= stop; x-- {
		y, _ := Verse(x)
		z := fmt.Sprintln(y)
		output.WriteString(z)
	}
	return output.String(), nil
}

func Verse(n int) (string, error) {
	if n > 2 && n < 100 {
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
	}

	if n == 2 {
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle of beer on the wall.\n", n, n, n-1), nil
	}

	if n == 1 {
		return fmt.Sprintf("%d bottle of beer on the wall, %d bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", n, n), nil
	}

	if n == 0 {
		return fmt.Sprintf("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"), nil
	}

	if n > 99 {
		return "", errors.New("Verse is too high")
	}

	return "", errors.New("unknown error")
}
