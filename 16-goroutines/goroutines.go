package main

import "fmt"

func do(work string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Doing", work, i, "times")
	}
}

func main() {

	do("stuff")

	go do("other stuff")

	go func(msg string) {
		for i := 0; i < 3; i++ {
			fmt.Printf("%d  anonymous function with parameter: %s\n", i, msg)
		}
	}("hello world")

	fmt.Scanln() //wait for input
	fmt.Println("Returning from main")
}
