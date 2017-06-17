package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
	// alternatively you can put func return directly in the range
	// rather then setting it to a variable first
	// for n := range puller(c) {
	// 	fmt.Println(n)
	// }
}

// chan<- would mean the channel can only send
// which is strange because then nothing would be able
// to read from it after adding values...

// if direction is not specified the channel is bi-directional

// <-chan that can only be used to receive ints
// meaning the channel that is returned from this func
// can't be written to after the return - only read from
func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// takes and returns channels that can only be received/read from not written to
func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
