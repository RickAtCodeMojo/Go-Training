package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	// We'll use an unsigned integer to represent our
	// (always-positive) counter.
	var counter uint64
	goRoutines := 50
	// To simulate concurrent updates, we'll start 50
	// goroutines that each increment the counter about
	// once a millisecond.
	for i := 0; i < goRoutines; i++ {
		go func() {
			for {
				// To atomically increment the counter we
				// use `AddUint64`, giving it the memory
				// address of our `counter` counter with the
				// `&` syntax.
				atomic.AddUint64(&counter, 1) //<-no surprise we need to pass the variable by reference
				// Wait a bit between increments.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Wait a second to allow some counter to accumulate.
	time.Sleep(time.Second)

	// In order to safely use the counter while it's still
	// being updated by other goroutines, we extract a
	// copy of the current value into `counterFinal` via
	// `LoadUint64`. As above we need to give this
	// function the memory address `&counter` from which to
	// fetch the value.
	counterFinal := atomic.LoadUint64(&counter) //use a copy
	fmt.Println("counter:", counterFinal)
}
