package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	// this for loop waits
	// for the puller func
	// to return the new cSum channel
	// and then loops over the channel values
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() chan int {
	out := make(chan int)
	// launch go routine(non-blocking)
	go func() {
		for i := 0; i < 10; i++ {
			// this for loop reads from channel c
			out <- i
		}
		close(out)
	}()
	// return immediately after go routine is launched passing back the created channel
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		// this for loop reads from channel c
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
