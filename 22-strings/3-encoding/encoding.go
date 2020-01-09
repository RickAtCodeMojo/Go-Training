package main

import (
	"fmt"
)

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {
	bytes := []byte(sample)
	fmt.Println("a slice of bytes")
	fmt.Println(bytes)
	fmt.Println("formatted hex")
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%x ", bytes[i])
	}
	fmt.Println("")
	// fmt.Printf("%x\n", bytes)
	// fmt.Printf("% x\n", bytes)
	fmt.Println("formatted q")
	fmt.Printf("%q\n", bytes)
	fmt.Println("formatted +q")
	fmt.Printf("%+q\n", bytes)
	fmt.Println("characters")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q ", sample[i])
	}
	fmt.Println()
}
