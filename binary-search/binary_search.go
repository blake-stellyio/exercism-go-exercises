package binarysearch

func SearchInts(list []int, key int) int {

	if len(list) < 1 {
		return -1
	}

	if key >= list[len(list)/2] {
		for i := len(list) / 2; i <= len(list)-1; i++ {
			if list[i] == key {
				return i
			}
		}
	} else {
		for i := 0; i <= len(list)-1; i++ {
			if list[i] == key {
				return i
			}
		}
	}
	return -1
}
