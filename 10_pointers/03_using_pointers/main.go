package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)  //address
	fmt.Println(*b) //value

	*b = 40
	fmt.Println(a) // 40 because b was simply a pointer to a and b was set to 40
}
