package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c1 <- "one"
	c2 <- "two"
	// go func() {
	// 	// for m := 0; m < 2; m++ { //<--what happens if do a couple of these
	// 	time.Sleep(1000 * time.Millisecond) //simulate some work being done
	// 	// }
	// }()
	// go func() {
	// 	time.Sleep(100 * time.Millisecond) //simulate some work being done
	// }()

	// for i := 0; i < 3; i++ { / /<--why do we have to change the number of
	//iterations to see all of the items in both channels?
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recieved:", msg1)
		case msg2 := <-c2:
			fmt.Println("recieved:", msg2)
		}
	}

}
