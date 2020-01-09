package main

import "fmt"

func reduce(ival int) {
	ival--
}

func reducePtr(iptr *int) {
	*iptr-- //pointer arithmatic
}

var movies map[string]float64

func main() {
	count := 10
	fmt.Println("initial:", count)

	reduce(count)
	fmt.Println("reduced:", count)

	reducePtr(&count)
	fmt.Println("reduced:", count)

	fmt.Println("address of count:", &count)

}
