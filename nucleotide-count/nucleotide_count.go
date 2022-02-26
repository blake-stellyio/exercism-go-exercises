package dna

import "errors"

type Histogram map[rune]int

type DNA []rune

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, x := range d {
		if x != 'A' && x != 'C' && x != 'G' && x != 'T' {
			err := errors.New("not a vaild nucleotide")
			return h, err
		} else {
			h[x]++
		}
	}
	return h, nil
}
