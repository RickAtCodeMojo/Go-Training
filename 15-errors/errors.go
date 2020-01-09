package main

import (
	"errors"
	"fmt"
	"math"
)

//https://blog.golang.org/error-handling-and-go

//NegativeSquareRootError returns an error specialized for this type of math error
type NegativeSquareRootError struct {
	message string
}

//NewNegativeSquareRootError returns a new error of this kind of error
func NewNegativeSquareRootError(message string) *NegativeSquareRootError {
	return &NegativeSquareRootError{
		message: message,
	}
}
func (e *NegativeSquareRootError) Error() string {
	return e.message
}
func main() {
	// print(math.Pi)
	z := getZ(7.5, 2.8)
	z, err := check(z)
	fmt.Printf("Result: %.2f, %s.\n", z, err)

	// fmt.Println("Better")
	// z = getZ(7.5, 2.8)
	// z, err = better(z)
	// fmt.Printf("Result: %.2f, %s.\n", z, err)

	fmt.Println("Best")
	z = getZ(7.5, 2.8)
	z, err = best(z)
	switch err.(type) {
	case *NegativeSquareRootError:
		{
			fmt.Println(err.Error())
			//try some recovery effort...
			fmt.Println("Doing something about the value of z...")
		}
	default:
		{
			fmt.Println("I guess an error occured!")
		}
	}
}
func getZ(x float64, y float64) float64 {
	return x + math.Sin(2.0*math.Pi/y)
}

func nan(x float64) float64 {
	return math.Sqrt(1.0 - math.Pow(x, 2.0))
}

func check(x float64) (float64, error) {
	z := nan(x)
	if math.IsNaN(z) {
		return 0, errors.New("cannot calculate the square root of a negative number")
	}

	return z, nil

}
func better(x float64) (float64, error) {
	of := 1.0 - math.Pow(x, 2.0)
	if of < 0.0 {
		return 0, fmt.Errorf("cannot calculate the square root of a negative number: (1.0 - math.Pow(x, 2.0) = %g, where x = %g", of, x)
	}
	return math.Sqrt(of), nil

}

func best(x float64) (float64, error) {
	of := 1.0 - math.Pow(x, 2.0)
	if of < 0.0 {
		message := fmt.Sprintf("cannot calculate the square root of a negative number: (1.0 - math.Pow(x, 2.0) = %g, where x = %g", of, x)
		return 0, NewNegativeSquareRootError(message)
	}
	return math.Sqrt(of), nil

}

/*
try to handle errors immediately

don't do this - it's fragile:
func someFunc() (int, error){
	err:=NewError()
		:
		:
	return 0, err
}
*/
