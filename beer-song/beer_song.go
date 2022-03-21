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

	if stop < 0 || start > 99 || start <= stop {
		return "", errors.New("Out of range")
	}

	var output strings.Builder
	for x := start; x >= stop; x-- {
		y, _ := Verse(x)
		output.WriteString(fmt.Sprintln(y))
	}
	return output.String(), nil
}

func Verse(n int) (string, error) {

	if n > 99 {
		return "", errors.New("Verse is too high")
	}
	s, p, count, count2 := " bottle", " bottles", strconv.Itoa(n), strconv.Itoa(n-1)
	bottles1, bottles2, takeDown := count+p, count2+p, "Take one down and pass it around,"
	bottlesFor0 := bottles1

	switch n {
	case 2:
		bottles2 = count2 + s
	case 1:
		bottles1, bottles2, takeDown = count+s, "no more"+p, "Take it down and pass it around,"
		bottlesFor0 = bottles1
	case 0:
		bottles1, bottlesFor0, bottles2, takeDown = "No more"+p, "no more"+p, "99"+p, "Go to the store and buy some more,"
	}

	return fmt.Sprintf("%s of beer on the wall, %s of beer.\n%s %s of beer on the wall.\n", bottles1, bottlesFor0, takeDown, bottles2), nil

}
