package main

import "fmt"

func main() {
	if !true {
		fmt.Println("This didn't run")
	}

	if !false {
		fmt.Println("This ran")
	}
}
