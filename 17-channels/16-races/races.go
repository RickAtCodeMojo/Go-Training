package main

//https://golang.org/doc/articles/race_detector.html
import (
	"fmt"
)

func main() {
	done := make(chan bool)
	creatives := make(map[string]string)
	go func() {
		creatives["Grand View Motors"] = "Spring Inventory Sale"
		done <- true
	}()
	creatives["High Park Chrysler"] = "New Arrivals"
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

SOLUTION in races-solution

The cost of race detection varies by program, but for a typical program,
memory usage may increase by 5-10x and execution time by 2-20x.
*/
