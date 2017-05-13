package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	// if you try to use food out of the
	// scope of the if statement you will get
	// an error
}
