package main

import "fmt"

// this is especially useful
// if you want to attach a method
// to a data type
// but don't want to modify
// the standard data type
type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)
}
