package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i == 111 {
			fmt.Println(i)
			break
		}
	}
}
