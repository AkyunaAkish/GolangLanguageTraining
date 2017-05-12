package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Medhi", "Daniel":
		fmt.Println("Sup Medhi and Daniel")
	case "Jacob":
		fmt.Println("Sup Jacob")
	default:
		fmt.Println("No matches")
	}
}
