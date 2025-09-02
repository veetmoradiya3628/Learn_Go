package thefarm
import ("fmt"; "errors")

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, noOfCows int) (float64, error) {
   	fa, err := fc.FodderAmount(noOfCows);
    if err != nil {
        return 0, err
    }
    ff, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    return (fa * ff) / float64(noOfCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, noOfCows int) (float64, error) {
	if noOfCows <= 0 {
        return 0, errors.New("invalid number of cows")
    }
    return DivideFood(fc, noOfCows)
}
// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
    noOfCows int
    message string
}
func (e *InvalidCowsError) Error() string {
  return fmt.Sprintf("%d cows are invalid: %s", e.noOfCows, e.message)
}
func ValidateNumberOfCows(noOfCows int) error {
	if noOfCows < 0 {
        return &InvalidCowsError{
            noOfCows: noOfCows,
            message: "there are no negative cows",
        }
    }
    if noOfCows == 0 {
        return &InvalidCowsError{
            noOfCows: noOfCows,
            message: "no cows don't need food",
        }
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
