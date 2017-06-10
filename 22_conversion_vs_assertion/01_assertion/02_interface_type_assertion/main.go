package main

import "fmt"

// this works because name is an interface
// and assertions are for interfaces
func main() {
	var name interface{} = "Sydney"
	str, ok := name.(string)

	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
