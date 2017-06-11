package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// init is a special function
// that can be used to configure/initialze
// a program
func init() {
	// this will use ALL CPUs
	// before Go 1.5 concurrency
	// would only run on one core unless you told
	// it too
	// now after 1.5 the default uses more than one core(not sure how many - maybe all)
	// if you want to make sure you are using all cores
	// you can do this:
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 450000; i++ {
		fmt.Println("Foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 450000; i++ {
		fmt.Println("Bar: ", i)
	}
	wg.Done()
}
