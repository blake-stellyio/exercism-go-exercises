package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Robot string

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const capacity int = (26 * 26 * 10 * 10 * 10)

var robotDB map[string]bool = make(map[string]bool, capacity)

func (r *Robot) Name() (string, error) {

	if *r != "" {
		return string(*r), nil
	}

	if len(robotDB) >= capacity {
		return "", errors.New("name capacity exhausted")
	}

	var output strings.Builder

	for {

		for i := 1; i <= 2; i++ {
			output.WriteByte(alphabet[rand.Intn(26)])
		}

		output.WriteString(strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)))

		if !robotDB[output.String()] {
			robotDB[output.String()] = true
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
