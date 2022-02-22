package reverse

import (
	"strings"
)

func Reverse(input string) string {
	reverseString := make([]string, len(input))
	for i, r := range input {
		reverseString[len(input)-i-1] = string(r)
	}
	return strings.Join(reverseString, "")
}
