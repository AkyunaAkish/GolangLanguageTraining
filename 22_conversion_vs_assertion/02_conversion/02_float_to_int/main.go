package main

import "fmt"

func main() {
	var x = 12
	var y = 12.123456
	fmt.Println(int(y) + x)
	// conversion of float64 to int
	// this conversion is narrowing not widening
	// beacause now the type float64 holds less values
}
