package main

import "fmt"

func zero(x int) {
	x = 0
	fmt.Println(&x) // different address
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(x)
	fmt.Println(x)
	// x is still 5 because only
	// the value was passed into the function,
	// not the memory address
}
