package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	//command line
	if len(os.Args) < 2 {
		fmt.Println("No Args")
		return
	}
	commandLine := os.Args
	args := os.Args[1:]

	// You can get individual args with normal indexing.
	arg := os.Args[1]
	fmt.Println("Args")
	fmt.Println(commandLine)
	fmt.Println(args)
	fmt.Println(arg)
	//go run main.go a b c d

	boolean := flag.Bool("b", false, "bool value")
	number := flag.Int("n", 0, "integer value")
	float := flag.Float64("f", 0.0, "float value")
	var str string
	flag.StringVar(&str, "s", "unknown", "string value")

	flag.Parse()

	fmt.Println("number:", *number)
	fmt.Println("flag:", *boolean)
	fmt.Println("string:", str)
	fmt.Println("float:", *float)
	fmt.Println("remaining:", flag.Args())
	//go run main.go -n=300 -s=something -b=true -f=90.50
}
