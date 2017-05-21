package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)
	// len of 3 capacity of 5 (still dynamic)
	// but that's how it starts out
	//immediate index value setting
	//is only available when youve
	//set the length of the slice
	//using make()
	//OTHERWISE you have to always use appen()
	greeting[0] = "yo"
	greeting[1] = "there"
	greeting[2] = "chelovek"
	// greeting[3] = "error" //this throws an error because len is only 3
	// use append instead
	greeting = append(greeting, "No Error")

	fmt.Println(greeting[3])

	//APPENDING ONE SLICE TO ANOTHER
	firstSlice := []int{1, 2, 3}
	secondSlice := []int{4, 5, 6}
	//must use "..." to spread items in slice as args
	newSlice := append(firstSlice, secondSlice...)
	fmt.Println(newSlice)
}
