package main

import "fmt"

func main() {
	//Jane and John are the arguments
	greet("Jane")
	greet("John")
}

//name is the parameter
func greet(name string) {
	fmt.Println(name)
}
