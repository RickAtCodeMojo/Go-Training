package main

import (
	"fmt"
	"time"
)

func progress() {
	time.Sleep(200 * time.Millisecond)
	fmt.Printf(".")

}
func worker(done chan struct{}) {
	fmt.Printf("working")
	for i := 0; i < 5; i++ {
		progress()
	}
	fmt.Println("done")
	done <- struct{}{}
}

func main() {

	//1. Start Here
	//unbuffered
	// messages := make(chan string, 1)

	// go func() {
	// 	messages <- "ping"
	// }()
	// msg := <-messages
	// fmt.Println(msg)
	// //<-- there's no wait here

	// msg = <-messages //<--because sends and receives block until both the sender and reciever are ready
	// // we need no other synchronization
	// fmt.Println(msg)
	// //2. Then Here
	// //buffered
	// buffered := make(chan string, 2)
	// buffered <- "buffered "
	// buffered <- "channel"

	// fmt.Printf(<-buffered)
	// fmt.Printf(<-buffered)
	// fmt.Println("")

	//3. Then Here
	//synchronization
	done := make(chan struct{}, 1)
	go worker(done)
	<-done //<-- ask the class how removing this (<-done) will affect how the program would run
}
