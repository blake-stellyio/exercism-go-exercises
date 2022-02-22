package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var sum int = 0
	for i := 1; i < limit; i++ {
		for _, x := range divisors {
			if x == 0 {
				continue
			} else if i%x == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
