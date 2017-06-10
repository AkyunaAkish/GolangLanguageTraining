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

// (c *circle) pointer receiver
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	// pass in pointer to circle instead of value
	info(&c)

	// this would break because the receiver for area is a pointer receiver
	// info(c)

	// pointer receivers are given a more rigid type which the Golang compiler
	// will not change - therefore a pointer receiver has to receive a pointer type
	// unlike a value receiver which is given an UNTYPED constant until the Golang
	// compiler gives the constant context/a Type

	// value receivers are more flexible - they can take value or pointer values
}
