package main

import "fmt"

func main() {
	// single quotes give you the ASCII
	// number for a character(a rune/int32)
	// "rune" is an alis for int32
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)
}
