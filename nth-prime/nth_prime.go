package prime

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	qtyPrimes := 1
	for x := 2; ; x++ {
		if isPrime(x) {
			if qtyPrimes == n {
				return x, true
			} else {
				qtyPrimes++
			}
		}

	}
}

func isPrime(p int) bool {
	for x := 2; x < p; x++ {
		if p%x == 0 {
			return false
		}
	}
	return true
}
