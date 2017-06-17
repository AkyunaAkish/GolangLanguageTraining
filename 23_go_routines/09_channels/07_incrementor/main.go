package main

import "fmt"

func main() {
	c1 := incrementor("Foo: ")
	c2 := incrementor("Bar: ")
	// passing created channels
	// into reader/receiver function
	c3 := puller(c1)
	c4 := puller(c2)

	// this statement:  <-c3+<-c4
	// blocks the main function from
	// exiting
	// until those channels are readable
	// if that wasn't there the main
	// function would exit right away
	// <- that notation just reads
	// a value and removes that value
	// from the channel
	fmt.Println("Final Counter: ", <-c3+<-c4)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			// sending data to the channel
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		// receiving data from the channel
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
