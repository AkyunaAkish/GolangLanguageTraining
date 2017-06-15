package main

import "fmt"

// n-times-to-1
// means n number of
// functions
// writing to one channel

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		// creates go routine
		// 9 times where each
		// routine
		// adds a value to the
		// "c" channel 9 times
		// and each routine
		// adds a boolean
		// "true" to the "done"
		// channel once
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		// waits for value on done
		// and removes a value
		// from done when one is a
		// added 9 times
		// and then closes the channel
		// "c"
		for i := 0; i < n; i++ {
			// this statement
			// is blocking
			// so the channel
			// "c" will not be
			// closed until
			// all values of the
			// "done" channel have
			// been removed
			<-done
		}
		close(c)
	}()

	// read values from the "c" channel
	// as they're added
	for n := range c {
		fmt.Println(n)
	}
}
