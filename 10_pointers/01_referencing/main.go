package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	var b = &a
	// *int means a memory address pointer to an int
	// var b *int = &a

	fmt.Println(b)
}
