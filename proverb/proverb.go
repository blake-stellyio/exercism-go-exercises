package proverb

import "fmt"

func Proverb(rhyme []string) []string {

	if len(rhyme) == 0 {
		return rhyme
	}

	var output []string

	for i := 0; i < len(rhyme)-1; i++ {
		if len(rhyme) > 1 {
			output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}

	output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return output
}
