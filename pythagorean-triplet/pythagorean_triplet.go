package pythagorean

type Triplet []int

func Range(min, max int) []Triplet {
	var output []Triplet
	var c int
	for n := 1; n <= max; n++ {
		for m := n + 1; c <= max; m++ {
			a := (m * m) - (n * n)
			b := 2 * m * n
			c = m*m + n*n
			aInitial, bInitial, cInitial := a, b, c
			if (a*a)+(b*b) == c*c && c <= max {
				if a >= min {
					output = append(output, []int{a, b, c})
				}
				for i := 1; c <= max; i++ {
					a += aInitial * i
					b += bInitial * i
					c += cInitial * i
					if a >= min && (a)*(a)+(b)*(b) == (c)*(c) && c == max {
						output = append(output, []int{a, b, c})
						return output
					}
				}

			}
		}
	}
	return output
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var output []Triplet
	for a := 1; a <= p/3; a++ {
		for b := a + 1; b <= p/2; b++ {
			c := p - (a + b)
			if (a*a)+(b*b) == c*c && a+b+c == p {
				output = append(output, []int{a, b, c})
				break
			}
		}
	}
	return output
}
