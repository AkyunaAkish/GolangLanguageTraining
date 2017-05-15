package main

import "fmt"

func main() {
	//anonymous func expression
	greeting := func() {
		fmt.Println("Hello World!")
	}

	greeting()
}
