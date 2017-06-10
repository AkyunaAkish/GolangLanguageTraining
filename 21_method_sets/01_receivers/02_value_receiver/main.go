package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// (c circle) value receiver
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	// pass in value c not pointer to c
	info(c)

	// passing a pointer value would also work the same
	// info(&c)
	// the golang compiler will translate it to a value for us in the area method
	// this is made possible by UNTYPED constants and their flexibility in Golang
	// if something is still in the UNTYPED constant state - once it's used in the program
	// Golang figures out how it's being used and applies the correct type
}
