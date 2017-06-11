package main

import "fmt"

// func main
// exits before foo and bar
// run
// which is why theres no output in the terminal
// you need a waitgroup to see all output
func main() {
	go foo()
	go bar()
}

// foo and bar run at the same time
func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Foo: ", i)
	}
}

func bar() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Bar: ", i)
	}
}
