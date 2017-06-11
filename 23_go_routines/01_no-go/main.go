package main

import "fmt"

func main() {
	foo()
	bar()
}

// this runs first
func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Foo: ", i)
	}
}

// this runs second
func bar() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Bar: ", i)
	}
}
