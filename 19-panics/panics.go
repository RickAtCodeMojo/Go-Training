package main

import "os"

func main() {

	//use panic when the result is unexpected and continuing would produce unexpected results
	//and you don't know how to gracefully hanlde the problem
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err) //we depend on the file for all futher processing
	}

	//how to generate a panic of your own
	panic("a problem") //of course the panic description should be much better than this

}
