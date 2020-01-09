package main

import (
	"fmt"
)

var balance string
var ok bool
var max uint8

func main() {
	fmt.Printf("initial values: >%s<, %t, %d\n", balance, ok, max)
	max = 255
	balance = "Balance:"
	money := 53273.95
	fmt.Printf("balance %f\n", money)
}

/*
counter := 1
	for {
		counter++
		money -= float64(counter * 1000)
		if counter > 5 {
			break
		}
	}
	ok := money > 40000
	result := "No"
	if ok {
		result = "Yes"
	}
	fmt.Println("Do we have more than 40000?", result)
	var x, y uint8
	x, y = max, max
	x, y = y/5, x/51
	fmt.Println(x, y)

*/
