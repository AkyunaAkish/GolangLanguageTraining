package main

import "fmt"

// Person exported
type Person struct {
	First string
	Last  string
	Age   int
}

// this struct DoubleZero
// is getting "promoted"
// by inheriting the fields in the person struct

// DoubleZero exported
type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) greeting() {
	fmt.Println("Hi there " + p.First)
}

func (dz DoubleZero) greeting() {
	fmt.Println("Hi " + dz.First)
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   42,
		},
		LicenseToKill: true,
	}

	p1.greeting()
	p1.Person.greeting()
}
