package main

import "fmt"

//the zero function takes a memory address of an integer
func zero(x *int) {
	*x = 0
	fmt.Println(x)
	// same address AND *x = 0
	// affects the original parameter
}

func main() {
	x := 5
	fmt.Println(&x)
	//pass the original memory address
	zero(&x)
	fmt.Println(x) // 0
}
