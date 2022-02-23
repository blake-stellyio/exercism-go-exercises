package prime

func Nth(n int) (int, bool) {
	var qtyPrimes int
	var result int
	for x := 2; qtyPrimes <= n; x++ {
		if isPrime(x) {
			qtyPrimes = +1
		}
		if isPrime(x) || qtyPrimes == n {
			result = x
			break
		}
	}
	return result, true
}

func isPrime(p int) bool {
	for x := 2; x < p; x++ {
		if p == 2 {
			return true
		}
		if p%x == 0 {
			return false
		}
	}
	return true
}
