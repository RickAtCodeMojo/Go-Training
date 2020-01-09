package main

import (
	"fmt"
)

func main() {
	//capacity is how much memory is reserved without having to reserve more - which would cost a call to get the memory
	// colours := make([]string, 3)
	// fmt.Println("Colours - Length:", len(colours), ", Capacity:", cap(colours))
	// primeColours := []string{"red", "green", "blue"}
	// secondaryColours := []string{"yellow", "cyan", "fuschia"}
	// bw := []string{"white", "black"}

	// // append two slices
	// colours = append(primeColours, secondaryColours...)
	// fmt.Println(colours)
	// fmt.Println("Colours - Length:", len(colours), ", Capacity:", cap(colours))

	// // append two slices
	// colours = append(colours, bw...)
	// fmt.Println(colours)
	// fmt.Println("Colours - Length:", len(colours), ", Capacity:", cap(colours))

	// //append individual items
	// colours = append(colours, "pink", "brown", "gray")
	// fmt.Println(colours)
	// fmt.Println("Colours - Length:", len(colours), ", Capacity:", cap(colours))

	// var numbers []int
	// for n := 0; n < 5000; n++ {
	// 	numbers = append(numbers, n)
	// 	if n%8 == 0 {
	// 		fmt.Println("Length:", len(numbers), ", Capacity:", cap(numbers))
	// 	}
	// }
	var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println(len(days), cap(days), days)
	days = append(days, "Saturday", "Sunday")
	week := days[:5]
	weekend := days[5:]
	fmt.Println(week)
	fmt.Println(weekend)
	days[0] = "Whaaaat"
	fmt.Println(week)
}
