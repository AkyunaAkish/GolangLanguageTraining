package main

import "fmt"

func main() {
	if !true && !false {
		fmt.Println("This didnt run because both need to be true")
	}

	if !false && true {
		fmt.Println("This ran")
	}
}
