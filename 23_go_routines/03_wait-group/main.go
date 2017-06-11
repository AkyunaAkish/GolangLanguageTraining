package main

import (
	"fmt"
	"sync"
)

// using waitgroup allows you to
// make a function wait for it's
// go routines to finish
// before exiting
var wg sync.WaitGroup

func main() {
	// add 2 go routines
	// to the wait group
	wg.Add(2)
	go foo()
	go bar()
	// when wait group becomes 0
	// the wait will finish
	wg.Wait()
	fmt.Println("Done with wait group")
}

func foo() {
	for i := 0; i < 1000000; i++ {
		fmt.Println("Foo: ", i)
	}
	// remove 1 go routine from wait group
	wg.Done()
}

func bar() {
	for i := 0; i < 1000000; i++ {
		fmt.Println("Bar: ", i)
	}
	// remove 1 go routine from wait group
	wg.Done()
}
