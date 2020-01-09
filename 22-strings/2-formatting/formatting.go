// Go offers excellent support for string formatting in
// the `printf` tradition. Here are some examples of
// common string formatting tasks.

package main

import (
	"fmt"
	"os"
)

type position uint

const (
	first position = iota
	second
	third
	fourth
	fifth
	sixth
	seventh
	eigth
	ninth
	tenth
	finished
	dnf
)

//Driver describes a driver and his season
type Driver struct {
	Name        string
	Country     string            `json:"country"` //when encoded the tag will be country not Country
	Constructor string            `json:"constructor"`
	Score       int               `json:"score"`
	Results     map[string]Result `json:"results"`
}

//Duration describes how much time something took
type Duration struct {
	Hours   int
	Minutes int
	Seconds float64
}

//Result struct describes a race result
type Result struct {
	Date     string
	Position position
	Laps     int
	Length   Duration
}

//Finished looks at the position and returns true if the result was not a dnf
func (r Result) Finished() bool {
	return r.Position != dnf
}

func main() {
	results := map[string]Result{
		"China":         Result{Date: "09 Apr 2017", Position: first, Laps: 56, Length: Duration{Hours: 1, Minutes: 37, Seconds: 36.158}},
		"Spain":         Result{Date: "14 May 2017", Position: first, Laps: 66, Length: Duration{Hours: 1, Minutes: 35, Seconds: 56.497}},
		"Canada":        Result{Date: "11 Jun 2017", Position: first, Laps: 70, Length: Duration{Hours: 1, Minutes: 33, Seconds: 05.154}},
		"Great Britain": Result{Date: "16 Jul 2017", Position: first, Laps: 51, Length: Duration{Hours: 1, Minutes: 21, Seconds: 27.430}},
		"Belgium":       Result{Date: "27 Aug 2017", Position: first, Laps: 44, Length: Duration{Hours: 1, Minutes: 24, Seconds: 42.820}},
		"Italy":         Result{Date: "03 Sep 2017", Position: first, Laps: 53, Length: Duration{Hours: 1, Minutes: 15, Seconds: 32.312}},
		"Singapore":     Result{Date: "17 Sep 2017", Position: first, Laps: 58, Length: Duration{Hours: 2, Minutes: 3, Seconds: 23.544}},
		"Japan":         Result{Date: "08 Oct 2017", Position: first, Laps: 53, Length: Duration{Hours: 1, Minutes: 27, Seconds: 31.194}},
		"United States": Result{Date: "22 Oct 2017", Position: first, Laps: 56, Length: Duration{Hours: 1, Minutes: 33, Seconds: 50.991}},
	}
	//print a struct
	k := "Italy"
	r := results[k]
	fmt.Printf("%v\n", r)  //use %v
	fmt.Printf("%+v\n", r) //use %+v to get the names of the struct fields
	fmt.Printf("%#v\n", r) //use %#v to show the source code that 'owns' the struct values
	fmt.Printf("%T\n", r)  //use %T for the type information

	fmt.Printf("%t\n", r.Finished()) //%t for bool

	fmt.Printf("%d\n", r.Laps) //many options for ints, %d most common
	fmt.Printf("%x\n", r.Laps) //%x is hex
	fmt.Printf("%b\n", r.Laps) //%b binary representation

	fmt.Printf("%c\n", 101) //%c the character equivalent

	fmt.Printf("%f\n", r.Length.Seconds) //%f for float
	t := float64(r.Length.Hours*60*60*1000) + float64(r.Length.Minutes*60*1000) + r.Length.Seconds*1000
	fmt.Printf("%e\n", t)
	fmt.Printf("%E\n", t)

	fmt.Printf("%s\n", k)            //%s for strings
	fmt.Printf("%s\n", "\"string\"") //\" for literal quotes
	fmt.Printf("%q\n", "\"string\"") //%q to double quote strings
	fmt.Printf("%q\n", k)            //%q to double quote strings

	fmt.Printf("%x\n", k) //hex with two characters per bytes

	fmt.Printf("%p\n", &r) //%p for a pointer

	sr := results["Singapore"]
	fmt.Printf("%2d:%2d:%08.2f\n", r.Length.Hours, r.Length.Minutes, r.Length.Seconds)  //% then a number 'n' controls the number of spaces or digists
	fmt.Printf("%2d:%2d:%.2f\n", sr.Length.Hours, sr.Length.Minutes, sr.Length.Seconds) //e.g. %2d or %.2f

	fmt.Printf("|%-8.2f|%-8.2f|\n", sr.Length.Seconds, r.Length.Seconds) //Left justify with '-'

	for k, v := range results {
		fmt.Printf("|%-15s on %15s|\n", k, v.Date) //size and alignment with strings
	}

	//use Sprintf to create strings
	var when string
	fmt.Printf("\nNext race:\n")
	when = fmt.Sprintf("|%-15s on %15s|\n", k, r.Date)
	fmt.Println(when)

	// // You can format+print to `io.Writers` other than
	// // `os.Stdout` using `Fprintf`.
	fmt.Fprintf(os.Stderr, "Error occured: %s\n", "some error")
}
