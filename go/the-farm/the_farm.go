package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator , quantity int) (float64, error) {

	amount , err :=	fc.FodderAmount(quantity)

	if err != nil{
		return 0 ,err
	}

	factor , err := fc.FatteningFactor()

	if err != nil {
		return  0 ,err
	}

	foodPerCow := amount * factor / float64(quantity)

	return foodPerCow , nil


}

// TODO: define the 'ValidateInputAndDivideFood' function

func ValidateInputAndDivideFood(fc FodderCalculator , quantity int) (float64 , error){
		if quantity > 0 {
			return DivideFood(fc , quantity)
		} else {
			return  0 , errors.New("invalid number of cows")
		}
	
}

type InvalidCowsError struct{
	message string
	cows int
}


// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(quantity int) error {
	if quantity < 0 {
		return &InvalidCowsError{
			cows:    quantity,
			message: "there are no negative cows",
		}
	}

	if quantity == 0 {
		return &InvalidCowsError{
			cows:    quantity,
			message: "no cows don't need food",
		}
	}

	return nil
}
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
