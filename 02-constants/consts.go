// Go supports _constants_ of character, string, boolean,
// and numeric values.

package main

import "fmt"
import "math"

type suit uint

// `const` declares a constant value.
const name string = "Index Exchange"

const (
	clubs suit = iota
	diamonds
	hearts
	spades
)

func main() {
	fmt.Println(name)
	const val = 500000000 //consts can be defined anywhere that a var can

	const small = math.MaxFloat32 //MaxFloat32 is also a constant

	const large = 3e20 / small //consts can be the result of an expression

	fmt.Println(large)

	const aFloat = 3e20 / val //arbitrary precision
	fmt.Println(aFloat)

	fmt.Println("math.Sin(", val, ") = ", math.Sin(val)) //Sin expects float64
	start := 0
	increment := 0
	increment++
	start = increment
	fmt.Println(start)
	fmt.Println(increment)
	for c := clubs; c <= spades; c++ {
		switch c {
		case clubs:
			fmt.Println("Clubs")
		default:
			fmt.Println(c)
		}
	}
}
