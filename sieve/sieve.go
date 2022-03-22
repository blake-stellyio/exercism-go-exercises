package sieve

func Sieve(limit int) []int {

	if limit < 2 {
		return []int{}
	}

	var nonPrimes []int
	var primes []int

	for x := 2; x <= limit; x++ {

		var listed bool = false

		for _, y := range nonPrimes {
			if y == x {
				listed = true
				continue
			}
		}

		if listed == true {
			continue
		}

		p, counter, y := x, 2, 0

		for y <= limit {
			y = counter * p
			nonPrimes = append(nonPrimes, y)
			counter++
		}

		primes = append(primes, x)
	}
	return primes
}
