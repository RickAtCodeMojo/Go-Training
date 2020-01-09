package main

import (
	"fmt"
	"runtime"
)

func first() {
	second() //line 9
}

func second() {
	third() //line 13
}

func third() {
	stackSlice := make([]byte, 512)
	s := runtime.Stack(stackSlice, false)
	fmt.Printf("\n%s", stackSlice[0:s])
	for c := 0; c < 6; c++ {
		fmt.Println(runtime.Caller(c))
	}
}

func one(panicked bool) []uintptr {
	return two(panicked) // line 26
}

func two(panicked bool) []uintptr {
	return three(panicked) // line 30
}

func three(panicked bool) []uintptr {
	if panicked {
		panic("three") // line 35
	}
	ret := make([]uintptr, 20)
	return ret[:runtime.Callers(0, ret)] // line 38
}

func callers(pcs []uintptr, panicked bool) {
	m := make(map[string]int, len(pcs)) //map of line numbers mapped to frame names
	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		if frame.Function != "" {
			m[frame.Function] = frame.Line
		}
		if !more {
			break
		}
	}

	var called []string
	for k := range m {
		called = append(called, k)
	}
	fmt.Printf("Functions called:\n")
	for _, fnctn := range called {
		fmt.Printf("%s\n", fnctn)
	}

	//if you want to assert where the calls are you could do something like this:
	var f3Line int
	if panicked {
		f3Line = 35
	} else {
		f3Line = 38
	}
	expected := []struct {
		name string
		line int
	}{
		{"main.one", 26},
		{"main.two", 30},
		{"main.three", f3Line},
	}
	for _, loc := range expected {
		if actual := m[loc.name]; actual != loc.line { //syntactic sugar: can check for assignment success and value comparison at the same time
			fmt.Printf("%s is line %d, want %d\n", loc.name, actual, loc.line)
		}
	}

}

func fakePanic() {
	defer func() { //will be reported as fakePanic.func1
		if r := recover(); r == nil {
			fmt.Println("\ndid not panic")
		}
		pcs := make([]uintptr, 20)
		pcs = pcs[:runtime.Callers(0, pcs)]
		callers(pcs, true) //<-- panic
	}()
	one(true)

}

//https://godoc.org/runtime#NumGoroutine
func main() {

	fmt.Println("\nStack ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	first()

	fmt.Println("\nCallers ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	callers(one(false), false) //gets the names of the stack frames

	fmt.Println("\nCallers (Panic) ~~~~~~~~~~~~~~~~~~~~~~")
	fakePanic()
}
