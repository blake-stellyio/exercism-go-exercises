package pythagorean

// import (
// 	"fmt"
// 	"math"
// )

type Triplet []int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var output []Triplet
	for n := 1; n < max; n++ {
		for m := n + 1; m <= max; m++ {
			a := (m * m) - (n * n)
			b := 2 * m * n
			c := m*m + n*n
			// fmt.Println(a, b, c)
			if (a*a)+(b*b) == c*c && c <= max {
				output = append(output, []int{a, b, c})
				tempTrip := []int{a, b, c}
				for i := a; i > max-2; i++ {

				}
				// 	for i := 1; i < 1000; i++ {
				// 		a *= i
				// 		b *= i
				// 		c *= i
				// 		fmt.Println(a, b, c)
				// 		if math.Signbit(float64(a)) || math.Signbit(float64(b)) || math.Signbit(float64(c)) {
				// 			fmt.Println("Something was negative!")
				// 			break
				// 		}
				// 		if a >= min && (a)*(a)+(b)*(b) == (c)*(c) && c < max {
				// 			output = append(output, []int{a, b, c})
				// 			continue
				// 		} else if a >= min && (a)*(a)+(b)*(b) == (c)*(c) && c == max {
				// 			output = append(output, []int{a, b, c})
				// 			return output
				// 		}
				// 	}

				// }
				// if c*c > max {
				// 	break
			}
		}
	}
	return output
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	panic("Please implement the Sum function")
}
