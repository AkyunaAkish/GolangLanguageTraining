package main

import "fmt"

func factorial(x int) int {
	if x == 2 {
		//final value is multiplyed by 2
		return 2
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}
