package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	var err error
	if len(digits) < span {
		err = errors.New("span cannot be longer than the series of digits")
		return 0, err
	} else if digits == "" && span == 0 {
		return 1, nil
	} else if span < 0 {
		err = errors.New("the span may only include positive integers")
		return 0, err
	}

	digitSlice := []int{}

	for _, i := range digits {
		if !unicode.IsDigit(i) {
			err = errors.New("the series of digits may only include positive integers")
		}
		x, _ := strconv.Atoi(string(i))
		if x < 0 {
			err = errors.New("the series of digits may only include positive integers")
			return 0, err
		}
		digitSlice = append(digitSlice, x)
	}

	var largestProduct int = 0

	for x := range digitSlice {
		var tempFactors []int
		if x+span > len(digitSlice) {
			break
		}
		tempFactors = digitSlice[x : x+span]
		var product int = 1
		if len(tempFactors) == span {
			for _, q := range tempFactors {
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

	return int64(largestProduct), err

}
