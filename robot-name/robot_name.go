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
var numberCounter, letterCounter map[string]int = make(map[string]int, 1000), make(map[string]int, 26*26)
var namesGenerated bool = false
var numberDB, letterDB []string

func (r *Robot) Name() (string, error) {

	if *r != "" {
		return string(*r), nil
	}

	if len(robotDB) >= capacity {
		return "", errors.New("name capacity exhausted")
	}

	if !namesGenerated {
		numberDB, letterDB = GenerateNames()
		namesGenerated = true
	}

	var output strings.Builder
	// var repeatCounter int = 0

	for {

		var numPosition, letPosition int

		if len(numberDB) > 0 && len(letterDB) > 0 {
			numPosition = rand.Intn(len(numberDB))
			letPosition = rand.Intn(len(letterDB))
		} else {
			break
		}

		letters, numbers := letterDB[letPosition], numberDB[numPosition]

		if numberCounter[numbers] == 26*26 {
			numberDB[numPosition] = numberDB[0]
			numberDB = numberDB[1:]
			continue
		}

		if numberCounter[letters] == 1000 {
			letterDB[letPosition] = letterDB[0]
			letterDB = letterDB[1:]
			continue
		}

		output.WriteString(fmt.Sprintf("%s%s", letters, numbers))

		if !robotDB[output.String()] {
			robotDB[output.String()] = true
			numberCounter[numbers] = numberCounter[numbers] + 1
			numberCounter[letters] = numberCounter[letters] + 1
			break
		}
		// repeatCounter++
		// fmt.Println(repeatCounter)
		output.Reset()

	}
	// fmt.Println("Made it here, ", output.String())
	*r = Robot(output.String())
	return output.String(), nil
}

func (r *Robot) Reset() {
	*r = ""
	r.Name()
}

func GenerateNames() ([]string, []string) {

	var outputNum, outputLet []string

	for x := 0; x < 1000; x++ {
		if x < 10 {
			outputNum = append(outputNum, fmt.Sprintf("00%d", x))
		}

		if x >= 10 && x < 100 {
			outputNum = append(outputNum, fmt.Sprintf("0%d", x))
		}

		if x > 99 {
			outputNum = append(outputNum, fmt.Sprint(x))
		}
	}

	for x := 0; x < 26; x++ {
		for i := 0; i < 26; i++ {
			outputLet = append(outputLet, fmt.Sprintf("%c%c", alphabet[x], alphabet[i]))
		}
	}

	return outputNum, outputLet

}
