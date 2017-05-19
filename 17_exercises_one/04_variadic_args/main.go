package main

import "fmt"

// write a func that can be called in all of these ways:
func main() {
	fmt.Println(foo(1, 2))
	fmt.Println(foo(1, 2, 3))
	aSlice := []int{1, 2, 3}
	fmt.Println(foo(aSlice...))
	fmt.Println(foo())
}

func foo(p ...int) []int {
	return p
}
