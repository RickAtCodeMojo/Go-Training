package main

//https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/
//http://tleyden.github.io/blog/2016/12/20/understanding-function-closures-in-go/
import "fmt"

func increase() func() int { //returns a func that returns an in "func() int"
	i := 0              // <-- initializes to 0 after each assignment
	return func() int { //returns the function
		i++
		return i
	}
}

func sum() func(int) int { //returns a func that takes an int and sums it
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	next := increase() //assignment i = 0

	fmt.Println(next()) //increments to 1 as nextInt is called
	fmt.Println(next()) //2
	fmt.Println(next()) //3

	sequence := increase()  //fresh assignment i = 0
	fmt.Println(sequence()) //increments to 1

	fmt.Println(next())

	// increment, decrement := sum(), sum()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(increment(i), decrement(-2*i))
	// }

}
