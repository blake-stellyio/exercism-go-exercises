package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {

	subjectChars := strings.Split(strings.ToLower(subject), "")
	sort.Strings(subjectChars)

	var anagams []string
	for _, i := range candidates {

		if strings.EqualFold(subject, i) {
			continue
		}

		candidateChars := strings.Split(strings.ToLower(i), "")
		sort.Strings(candidateChars)

		if len(subjectChars) != len(candidateChars) {
			continue
		}

		var isAnagram bool
		for i := range subjectChars {
			if subjectChars[i] != candidateChars[i] {
				isAnagram = false
				break
			} else {
				isAnagram = true
			}
		}
		if isAnagram {
			anagams = append(anagams, i)
		}

	}
	return anagams
} 
