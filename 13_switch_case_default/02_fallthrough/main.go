package main

import "fmt"

func main() {
	//fallthrough allows you to say if a case is true
	//act as if the next case is also true
	//and run both cases
	switch "Medhi" {
	case "Medhi":
		fmt.Println("Sup Medhi")
		fallthrough
	case "Daniel":
		fmt.Println("Sup Daniel")
	case "Jacob":
		fmt.Println("Sup Jacob")
	default:
		fmt.Println("No matches")
	}
}
