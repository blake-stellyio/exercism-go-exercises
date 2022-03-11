package binarysearch

func SearchInts(list []int, key int) int {

	if len(list) < 1 {
		return -1
	}

	low, high := 0, len(list)-1
	mid := low + (high-low)/2

	for i := 0; i <= (len(list) / 2); i++ {

		switch {
		case key > list[mid]:
			low = mid
		case key < list[mid]:
			high = mid
		}

		if key == list[mid] {
			return mid
		}

		if mid == len(list)-2 {
			mid++
		} else {
			mid = low + (high-low)/2
		}
	}

	return -1
}
