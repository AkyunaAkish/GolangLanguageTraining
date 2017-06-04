package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 25
	fmt.Printf("%T %v \n", yourAge, yourAge)

	// this doesn't work:
	// fmt.Println(myAge + myAge)

	// this conversion works:
	// fmt.Println(int(myAge) + yourAge)
}
