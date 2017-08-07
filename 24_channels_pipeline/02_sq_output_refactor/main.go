package main

import "fmt"

func main() {
	// set up the pipeline and consume the output
	// in a way that ranges over any amount
	// of numbers instead of needing to
	// make a print statement for every num
	for n := range sq(gen(2, 3)) {
		fmt.Println(n) // 4 then 9
	}
	// return a squared channel then square it again
	// for n := range sq(sq(gen(2, 3))) {
	// 	fmt.Println(n) // 16 then 81
	// }
}

func gen(nums ...int) chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)
	}()
	return out
}

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
