package main

import (
	"fmt"
	"time"
)

func main() {

	//switch on time
	now := time.Now()
	switch now.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Day off")
	default:
		fmt.Println("Work day")
	}

	//switch without a an expression is like an if else
	switch {
	case now.Hour() > 11 && now.Hour() < 17:
		fmt.Println("Core Hours")
	default:
		fmt.Println("You don't have to be here")
	}

}
