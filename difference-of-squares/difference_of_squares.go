package diffsquares

func SquareOfSum(n int) int {

	digits := []int{n}
	x := n
	for x-1 >= 0 {
		digits = append(digits, x-1)
		x -= 1
	}
	var sum int
	for _, y := range digits {
		sum += y
	}
	result := sum * sum
	return result
}

func SumOfSquares(n int) int {
	digits := []int{n}
	x := n
	for x-1 >= 0 {
		digits = append(digits, x-1)
		x -= 1
	}
	var sum int
	for _, y := range digits {
		sum += y * y
	}
	return sum
}
func Difference(n int) int {
	result := SquareOfSum(n) - SumOfSquares(n)
	return result
}
