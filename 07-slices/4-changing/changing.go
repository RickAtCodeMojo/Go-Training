package main

import (
	"fmt"
	"strings"
)

func main() {

	colours := []string{"red", "green", "blue", "yellow", "cyan", "fuschia", "white", "black"}
	fmt.Println(colours)

	lower := 2
	upper := 5
	portion := colours[lower:upper]
	fmt.Println("elements between", upper, "and", lower, ":", portion)

	// This slices up to but excluding upper.
	portion = colours[:upper]
	fmt.Println("elements up to", upper, ":", portion)

	// And this slices up from and including lower.
	portion = colours[lower:]
	fmt.Println("elements from:", lower, "to last:", portion)

	//the last element
	last := colours[len(colours)-1]
	fmt.Println("Last", last)

	//drop the last element
	colours = colours[:len(colours)-1]
	fmt.Println(colours)

	//pop last
	colours, last = colours[:len(colours)-1], colours[len(colours)-1]
	fmt.Println(last, colours)

	//return last and pop last
	last, colours = colours[len(colours)-1], colours[:len(colours)-1]
	fmt.Println(last, colours)

	//first
	first := colours[0]

	//pop front
	colours = colours[1:]
	fmt.Println(first, colours)

	//first & pop front
	first, colours = colours[0], colours[1:]
	fmt.Println(first, colours)

	//Join a slice of strings with a string intersperces the strings with the string
	fmt.Println(strings.Join(colours, " "))

	fmt.Println("Index Exchange"[6:])
	name := "Index Exchange"
	name = "An Ad Serving " + name[6:]
	fmt.Println(name)
}
