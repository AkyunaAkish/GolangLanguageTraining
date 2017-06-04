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

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   42,
		},
		LicenseToKill: true,
	}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p1.Person.First) // alternate syntax
}
