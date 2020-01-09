package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

var flg = flag.Int("n", 1e5, "Number of goroutines to create")

var ch = make(chan byte)
var counter = 0

func f() {
	counter++
	<-ch // Block this goroutine
}

func main() {
	flag.Parse()
	if *flg <= 0 {
		fmt.Fprintf(os.Stderr, "invalid number of goroutines")
		os.Exit(1)
	}
	fmt.Printf("%d processors available\n", runtime.NumCPU())
	// Limit the number of spare OS threads to just 1
	runtime.GOMAXPROCS(1)

	// Make a copy of MemStats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	started := time.Now().UnixNano()
	for i := 0; i < *flg; i++ {
		go f()
	}
	runtime.Gosched()
	ended := time.Now().UnixNano()
	runtime.GC() //force a garbage collection

	// Make a copy of MemStats
	var copiedMemStats runtime.MemStats
	runtime.ReadMemStats(&copiedMemStats)

	if counter != *flg {
		fmt.Fprintf(os.Stderr, "failed to begin execution of all goroutines")
		os.Exit(1)
	}

	fmt.Printf("Number of goroutines: %d\n", *flg)
	fmt.Printf("Per goroutine:\n")
	fmt.Printf("  Memory: %.2f bytes\n", float64(copiedMemStats.Sys-memStats.Sys)/float64(*flg))
	fmt.Printf("  Time:   %f µs\n", float64(ended-started)/float64(*flg)/1e3)
}
