package main

import "fmt"
import "os"

// Suppose we wanted to create a file, write to it,
// and then close when we're done. Here's how we could
// do that with `defer`.
func main() {

	file := create("defer.txt")

	defer close(file) //make sure that what ever else happens, the file is closed
	defer Showme()

	writeTo(file)
}

func create(p string) *os.File {
	fmt.Println("Creating", p)
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func Showme() {
	fmt.Println("Showing stuff ")
}
func writeTo(f *os.File) {
	fmt.Println("Writing to", f.Name())
	for i := 0; i < 25; i++ {
		fmt.Fprintln(f, fmt.Sprintf("%d data", i))
	}

}

func close(f *os.File) {
	fmt.Println("Closing", f.Name())
	f.Close()
}
