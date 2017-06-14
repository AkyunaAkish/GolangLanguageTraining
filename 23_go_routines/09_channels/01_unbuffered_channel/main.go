package main

import (
	"fmt"
	"time"
)

// channels are a way to communicate data
// between parallel running functions
// when set set a value on a channel
// that channel will wait until the value
// is read by another function
// before moving onto setting the next value
// channels have a "blocking" nature
func main() {
	// create unbuffered channel that can communicate an int
	c := make(chan int)

	// buffered channel example
	// can hold 10 things
	// more advanced technique
	// which is best to avoid
	// unless you have a use case
	// c := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			// set value of channel to equal i
			c <- i
		}
	}()

	go func() {
		for {
			// read and print current value of channel
			fmt.Println(<-c)
		}
	}()
	// sleep at the end for a second
	// is needed in order to see the
	// values printed
	// the functions run too fast
	// to see the output without sleeping
	// I don't know if the sleeping is neccessary
	// to actually run the functions or if it is
	// only needed to see the output
	time.Sleep(time.Second)
}
