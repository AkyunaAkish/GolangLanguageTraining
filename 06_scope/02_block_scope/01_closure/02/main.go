package main

import "fmt"

func main() {
	x := 0
	//anonymous function only "enclosed"/visible
	//in this "main" function
	//FYI this is the only way to make a function
	//within another function
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
