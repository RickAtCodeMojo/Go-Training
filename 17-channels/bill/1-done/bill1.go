package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	done := make(chan bool)
	go func() {
		p := <-ch // Receive
		fmt.Println(p)
		done <- true
	}()
	ch <- "paper" // Send
	// ok := <-done
	<-done
}
