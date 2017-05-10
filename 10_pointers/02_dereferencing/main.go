package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)
	//*b means dereference the memory address
	//and give the value instead
	//* is an operator that give a value instead of an address
	fmt.Println(*b)
}
