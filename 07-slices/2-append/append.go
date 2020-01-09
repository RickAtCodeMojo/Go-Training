package main

import (
	"fmt"
)

func main() {

	primeColours := []string{"red", "green", "blue"}
	secondaryColours := []string{"yellow", "cyan", "fuschia"}
	// bw := []string{"white", "black"}
	var colours []string
	colours = append(primeColours, secondaryColours...)
	for i, _ := range colours {
		// fmt.Println(&c)
		fmt.Println(&colours[i])
	}
	colours = append(colours, "pink")
	for i, _ := range colours {
		// fmt.Println(&c)
		fmt.Println(&colours[i])
	}

	//append is variadic
	colours = append(colours, "pink", "brown", "gray")
	fmt.Println(colours)

}
