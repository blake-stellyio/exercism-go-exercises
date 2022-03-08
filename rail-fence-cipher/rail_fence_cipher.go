package railfence

import (
	"fmt"
	"strings"
)

func Encode(message string, rails int) string {

	railsIdent := make(map[int][]string)

	// for i, x := range message {
	// 	railsIdent[i%rails] = append(railsIdent[i%rails], string(x))

	// }

	// fmt.Println(railsIdent)

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
	panic("Please implement the Decode function")
}
