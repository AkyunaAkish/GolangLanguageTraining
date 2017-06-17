package main

import "fmt"

// func main() {
// 	c := make(chan int)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c <- i
// 		}
// the channel does not get closed properly either
// 	}()
// this only pulls one value from c
// 	fmt.Println(<-c)
// }

// Why does this only print zero?
// and what can you do to get it to print
// all 0 - 9 numbers?

// my solution
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// close channel
		close(c)
	}()
	// read all values from channel
	for n := range c {
		fmt.Println(n)
	}
}
