package clock

import (
	"fmt"
	"time"
)

type Clock string

const clockForm = "03:04"

func New(h, m int) Clock {
	initialTime := time.Date(0000, 00, 00, h, m, 00, 00, time.FixedZone("null", 0))
	var output Clock = Clock(fmt.Sprintln(initialTime))
	fmt.Println(output)
	return output
}

func (c Clock) Add(m int) Clock {
	panic("Please implement the Add function")
}

func (c Clock) Subtract(m int) Clock {
	panic("Please implement the Subtract function")
}

func (c Clock) String() string {
	panic("Please implement the String function")
}
