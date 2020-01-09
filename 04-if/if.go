package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// i := 0
	// // basic if statment
	// if i < 1 {
	// 	fmt.Println("i is less than 1")
	// }

	// //else
	// i = upto(9)
	// if i > 4 {
	// 	fmt.Println("i is greater than 4")
	// } else {
	// 	fmt.Println("i is less than or equal to 4")
	// }

	//else if
	//you can put a statement before the conditional
	//and you can have multiple branches
	if i := upto(9); i > 6 {
		fmt.Printf("%d is greater than 6\n", i)
	} else if i > 3 {
		fmt.Printf("%d is greater than 3\n", i)
	} else {
		fmt.Printf("%d is less than or equal to 3\n", i)
	}

}

func upto(max int) int {
	var seed = rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	return r.Intn(max)
}
