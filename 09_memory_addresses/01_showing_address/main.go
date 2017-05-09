package main

import "fmt"

func main() {
	a := 42

	fmt.Println("a - ", a)
	//hexadecimal memory address
	fmt.Println("a's memory address hexadecimal - ", &a)
	//decimal memory address
	fmt.Printf("%d", &a)
}
