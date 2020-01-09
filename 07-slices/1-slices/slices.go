//Slices
//https://github.com/golang/go/wiki/SliceTricks
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//define a slice with var...
	var seasons []string
	seasons = []string{"spring", "summer", "winter"}
	fmt.Println("length of seasons:", len(seasons), "capacity of seasons:", cap(seasons))
	seasons = append(seasons, "fall")
	fmt.Println("length of seasons:", len(seasons), "capacity of seasons:", cap(seasons))

	//define a slice with make...
	colours := make([]string, 3, 5)
	// fmt.Println("empty", colours)

	// We can set and get just like with arrays.
	colours[0] = "red"
	colours[1] = "green"
	colours[2] = "blue"
	fmt.Println("entire slice:", colours)
	fmt.Println("element 2:", colours[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(colours))

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin append, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.
	colours = append(colours, "yellow")
	colours = append(colours, "cyan", "fuschia")
	fmt.Println("appended colours:", colours)

	hues := make([]string, len(colours))
	copy(hues, colours) //copies one slice into another
	fmt.Println("copied colours:", hues)

	lower := 2
	upper := 5
	portion := colours[lower:upper]
	fmt.Println("elements between", upper, "and", lower, ":", portion)

	// This slices up to (but excluding) `s[5]`.
	portion = colours[:upper]
	fmt.Println("elements up to", upper, ":", portion)

	// And this slices up from (and including) `s[2]`.
	portion = colours[lower:]
	fmt.Println("elements from:", lower, "to last:", portion)

	// We can declare and initialize a variable for slice
	// in a single line as well.
	nouns := []string{"book", "train", "elbow", "scroll", "face", "space", "neutral", "basket"}
	plurals := nouns
	print(plurals)
	plural(nouns)
	print(nouns)
	print(plurals)
	plurals[1] = "locomotives"
	print(nouns)
	fmt.Println("pop the first two elements from nouns")
	nouns = nouns[2:]
	print(nouns)
	print(plurals)
	words := [4]string{"pillow", "insect", "pot", "board"}
	plural(words[:])
	// plural(words)
	print(words[:])

	rand.Seed(time.Now().UTC().UnixNano())
	//2 dimensional arrays can be ragged
	bounds := 5
	RaggedSlice := make([][]int, bounds)
	for i := 0; i < bounds; i++ {
		innerLen := 1 + rand.Intn(bounds)
		RaggedSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			RaggedSlice[i][j] = i + j
		}
	}
	fmt.Println("Ragged: ", RaggedSlice)

	// Sorting the builtin string type
	sort.Strings(colours)
	fmt.Println("Sorted Colours:", colours)

	// Sorting an array means passing in a slice
	sort.Slice(words[:], func(i, j int) bool {
		return words[i] < words[j]
	})
	fmt.Println("Sorted Words:", words)
}

type team struct {
	name     string
	games    uint
	wins     uint
	losses   uint
	otlosses uint
}

func plural(strings []string) {
	for i, s := range strings {
		s += "s"
		strings[i] = s
	}
}
func print(strings []string) {
	for _, s := range strings {
		fmt.Printf("%s ", s)
	}
	fmt.Printf("\n")
}
