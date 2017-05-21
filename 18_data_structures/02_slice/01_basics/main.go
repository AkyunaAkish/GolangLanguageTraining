package main

import "fmt"

func main() {
	//with no number between the []
	//that tells you it is a slice
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice)) //len of slice
	fmt.Println(cap(mySlice)) //len of slice and array beyond slice that slice is referenced to
	// to set the length and capacity of a slice
	// you can use make([]int, length, capacity)
	// OR new([100]int)[0:50]
	// using make() is good if you know the ballpark size
	// your slice might grow to in order to limit
	// the amount of times golang needs to create
	// and double the capacity of the array underlying
	// the slice

	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	//USE ***********make()**************** most of the time
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	// using just 1 number sets both length and capacity
	// make([]string, 3)

	// var example(var creating a slice of slice of strings)
	// var students [][]string
}
