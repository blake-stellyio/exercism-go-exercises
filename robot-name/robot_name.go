package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Robot string

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const capacity int = (26 * 26 * 10 * 10 * 10)

var robotDB map[string]bool = make(map[string]bool, capacity)
var namesGenerated bool = false
var nameDB []string

func (r *Robot) Name() (string, error) {

	var output strings.Builder

	if *r != "" {
		return string(*r), nil
	}

	if len(robotDB) >= capacity {
		return "", errors.New("name capacity exhausted")
	}

	if !namesGenerated {
		nameDB = GenerateNames()
		namesGenerated = true
	}

	position := rand.Intn(len(nameDB))
	output.WriteString(nameDB[position])
	nameDB[position] = nameDB[0]
	nameDB = nameDB[1:]

	robotDB[output.String()] = true

	*r = Robot(output.String())
	return output.String(), nil

}

func (r *Robot) Reset() {
	*r = ""
	r.Name()
}

func GenerateNames() []string {

	var num, let, output []string

	for x := 0; x < 1000; x++ {
		if x < 10 {
			num = append(num, fmt.Sprintf("00%d", x))
		}

		if x >= 10 && x < 100 {
			num = append(num, fmt.Sprintf("0%d", x))
		}

		if x > 99 {
			num = append(num, fmt.Sprint(x))
		}
	}

	for x := 0; x < 26; x++ {
		for i := 0; i < 26; i++ {
			let = append(let, fmt.Sprintf("%c%c", alphabet[x], alphabet[i]))
		}
	}

	for _, x := range let {
		for _, i := range num {
			output = append(output, fmt.Sprintf("%s%s", x, i))
		}
	}

	return output

}
