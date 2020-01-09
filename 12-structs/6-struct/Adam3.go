package main

import (
	"fmt"
)

type Legs struct {
}

type Wheels struct {
}

type Moveable interface {
	Move()
}

// both structs now satisfy the interface

func (l Legs) Move() {
	fmt.Println("I'm running slow")
}

func (w Wheels) Move() {
	fmt.Println("I'm rolling fast")
}

//truckasaurus is composed of an interface
type Truckasaurus struct {
	Moveable
}

func main() {
	var t Truckasaurus
	//var l Legs

	//t.Moveable = l //we can dynamically compose a struct, giving it new behaviour on the fly
	t.Move() //but be careful (comment out previous line)

}
