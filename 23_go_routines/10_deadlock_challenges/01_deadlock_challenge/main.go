package main

import "fmt"

// output: fatal error: all goroutines are asleep - deadlock!
// func main() {
// 	c := make(chan int)
// this puts a value on the channel
// though there is nothing to
// receive it
// 	c <- 1
// 	fmt.Println(<-c)
// }

// this results in a deadlock
// can you determine why?
// and what would you do to fix it?

// my solution
func main() {
	c := make(chan int)
	// a go routine is needed
	// to set values to a channel
	// without blocking a function
	go func() {
		c <- 1
		close(c)
	}()
	// the go routine allows
	// this receiver
	// to properly wait
	// and block the main func
	// to receive
	fmt.Println(<-c)
}

// a deadlock is when there is a "miss"
// between a send and receive in a channel
// the data wasn't properly delivered
