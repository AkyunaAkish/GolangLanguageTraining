package main

import "fmt"

// Square ...
type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

// Shape ...
type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println("z...(Square struct)", z)
	fmt.Println("z area...(Square's area)", z.area())
}

func main() {
	s := Square{10}
	info(s)
}

// An interface defines functionality
// any method that has this signature: area() float64
// implements the Shape interface

// If you(a data type) want to implement the Shape interface
// you need to have a method with this signature(in this example):
// area() float64

// Because struct Square has that method signature
// it implements interface Shape
// The info() method takes as a parameter
// anything that implements interface Shape
