package main

import (
	"fmt"
)

type Body struct {
	Health int
}

func (b Body) GetHealth() {
	fmt.Println(b.Health)
}

func (b Body) Jump() {
	fmt.Println("I'm Jumping")
}

type Mind struct {
	Feeling string
}

func (m Mind) GetHealth() {
	fmt.Println(m.Feeling)
}

func (b Body) Think() {
	fmt.Println("I'm Thinking")
}

type Human struct {
	Name string
	Mind
	Body
}

func (h Human) GetHealth() {
	fmt.Println("I'm feeling human")
}

func main() {
	var x Human
	x.Name = "Brodie"
	x.Feeling = "Feeling good"
	x.Health = 10

	// Human is composed of both body and mind, and has the behaviour of both
	x.Jump()
	x.Think()

	// The child structs can still be accessed directly.
	x.Mind.GetHealth()
	x.Body.GetHealth()
	x.GetHealth()

}
