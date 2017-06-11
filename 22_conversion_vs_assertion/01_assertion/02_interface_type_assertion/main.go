package main

import "fmt"

// this works because name is an interface
// and assertions are for interfaces
func main() {
	var name interface{} = "Sydney"
	// var name interface{} = 12 // this would be false because the assertion would said that the value is not a strint
	str, ok := name.(string)

	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
