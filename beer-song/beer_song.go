package beer

import (
	"errors"
	"fmt"
	"strconv"
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

	if n > 99 {
		return "", errors.New("Verse is too high")
	}

	var bottles1, bottlesFor0, bottles2, takeDown string

	switch {
	case n > 2:
		bottles1, bottlesFor0, bottles2, takeDown = strconv.Itoa(n)+" bottles", strconv.Itoa(n)+" bottles", strconv.Itoa(n-1)+" bottles", "Take one down and pass it around,"
	case n == 2:
		bottles1, bottlesFor0, bottles2, takeDown = strconv.Itoa(n)+" bottles", strconv.Itoa(n)+" bottles", strconv.Itoa(n-1)+" bottle", "Take one down and pass it around,"
	case n == 1:
		bottles1, bottlesFor0, bottles2, takeDown = strconv.Itoa(n)+" bottle", strconv.Itoa(n)+" bottle", "no more"+" bottles", "Take it down and pass it around,"
	case n == 0:
		bottles1, bottlesFor0, bottles2, takeDown = "No more"+" bottles", "no more"+" bottles", "99"+" bottles", "Go to the store and buy some more,"

	}

	return fmt.Sprintf("%s of beer on the wall, %s of beer.\n%s %s of beer on the wall.\n", bottles1, bottlesFor0, takeDown, bottles2), nil

}
