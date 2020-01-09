package main

func main() {
	// ** nil channel

	// // A channel is in a nil state when it is declared to its zero value
	// var ch chan string

	// // A channel can be placed in a nil state by explicitly setting it to nil.
	// ch = nil //this will make it block forever
	// ch <- "message"

	// ** open channel
	// // A channel is in a open state when it’s made using the built-in function make.
	chA := make(chan string)

	// ** closed channel
	// A channel is in a closed state when it’s closed using the built-in function close.
	close(chA)
	chA <- "message"
}
