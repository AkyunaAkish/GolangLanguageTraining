package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// add values to channel
			c <- i
		}
		// close channel/stop adding values
		// values can still be read from the channel
		// but you cannot add more to the channel
		// close will not run until all of the
		// values have been printed
		// in the range for loop
		close(c)
	}()

	// read channel in a range
	// every time a channel is read from
	// the value that was read is removed
	// from the channel
	// therefore the for loop will end
	// because every loop over the range
	// removes an item from the channel
	// and will eventually be empty and end the loop
	for n := range c {
		// time.Sleep was not needed here because
		// this forloop blocked the function execution
		// until the channel has finished being read from
		fmt.Println(n)
	}
}
