package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Primitive unprotected variable

}

type Watchdog struct{ last int64 }

//KeepAlive writes to a Watchdog struct
// Data races can happen on variables of primitive types as well (bool, int, int64, etc.), as in this example:
func (w *Watchdog) KeepAlive() {
	w.last = time.Now().UnixNano() // First conflicting access.
}

//Start reads the watchdog
func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Second conflicting access.
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}

// Even such "innocent" data races can lead to hard-to-debug problems caused by non-atomicity of the memory accesses, interference with compiler optimizations, or reordering issues accessing processor memory .

// A typical fix for this race is to use a channel or a mutex. To preserve the lock-free behavior, one can also use the sync/atomic package.
/*
type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	atomic.StoreInt64(&w.last, time.Now().UnixNano())
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}
*/
