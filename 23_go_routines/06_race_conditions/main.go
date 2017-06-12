package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

// because this counter is global
// multiple go routines will be modifying
// and reading from the same value
// which will cause race conditions
// meaning that the value will not be
// what is expected due to outer influences
// on your value that shouldn't really be there
var counter int

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		// this resetting of counter will cause problems
		// with other functions running concurrently/in parallel
		// due to a dependency on the global variable
		counter = x
		fmt.Println(3, i, "Counter: ", counter)
		// you can run this
		// file to check if a race condition exists
		// $ go run -race main.go (one or two dashes for race flag)
	}
	wg.Done()
}
