package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday", "Wednesday", "Thursday"}
	// fmt.Println()
	//removes "Tuesaday"
	//:1 takes everything before position 1(not including position 1)
	//2: takes the second position and beyond
	mySlice = append(mySlice[:1], mySlice[2:]...)

	fmt.Println(mySlice)
}
