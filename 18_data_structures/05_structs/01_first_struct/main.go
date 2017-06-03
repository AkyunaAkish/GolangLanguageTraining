package main

import "fmt"

// in Golang, instead of creating a class that you can
// then go on to instantiate new instances of
// you create a struct type
// and then assign variables that type
type person struct {
	first string // field
	last  string // field
	age   int    // field
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Alice", "Wonderland", 14}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}

// A struct is how you can create a reusable constructor object
// (using javascript terms) AKA class
