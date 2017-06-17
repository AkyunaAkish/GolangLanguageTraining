package main

import "fmt"

// func main() {
// 	f := factorial(4)
// 	fmt.Println("Total: ", f)
// }

// func factorial(n int) int {
// 	total := 1
// 	for i := n; i > 0; i-- {
// 		total *= i
// 	}
// 	return total
// }

// challenge: use go routines to solve this
// factorial of 4 = 4 * 3 * 2 * 1

func main() {
	f := factorial(4)
	fmt.Println("Total: ", <-f)
	// OR
	// for n := range c {
	// 	fmt.Println(n)
	// }
}

func factorial(n int) chan int {
	c := make(chan int)

	go func() {
		total := 1

		for i := n; i > 0; i-- {
			total *= i
		}

		c <- total
		close(c)
	}()

	return c
}

// why would you want to use a channel in this instance?
// if you needed to run many factorial calculations in short period of time
// using go routines/channels would allow all of those calculations to be
// ran in parallel using all of the CPU core power of your machine
// and therefore complete faster
