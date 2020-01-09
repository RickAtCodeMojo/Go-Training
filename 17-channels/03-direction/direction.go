package main

import "fmt"

func ping(pings chan<- string, msg string) {
	fmt.Println("ping!")
	fmt.Println("putting message into pings")
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	fmt.Println("pong!")
	fmt.Println("getting message from pings")
	msg := <-pings
	fmt.Println("putting message into pongs")
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hello World")
	pong(pings, pongs)
	fmt.Println("printing message from pongs")
	fmt.Println(<-pongs)
}
