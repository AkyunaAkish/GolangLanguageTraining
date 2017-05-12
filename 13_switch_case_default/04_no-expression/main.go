package main

import "fmt"

func main() {
	//a switch with no expression will only run the first case
	//that is true otherwise only default if none are true
	switch {
	case "Medhi" == "Medhi":
		fmt.Println("Sup Medhi and Daniel")
	case "Jacob" == "Not Jacob":
		fmt.Println("Sup Jacob")
	case "Daniel" == "Daniel":
		fmt.Println("Sup Daniel")
	default:
		fmt.Println("No matches")
	}
}
