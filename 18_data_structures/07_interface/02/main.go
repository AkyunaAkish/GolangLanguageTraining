package main

import (
	"fmt"
	"math"
)

// Square ...
type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

// Circle ...
type Circle struct {
	radius float64
}

func (z Circle) area() float64 {
	return math.Pi * z.radius * z.radius
}

// Shape ...
type Shape interface {
	area() float64
}

// Takes any struct that implements the Shape interface
func info(z Shape) {
	fmt.Println("z...(struct)", z)
	fmt.Println("z area...(area)", z.area())
}

// func info(z string) {
// 	fmt.Println("z...(struct)", z)
// 	fmt.Println("z area...(area)", z.area())
// }

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}

// The power of interfaces
// is polymorphism(in this case: one thing can change/morph many other things)
// The interface allows for a function to be able to take
// Different types of structs as long as they have the
// correct signatures to implement the Shape interface

// When a type implements an interface
// an entire world of functionality
// can be opened up to a value of that interface's type

// When passing an interface to a func
// Only the fields/methods defined in the interface
// declaration are exposed in that method
// NOT the additional fields/methods that
// are on the original type implementing the interface
