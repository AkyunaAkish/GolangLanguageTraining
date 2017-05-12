package main

import "fmt"

func main() {
	//no need for break statements
	//because by default there is no fallthrough in switches
	//which means if you forget a break the code won't continue
	//to check for matches after a case runs
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Sup Daniel")
	case "Medhi":
		fmt.Println("Sup Medhi")
	default:
		fmt.Println("No matches")
	}
}
