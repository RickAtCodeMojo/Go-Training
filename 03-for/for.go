package main

import (
	"fmt"
	"time"
)

func main() {

	//one condition
	fmt.Println("One condition, i is less than 5")
	i := 0
	for i < 5 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Print("\n\n")

	//classic for loop
	fmt.Println("Classic for loop")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\n")

	//continue
	fmt.Println("Using continue")
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			fmt.Printf("continue:")
			pause()
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Print("\n")

	//endless loop
	fmt.Println("Endless loop. ")
	for {
		fmt.Printf(". ")
		pause()
		i++
		if i > 9 {
			fmt.Print(" break")
			break
		}
	}
	fmt.Printf("\n\n")
}
func pause() {
	time.Sleep(250 * time.Millisecond)

}
