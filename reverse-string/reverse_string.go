package reverse

import (
	"fmt"
)

func Reverse(input string) string {
	var reverseString []rune
	for i := range input {
		reverseChar := rune(input[len(input)-(i+1)])
		reverseString = append(reverseString, reverseChar)
	}
	return fmt.Sprint(reverseString)
}

/*
func Reverse(input string) string {
	var reverseString []string
	stringInRune := []rune(input)
	for i := range stringInRune {
		reverseChar := string(stringInRune[len(input)-(i+1)])
		reverseString = append(reverseString, reverseChar)
	}
	return strings.Join(reverseString, "")
}
*/
