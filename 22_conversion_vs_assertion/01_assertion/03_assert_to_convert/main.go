package main

import "fmt"

func main() {
	var val interface{} = 7
	fmt.Println(val.(int) + 6) // works
	// fmt.Println(val + 6) // doesnt work because interface needs to be asserted to be converted to an int
	// this is different than type casting like so:
	// int(val) // this would break because you cant convert and interface, only assert
}
