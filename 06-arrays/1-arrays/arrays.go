package main

import (
	"fmt"
	"math"
)

func main() {
	var results [7]float64
	fmt.Println("results:", results)
	// The builtin len function returns the length of an array.
	fmt.Println("number of elemenst in 'results':", len(results))
	// The builtin cap function returns the length of an array.
	fmt.Println("capacity of 'results':", cap(results))

	results[4] = 90.0 * math.Pi
	print("results[%d] = %g\n", results)
	fmt.Println("element 5:", results[4])

	// //arrays are passed by value (copied)
	// update(results)
	// print("results[%d] is still %g\n:", results)

}
func print(message string, days [7]float64) {
	for i, f := range days {
		fmt.Printf(message, i, f)
	}
}

func update(days [7]float64) {
	for i := 0; i < len(days); i++ {
		days[i] = float64(i*10) * math.Pi //the copy is being updated
	}
	print(" copy of results[%d] = %g\n", days)
}
