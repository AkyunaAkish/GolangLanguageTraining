package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	// adding to the wait group here is correct
	// so that is doesn't happen within concurrent
	// functions which would cause a race condition
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// remove one from wait group
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		// after wait group is finished
		// close the channel
		close(c)
	}()

	for n := range c {
		// as funcs above write values to the channel
		// this for loop will read the value
		// step by step until the channel is closed
		fmt.Println(n)
	}
}
