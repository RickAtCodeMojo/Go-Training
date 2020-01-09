package main

import (
	"fmt"
)

//A struct can contain a struct
var ellipse Circle
var wheel Circle

type Circle struct {
	Radius int
	Point
}
type Point struct {
	x int
	y int
	z struct {
		a int
		b int
	}
}

func main() {

	ellipse.x = 3
	ellipse.y = 4
	ellipse.Radius = 10
	wheel = Circle{Point: Point{x: 8, y: 9}, Radius: 6}

	fmt.Println(ellipse)
	fmt.Println(wheel)
}
