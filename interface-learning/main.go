package main

import "fmt"

//An interface is a set of method signatures.

// “If something can do these things, I don’t care what it is.”

type Rectangle struct {
	lenght float64
	width  float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.lenght * r.width
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var s Shape
	s = Rectangle{10, 20}
	fmt.Println(s.Area())

	s = Circle{10}
	fmt.Println(s.Area())
}
