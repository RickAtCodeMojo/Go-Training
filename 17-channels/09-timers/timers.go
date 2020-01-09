package main

import "time"
import "fmt"

func showProgress(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("")
			return
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("=")
		}
	}
}
func main() {

	// Timers represent a single event in the future. You
	// tell the timer how long you want to wait, and it
	// provides a channel that will be notified at that
	// time. This timer will wait 2 seconds.
	done := make(chan bool)
	timer1 := time.NewTimer(2 * time.Second)
	go showProgress(done)
	// The `<-timer1.C` blocks on the timer's channel `C`
	// until it sends a value indicating that the timer
	// expired.
	<-timer1.C
	done <- true
	fmt.Println("\nTimer 1 expired")

	// If you just wanted to wait, you could have used
	// `time.Sleep`. One reason a timer may be useful is
	// that you can cancel the timer before it expires.
	// Here's an example of that.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
