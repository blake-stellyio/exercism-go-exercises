package prime

import (
	"fmt"
	"math"
)

func Factors(n int64) []int64 {
	var primeFactors []int64
	var i int64
	for i = 2; i <= n; i++ {
		if n == i && len(primeFactors) > 1 {
			break
		}
		if n%i == 0 {
			if isPrime(i) {
				primeFactors = append(primeFactors, i)

			}
			// fmt.Println(i)
			tempFactor := n / i
			for tempFactor%i == 0 {
				if isPrime(i) {
					primeFactors = append(primeFactors, i)
				}
				tempFactor = tempFactor / i
			}
		}
	}
	fmt.Println(primeFactors)
	return primeFactors
}

func isPrime(n int64) bool {
	fmt.Println(n)
	var i int64
	for i = 2; i <= int64(math.Floor(math.Sqrt(float64(n)))); i++ {
		// fmt.Println(i)
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
