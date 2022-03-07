package prime

func Factors(n int64) []int64 {
	var primeFactors []int64
	for i := int64(2); i <= n; i++ {
		for n%i == 0 {
			if isPrime(i) {
				primeFactors = append(primeFactors, i)
			}
			n = n / i
		}
	}
	return primeFactors
}

func isPrime(n int64) bool {
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}

	}
	return n > 1
}
