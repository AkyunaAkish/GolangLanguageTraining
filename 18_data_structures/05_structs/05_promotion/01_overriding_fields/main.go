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
	First         string // this field overrides the "First" field in person
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   42,
		},
		First:         "YourNameHere",
		LicenseToKill: true,
	}

	// p1.First gives the overridden field value
	// though p1.Person.First is still accessible
	fmt.Println(p1.First, p1.Person.First)
}
