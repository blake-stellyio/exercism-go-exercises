package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	// "strconv"
	"strings"
)

type Robot string

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const capacity int = (26 * 26 * 10 * 10 * 10)

var robotDB map[string]bool = make(map[string]bool, capacity)
var numberCounter map[string]int = make(map[string]int, 1000)
var namesGenerated bool = false
var numberDB []string

func (r *Robot) Name() (string, error) {

	if *r != "" {
		return string(*r), nil
	}

	if len(robotDB) >= capacity {
		return "", errors.New("name capacity exhausted")
	}

	if !namesGenerated {
		numberDB = GenerateNumbers()
	}

	var output strings.Builder

	for {

		var position int

		if len(numberDB) > 0 {
			position = rand.Intn(len(numberDB))
		} else {
			break
		}

		number := numberDB[position]

		if numberCounter[number] == 26*26 {
			numberDB[position] = numberDB[0]
			numberDB = numberDB[1:]
			continue
		}

		for i := 1; i <= 2; i++ {
			output.WriteByte(alphabet[rand.Intn(26)])
		}

		output.WriteString(number)

		if !robotDB[output.String()] {
			robotDB[output.String()] = true
			numberCounter[number] = numberCounter[number] + 1
			break
		}

		output.Reset()
	}
	fmt.Println("Made it here, ", output.String())
	*r = Robot(output.String())
	return output.String(), nil
}

func (r *Robot) Reset() {
	*r = ""
	r.Name()
}

func GenerateNumbers() []string {
	var output []string
	for x := 0; x < 1000; x++ {
		if x < 10 {
			output = append(output, fmt.Sprintf("00%d", x))
		}

		if x >= 10 && x < 100 {
			output = append(output, fmt.Sprintf("0%d", x))
		}

		if x > 99 {
			output = append(output, fmt.Sprint(x))
		}
	}
	return output
}
