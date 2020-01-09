package main

import (
	"fmt"
)

func main() {
	fmt.Println("Result", sum(1, 2, 3, 4, 5))
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Result", sum(75, numbers...))
	//fmt.Println("Another result", sum(1, 2, 3, 4))
}

func sum(money float64, numbers ...int) {
	s := 0
	for _, f := range numbers {
		s += f
	}
	fmt.Printf("\n%T\n", numbers)
	return s, 75.5
}
