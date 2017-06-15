package main

import "fmt"

func main() {
	// create two channels and launch
	// 3 go routines
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// instead of wg.Done() we set the done
		// channel to true
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// instead of wg.Done() we set the done
		// channel to true
		done <- true
	}()

	// the semaphore pattern should be done
	// in a go routine function
	// if the semaphore pattern is used
	// outside of one it will not work properly
	// the function will "hang" because
	// the channels will not close properly
	// due to the for loop reading the "c" channel
	// values being blocked by the <-done statements
	// in the same scope. <-done statments being in
	// a go routine func allows the for loop
	// to read the channel values and keep the
	// cycle going
	go func() {
		// each <-done statement
		// waits for a value to go onto
		// the done channel
		// once a value comes on
		// the statement will take that value
		// off of the channel
		// we write this twice
		// because 2 go routines are
		// going to be placing values onto
		// the done channel
		// the <-done statements block this function's
		// execution which will cause close(c)
		// to only run once the 2 <-done
		// statements have finished running
		<-done
		<-done // waits for a value and takes the value off the channel when the value is added
		close(c)
	}()

	// this for loop is blocking
	// because we're using unbuffered channels
	// this allows us to see the function execution
	// and print values
	// once the "c" channel has been closed
	// this range "drains" the channel
	// of it's REMAINING values and uses them in the for loop
	for n := range c {
		fmt.Println(n)
	}
}

// SEMAPHORE:
// A trivial semaphore is a plain variable that is
// changed(incremented/decremented/toggled) depending
// on a programmer-defined condition/s. The variable
// is then used as a condition to control access to
// some system resource
// similar to a flag  git push origin master --force(force being the flag)

// A semaphore is a way of "purely" using channels
// without relying on package methods such as
// sync.WaitGroup or atomicity to control your
// channel flow
