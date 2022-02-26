package pythagorean

type Triplet []int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var sides Triplet
	for m := min; m < max; m++ {
		for n := m + 1; n <= max; n++ {
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n
			if a*a+b*b == c*c && c*c <= max {
				tempSides := []int{a, b, c}
				sides = append(sides, tempSides)
				// ? having trouble figuring out how to return the answer in the triplet format
			}
		}
	}
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	panic("Please implement the Sum function")
}
