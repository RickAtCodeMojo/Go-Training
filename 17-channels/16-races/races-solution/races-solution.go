package main

//https://golang.org/doc/articles/race_detector.html
import (
	"fmt"
)

func main() {
	done := make(chan bool)
	creatives := make(map[string]string)
	// var mtx = &sync.Mutex{}
	go func() {
		// mtx.Lock()
		creatives["Grand View Motors"] = "Spring Inventory Sale"
		// mtx.Unlock()
		done <- true
	}()
	// mtx.Lock()
	creatives["High Park Chrysler"] = "New Arrivals"
	// mtx.Unlock()
	<-done
	for k, v := range creatives {
		fmt.Println(k, v)
	}
}

/*
$ go test -race mypkg    // to test the package
$ go run -race mysrc.go  // to run the source file
$ go build -race mycmd   // to build the command
$ go install -race mypkg // to install the package
*/

/*
ask the class how they would fix these race conditions
reminder to run the binary under load - this is the best way to be sure

The cost of race detection varies by program, but for a typical program,
memory usage may increase by 5-10x and execution time by 2-20x.
*/
