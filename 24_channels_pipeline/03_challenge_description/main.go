package main

import "fmt"

// solution
func main() {
	in := gen()

	f := factorial(in)
	// this range will block
	// main from exiting
	// until all the channel
	// data is read and the
	// channel closes
	for n := range f {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		// 10 sets of 3 - 12 will be written to the channel
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

// func main() {
// 	c := factorial(4)
// 	for n := range c {
// 		fmt.Println(n)
// 	}
// }
//
// func factorial(n int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		total := 1
// 		for i := n; i > 0; i-- {
// 			total *= i
// 		}
// 		out <- total
// 		close(out)
// 	}()
// 	return out
// }
//
// Challenge:
// -- Change the code above to execute 100 factorial
// -- computations concurrently and in parallel.
// -- Use the pipeline pattern to accomplish this
