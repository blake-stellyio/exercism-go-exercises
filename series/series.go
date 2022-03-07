package series

import (
	"fmt"
	"strings"
)

func All(n int, s string) []string {
	var output []string
	for x := range s {
		// fmt.Println(x)
		var tempList []string
		for i := x; i < x+n && i < len(s); i++ {
			fmt.Println(i)
			tempList = append(tempList, string(s[i]))
			if len(tempList) == n {
				output = append(output, strings.Join(tempList, ""))
			}
		}
	}
	return output
}

func UnsafeFirst(n int, s string) string {
	for x := range s {
		// fmt.Println(x)
		var tempList []string
		for i := x; i < x+n && i < len(s); i++ {
			fmt.Println(i)
			tempList = append(tempList, string(s[i]))
			if len(tempList) == n {
				output := strings.Join(tempList, "")
				return output
			}
		}
	}
	return "failed"
}
