package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4]) //[3 4]
	fmt.Println(mySlice[2])   //3
	fmt.Println("mySlice"[2]) //83(returns ASCII number for string value "S")
}
