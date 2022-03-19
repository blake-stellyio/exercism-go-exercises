package robotname

import (
	// "fmt"
	"errors"
	// "fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Robot string

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const capacity int = 26 * 26 * 10 * 10 * 10

var capacityCounter int = 0
var robotDB map[string]bool = make(map[string]bool, capacity)

func (r *Robot) Name() (string, error) {

	if capacityCounter == capacity {
		return "", errors.New("namespace exhausted")
	}

	var output strings.Builder
	for {

		for i := 1; i <= 2; i++ {
			output.WriteByte(alphabet[rand.Intn(25)])
		}

		output.WriteString(strconv.Itoa(rand.Intn(999-100) + 100))

		if !robotDB[output.String()] {
			robotDB[output.String()] = true
			capacityCounter++
			break
		}
		output.Reset()
	}
	return output.String(), nil
}

func (r *Robot) Reset() {
	// robotDB[fmt.Sprint(r)] = false
	*r = ""
}
