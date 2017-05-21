package main

import "fmt"

func main() {
	mySlice := make([]int, 5)
	fmt.Println(mySlice) //[0 0 0 0 0]
	mySlice[0] = 5
	fmt.Println(mySlice) //[5 0 0 0 0]
	mySlice[0]++         // OR mySlice[0] += 1 OR mySlice[0] = mySlice[0] + 1
	fmt.Println(mySlice) //[6 0 0 0 0]
}
