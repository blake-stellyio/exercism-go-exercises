package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	s string
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	if cows == 0 {
		err := errors.New("division by zero")
		return 0, err
	}

	if cows < 0 {
		err := errors.New(fmt.Sprintf("silly nephew, there cannot be %d cows", cows))
		return 0, err
	}

	fodder, err := weightFodder.FodderAmount()

	if err == ErrScaleMalfunction {
		if fodder > 0 {
			fodder *= 2
			err = nil
		} else {
			err = errors.New("negative fodder")
			return 0, err
		}
	}

	if err != nil {
		return 0, err
	}

	if fodder < 0 {
		err = errors.New("negative fodder")
		return 0, err
	}

	fodderPerCow := fodder / float64(cows)

	return fodderPerCow, nil

}
