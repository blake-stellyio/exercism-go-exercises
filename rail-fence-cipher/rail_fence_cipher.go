package railfence

import (
	"fmt"
	"strings"
)

func Encode(message string, rails int) string {

	railsIdent := make(map[int][]string)

	var counter int = 1
	var previousCounter int

	for _, x := range message {

		if previousCounter > counter && counter != 1 || counter == rails {
			railsIdent[counter] = append(railsIdent[counter], string(x))
			previousCounter = counter
			counter--
		} else if previousCounter < counter && counter != rails || counter == 1 {
			railsIdent[counter] = append(railsIdent[counter], string(x))
			previousCounter = counter
			counter++
		}

	}

	var output string

	for i := 0; i <= rails; i++ {
		output = output + strings.Join(railsIdent[i], "")
	}

	fmt.Println(output)

	return output
}

func Decode(message string, rails int) string {

	railsDec := make(map[int][]string)

	var remainder int

	if rails%2 == 0 {
		remainder = (len(message) % rails)
	} else {
		remainder = ((len(message) + 1) % rails)
	}
	inside := ((len(message) - remainder) / (rails - 1))
	outside := inside / 2

	var counter int = 1
	var whichInside int = 2

	for i, x := range message {
		if i < outside+remainder-1 {
			railsDec[1] = append(railsDec[1], string(x))
		}
		if i >= (outside+remainder-1) && i < len(message)-(outside) && rails-2 > 0 {
			if counter <= inside+remainder {
				railsDec[whichInside] = append(railsDec[whichInside], string(x))
				counter++
				continue
			} else {
				counter = 1
				whichInside += 1
				railsDec[whichInside] = append(railsDec[whichInside], string(x))
				continue
			}
		}

		if i >= len(message)-(outside) {
			railsDec[rails] = append(railsDec[rails], string(x))
		}

	}

	fmt.Println(railsDec)

	position := 0
	counter = 1
	previousCounter := 0
	var output []string

	for x := 0; x <= inside-1; x++ {

		if x == 0 {
			output = append(output, railsDec[1][x])
			continue
		}

		for i := 1; railsDec[i][0] != ""; i++ {

		}

	}

	return strings.Join(output, "")
}
