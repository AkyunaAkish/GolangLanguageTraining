// a pipline is a series of stages connected
// by channels, where each stage is a group of
// goroutines running the same function.
// In each stage, the goroutines
// receive values from upstaream via inbound
// channels
// perform some function on that data, usually
// producing new values
// and send values downstream via outbound channels
package main

import "fmt"

func main() {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	// since gen writes 2 values
	// to it's returned channel
	// you need to read from the "out"
	// channel twice
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

// nums is a variadic parameter
// because of the ...
// it can be an unlimited number
// of ints
func gen(nums ...int) chan int {
	out := make(chan int)
	// start the go routine using
	// an anonymouse self executing func
	// and write to the channel then close
	// the channel when you have no more
	// data to range over
	// because the sq func is reading from the
	// channel you will not create a
	// deadlock
	// channels always need to be written
	// to and read from the same amount of times
	go func() {
		for _, n := range nums {
			out <- n
		}
		// because out is being closed
		// this func returnes a receive only channel(<-chan)
		close(out)
	}()
	return out
}

// continues the pipeline
// by reading from the original
// channel then returning a new
// channel continuing the
// pipeline sequence
func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
