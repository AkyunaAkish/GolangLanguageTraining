package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// func (receiver) funcName params() return funcBody{}
func (p person) fullName() string {
	return p.first + " " + p.last
}

// the receiver of a function attaches that function to the
// type that is passed into the receiver

func main() {
	p1 := person{"James", "Bond", 42}
	p2 := person{"Benjamin", "Franklin", 311}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
