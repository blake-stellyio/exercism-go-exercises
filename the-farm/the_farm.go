package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	if err == ErrScaleMalfunction && fodder > 0 {
		fodder *= 2
		err = nil
	}

	if err != nil && fodder > 0 {
		return 0, err
	}

	if fodder < 0 {
		err = errors.New("negative fodder")
		return 0, err
	}

	if cows == 0 {
		err = errors.New("division by zero")
		return 0, err
	}

	if cows < 0 {
		err = errors.New(fmt.Sprintf("silly nephew, there cannot be %d cows", cows))
		return 0, err
	}

	fodderPerCow := fodder / float64(cows)

	return fodderPerCow, nil

}
