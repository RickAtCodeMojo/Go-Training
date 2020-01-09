package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var txt string
	var val int

	// basic switch statment
	switch upto(4) {
	case 0:
		txt = "Zero"
	case 1:
		txt = "One"
	case 2:
		txt = "Two"
	case 3:
		txt = "Three"
	}
	fmt.Println("txt:", txt)

	//can switch on string
	switch txt {
	case "Zero":
		val = 0
	case "One":
		val = 1
	case "Two":
		val = 2
	case "Three":
		val = 3
	}
	fmt.Println("val:", val)

}

func upto(max int) int {
	var seed = rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	return r.Intn(max)
}
