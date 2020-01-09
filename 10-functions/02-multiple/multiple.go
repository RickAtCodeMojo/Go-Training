package main

import "fmt"

var max float64

func main() {
	var ok bool
	amount := 2.99
	max, ok = bid(amount, max)
	if ok {
		fmt.Println(amount, "is the winning bid so far")
	} else {
		fmt.Println(max, "is still the winning bid")
	}
	next := 1.99
	_, ok = bid(next, max)
	if ok {
		fmt.Println(next, "is the winning bid so far")
	} else {
		fmt.Println(max, "is still the winning bid")
	}

	last := 3.99
	max, _ = bid(last, max)
	if last > max {
		fmt.Println(last, "is the winning bid so far")
	}
	_, ok = bid(last, max)
	fmt.Println(ok, " last is greater than max")
}

func bid(amount float64, max float64) (float64, bool) {
	if amount > max {
		return amount, true
	}
	return max, false
}
