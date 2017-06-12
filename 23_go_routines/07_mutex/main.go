package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// mutex means mutual exclusion object
var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo: ")
	go incrementor("Bar: ")
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		// mutex.Lock() will lock down some executed code
		// and pause other threads that may be trying to access
		// values within the locked code so that there will be no
		// conflicts with values
		// this is handling race conditions
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter: ", counter)
		mutex.Unlock()
	}
	wg.Done()
}
