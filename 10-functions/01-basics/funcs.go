package main

import (
	"fmt"
	"math"
)

type result struct {
	value float64
}

func main() {

	// var x uint64
	// x = math.MaxUint64
	// a := uint8(x)

	// fmt.Println(a)
	// // print(math.Pi)
	// z := getZ(7.5, 2.8)

	// fmt.Println(z)
	res := result{value: 9999}
	fmt.Println(res)
	convert(res)
	fmt.Println(res.value)
	convPtr(&res)
	fmt.Println(res.value)
	// fmt.Println(nan(z))
}

func print(number float64) {
	fmt.Printf("%.2f\n", number)
}

func getZ(x float64, y float64) float64 {
	return x + math.Sin(2.0*math.Pi/y)
}

func convert(x result) {
	z := 1.0 / math.Sqrt(x.value)
	x.value = z
}
func convPtr(x *result) {
	z := 1.0 / math.Sqrt(x.value)
	x.value = z

}
func nan(x float64) float64 {
	return math.Sqrt(1.0 - math.Pow(x, 2.0))
}
