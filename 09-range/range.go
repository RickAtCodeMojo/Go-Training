package main

import "fmt"

func main() {

	//an intereting slice of numbers
	fibs := []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	sum := 0
	for i, f := range fibs { //ignore the index
		sum += f
		fmt.Println(i, f)
	}
	fmt.Println("sum:", sum)

	for i, f := range fibs { //refer to the index, and the value returned by range
		if f == 11 {
			fmt.Println("11 is at index", i)
		}
	}

	// range on a map iterates over key/value pairs.
	teams := map[string]int{"Mercedes": 1, "Ferrari": 2, "Red Bull": 3, "Force India": 4, "Renault": 5, "Williams": 6, "Haas": 7}

	for k, t := range teams {
		fmt.Printf("%s -> %d\n", k, t)
	}

	// range can also iterate over just the keys of a map.
	for k := range teams {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the rune and the second the rune itself.
	fmt.Println("winds 🀀🀁🀂🀃")
	for i, c := range "winds 🀀🀁🀂🀃" {
		fmt.Println(i, c)
	}
}
