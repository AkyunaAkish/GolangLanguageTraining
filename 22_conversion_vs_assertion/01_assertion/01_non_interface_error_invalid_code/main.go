package main

import "fmt"

// this code fails because assertion is only for interfaces
// not types
func main() {
	name := "Sydney"
	str, ok := name.(string) // error invalid type assertion non-interface type

	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
