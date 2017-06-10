package main

import "fmt"

// reminder a rune is an alias to int32
func main() {
	var x = 'a'
	// var x rune = 'a'
	var y int32 = 'b'
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(string(x))
	fmt.Println(string(y))
	// conversion rune to string
}
