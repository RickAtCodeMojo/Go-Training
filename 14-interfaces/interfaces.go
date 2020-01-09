package main

import (
	"fmt"
	"math"
)

type areaOnly interface {
	area() float64
}
type geometry interface {
	area() float64
	perimeter() float64
}
type triangle struct {
	a float64
	b float64
	c float64
}
type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}
type ellipse struct {
	a     float64
	b     float64
	color uint
}

func (t triangle) area() float64 {
	var s float64
	var Area float64
	fmt.Println("Enter three sides of triangle : ")

	s = (t.a + t.b + t.c) / 2.0
	Area = math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))

	return Area
}
func (t triangle) perimeter() float64 {
	return t.a + t.b + t.c
}
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (e ellipse) area() float64 {
	return math.Pi * e.a * e.b
}

//ellipse perimiter too hard to calculate...

func main() {

	t := triangle{a: 5.6, b: 8.1, c: 10.3}
	// fmt.Printf("The area of a triangle with sides %g, %g, %g is %g\n", t.a, t.b, t.c, t.area())
	// c := circle{radius: 4.5}
	// fmt.Printf("The area of a circle with a radius of %g is %g\n", c.radius, c.area())
	// r := rectangle{width: 7.5, height: 5.4}
	// measure(r)
	// measure(c)
	// What happens if you try to call measure on an interface that is not fully implemented
	// e := ellipse{a: 4.5, b: 8.4}
	// measure(e)
	check(t)
	check("hello")
}
func check(i interface{}) {
	switch i.(type) { //use a type switch to see the underlying type
	case string:
		fmt.Println(i)
	default:
		fmt.Println("I only work on strings")
	}

}
func area(a areaOnly) {

}
func measure(g geometry) {
	switch g.(type) { //use a type switch to see the underlying type
	case rectangle:
		fmt.Println("Sides:", g)
		r, ok := g.(rectangle) //use a type assertion to get the rectangle from the interface
		if ok {
			fmt.Printf("The area of the rectangle of width %g and height %g is %g\n", r.width, r.height, g.area())
			fmt.Printf("The perimeter of the same rectangle is %g\n", g.perimeter())
		}
	case circle:
		fmt.Println("Radius:", g)
		fmt.Printf("The area of the corcle is %g\n", g.area())
		fmt.Printf("The perimeter of the circle is %g\n", g.perimeter())
	default:
		fmt.Printf("The area is %g\n", g.area())
		fmt.Printf("The perimeter is %g\n", g.perimeter())
	}
}
