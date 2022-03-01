package pythagorean

import (
	"fmt"
	"math"
)

type Triplet []int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var output []Triplet
	var a int
	var b int
	var c int
	var aInitial int
	var bInitial int
	var cInitial int
	for n := 1; n <= max; n++ {
		for m := n + 1; c <= max; m++ {
			a = (m * m) - (n * n)
			b = 2 * m * n
			c = m*m + n*n
			fmt.Println(a, b, c)
			aInitial = a
			bInitial = b
			cInitial = c
			if (a*a)+(b*b) == c*c && c <= max {
				if a >= min {
					output = append(output, []int{a, b, c})
				}
				// tempTrip := []int{a, b, c}
				for i := 1; c <= max; i++ {
					a += aInitial * i
					b += bInitial * i
					c += cInitial * i
					fmt.Println(a, b, c, "in nested loop")
					if math.Signbit(float64(a)) || math.Signbit(float64(b)) || math.Signbit(float64(c)) {
						fmt.Println("Something was negative!")
						break
					}
					if a >= min && (a)*(a)+(b)*(b) == (c)*(c) && c < max {
						output = append(output, []int{a, b, c})
						fmt.Println(a, b, c, "I appened to output with c less than max")
						continue
					} else if a >= min && (a)*(a)+(b)*(b) == (c)*(c) && c == max {
						output = append(output, []int{a, b, c})
						fmt.Println(a, b, c, "I appened to output with c less equal to max")
						return output
					}
				}

			}
			// fmt.Println("made it")
		}
		fmt.Println("made it")

	}
	return output
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var output []Triplet
	for a := 1; a < p-2; a++ {
		for b := a + 1; b < p-1; b++ {
			c := p - (a + b)
			if (a*a)+(b*b) == c*c && a+b+c == p {
				output = append(output, []int{a, b, c})
				break
			}
		}
	}
	return output
}
