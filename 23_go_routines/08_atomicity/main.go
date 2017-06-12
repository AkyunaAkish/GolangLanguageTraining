package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// atomicity is another way to sync processes
// similar to mutex where is can lock down a value
// the main difference is that it doesn't lock down a
// whole chunk of code - it just locks down a single value
// meaning only one process can access a particular value at a time
// even during concurrency and running parallel code

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementor("Foo: ")
	go incrementor("Bar: ")
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		// with race condition:
		// counter++
		// Without race condition:
		atomic.AddInt64(&counter, 1) // adds to counter but locks down value for no race conditions with other processes
		fmt.Println(s, i, "Counter: ", counter)
	}
	wg.Done()
}

// go run -race main.go
// or
// go run main.go
