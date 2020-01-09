package main

import (
	"fmt"
	"sync"
)

func do(work string, wg *sync.WaitGroup) {
	//wg.Add(1)
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("Doing", work, i, "times")
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func(msg string) {
		defer wg.Done() //is aware of wg!
		for i := 0; i < 3; i++ {
			fmt.Printf("%d  anonymous function with parameter: %s\n", i, msg)
		}
	}("hello world")
	go do("stuff", &wg)
	wg.Wait()
	fmt.Println("Returning from main")
}
