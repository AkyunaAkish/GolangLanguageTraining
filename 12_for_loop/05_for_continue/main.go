package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
			// if number is even
			// don't run following logic,
			// just go to the next iteration of the loop
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
