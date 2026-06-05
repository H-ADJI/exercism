package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(foodCalc FodderCalculator, nCows int) (float64, error) {
	amount, err := foodCalc.FodderAmount(nCows)
	if err != nil {
		return 0, err
	}
	fatFactor, err := foodCalc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return amount * fatFactor / float64(nCows), nil

}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(foodCalc FodderCalculator, nCows int) (float64, error) {
	if nCows > 0 {
		return DivideFood(foodCalc, nCows)
	}
	return 0, errors.New("invalid number of cows")
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	nCows   int
	message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.nCows, err.message)
}
func ValidateNumberOfCows(ncows int) error {
	if ncows > 0 {
		return nil
	}
	if ncows == 0 {
		return &InvalidCowsError{nCows: 0, message: "no cows don't need food"}
	}
	return &InvalidCowsError{nCows: ncows, message: "there are no negative cows"}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
