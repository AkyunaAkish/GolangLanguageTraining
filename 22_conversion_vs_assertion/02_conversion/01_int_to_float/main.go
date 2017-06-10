package main

import "fmt"

func main() {
	var x = 12
	var y = 12.123456
	fmt.Println(y + float64(x))
	// conversion of int to float64
	// this conversion is widening not narrowing
	// beacause now the type int can hold more value
}
