package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(1000 * time.Millisecond)
		c1 <- "completed: 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)

	case <-time.After(500 * time.Millisecond):
		fmt.Println("timed out: 1")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(1000 * time.Millisecond)
		c2 <- "completed: 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)

	case <-time.After(1500 * time.Millisecond):
		fmt.Println("timed out: 2")
	}
}

/*

result should show :

time out: 1
completed: 2

the first operation doesn't complete because it timed out
the second completes because it finishes before the timer goes off
*/
