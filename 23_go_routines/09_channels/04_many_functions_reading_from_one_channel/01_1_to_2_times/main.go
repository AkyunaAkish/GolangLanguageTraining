package main

import "fmt"

// this method does not make sense for most use cases
// but if you happen to need 2+ funcs reading from one channel
// you can do that - though each channel reading will contain a random
// item from the channel - not sequential data
func main() {
	// create channels
	c := make(chan int)
	done := make(chan bool)
	// launch 3 go routes
	go func() {
		for i := 0; i < 100000; i++ {
			// add values to channel
			c <- i
		}
		close(c)
	}()

	go func() {
		// reads from c channel
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	go func() {
		// reads from c channel
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	// the output of this code doesn't have duplication
	// the reason is the each loop that reads from the c channel
	// simply read one value and remove it from the channel
	// whichever go routine reads a channel value first prints that next value
	// each for loop will not need to read the same value sequentially - just values as they come through the channel

	// waiting for 2 done values to be added to the "done" channel when funcs
	// with for loops in them complete
	// doing this allows the done channel to close and be emptied
	<-done
	<-done
}
