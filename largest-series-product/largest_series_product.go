package lsproduct

func LargestSeriesProduct(digits string, span int) (int64, error) {
	digitSlice := []int{}

	for _, i := range digits {
		digitSlice = append(digitSlice, int(i))
	}

	var largestProduct int = 0

	for x := range digitSlice {
		var tempFactors []int
		for z := x; z <= x+span; z++ {
			if z > len(digitSlice)-1 {
				break
			}
			tempFactors = append(tempFactors, digitSlice[z])
		}
		var product int = 1
		if len(tempFactors) == span {
			for q := range tempFactors {
				product = product * q
			}
			if product <= largestProduct {
				continue
			} else {
				largestProduct = product
			}
		} else {
			break
		}

	}

	return int64(largestProduct), nil

}
