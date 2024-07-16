package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(calculator FodderCalculator, cowCount int) (float64, error) {
	totalFodder, err := calculator.FodderAmount(cowCount)
	if err != nil {
		return 0, err
	}

	ff, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return totalFodder * ff / float64(cowCount), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calculator FodderCalculator, cowCount int) (float64, error) {
	if cowCount < 1 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(calculator, cowCount)
}

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %v", err.numberOfCows, err.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cowCount int) error {
	if cowCount < 0 {
		return &InvalidCowsError{numberOfCows: cowCount, message: "there are no negative cows"}
	}

	if cowCount == 0 {
		return &InvalidCowsError{numberOfCows: cowCount, message: "no cows don't need food"}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
